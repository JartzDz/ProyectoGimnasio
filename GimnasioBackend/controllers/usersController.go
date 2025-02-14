package controllers

import (
	"backendGimnasio/models"
	"backendGimnasio/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"github.com/google/uuid"
	"fmt"
	//"io"

)

func genTarjetaID() string{
	tarjetaID := uuid.New().String()
	return tarjetaID
}


// Función para registrar un nuevo usuario
func CreateUsuario(c *gin.Context, db *gorm.DB) {
	var usuario models.Usuario

	
	// Parsear el JSON del request
	if err := c.ShouldBindJSON(&usuario); err != nil {
		fmt.Println("Error en el binding:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	fmt.Printf("Datos después del binding: %+v\n", usuario)


	contraseniaGenerada := utils.GenerarContrasenia(12)

	hashedPassword, err := utils.EncriptarContrasenia(contraseniaGenerada)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encriptar la contraseña"})
		return
	}
	usuario.Contrasenia = hashedPassword

	// Lógica para asignar valores según el tipo de usuario
	if usuario.TipoUsuario == 1 {
		usuario.PagoMensual = false
		usuario.TarjetaID = ""
	} else {
		usuario.TarjetaID = genTarjetaID() // Función genTarjetaID() debe estar en utils o en el mismo controlador
	}

	fmt.Printf("Datos recibidos: %+v\n", usuario)

	// Crear el usuario en la base de datos
	if err := models.CreateUsuario(db, &usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}
	
	go utils.EnviarCorreo(usuario.Email, "Registro exitoso", "Bienvenido " +usuario.Nombre+ "\nTu contraseña es: "+contraseniaGenerada +" \nRecuerda cambiar tu contraseña")

	c.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado exitosamente. Revisa tu correo para la contraseña."})
}

// Función para renovar la mensualidad de un usuario
func RenovarMensualidad(c *gin.Context, db *gorm.DB) {
	usuarioID := c.Param("id")

	id, err := strconv.Atoi(usuarioID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// Llamar al modelo para renovar la mensualidad
	if err := models.RenovarMensualidad(db, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al renovar la mensualidad"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Mensualidad renovada con éxito"})
}
