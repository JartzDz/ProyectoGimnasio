package routes

import (
	"backendGimnasio/models"
	"backendGimnasio/utils"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)
func genTarjetaID() string{
	tarjetaID := uuid.New().String()
	return tarjetaID
}
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/usuarios", func(c *gin.Context){
		var usuario models.Usuario

		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(400,gin.H{"error":"Datos invalidos"})
			return
		}
		if usuario.TipoUsuario == 1 {
            usuario.PagoMensual = false
            usuario.TarjetaID = "" 
        } else {
          
            tarjetaID := genTarjetaID()
            usuario.TarjetaID = tarjetaID
        }

		hashedPassword, err := utils.EncriptarContrasenia(usuario.Contrasenia)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al encriptar la contraseña"})
			return
		}
		usuario.Contrasenia = hashedPassword
		tarjetaID := genTarjetaID()
		usuario.TarjetaID = tarjetaID

		if err := models.CreateUsuario(db, &usuario); err != nil {
			c.JSON(500,gin.H{"error":"Error al crear el usuario"})
			return
		}

	

		c.JSON(201,usuario)
	})

	// Ruta para renovar la mensualidad de un usuario
	r.POST("/usuarios/:id/renovar", func(c *gin.Context) {
		usuarioID := c.Param("id")

		id, err := strconv.Atoi(usuarioID)
		if err != nil {
			c.JSON(400, gin.H{"error": "ID de usuario inválido"})
			return
		}

		if err := models.RenovarMensualidad(db, uint(id)); err != nil {
			c.JSON(500, gin.H{"error": "Error al renovar la mensualidad"})
			return
		}

		c.JSON(200, gin.H{"message": "Mensualidad renovada con éxito"})
	})
}