package controllers

import (
	"backendGimnasio/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
)


// Estructura para la solicitud de login
type LoginRequest struct {
	Email       string `json:"email" binding:"required"`
	Contrasenia string `json:"contrasenia" binding:"required"`
}

// Estructura para la respuesta de login
type LoginResponse struct {
	Token string `json:"token"`
}

// Función para generar el JWT
func GenerarJWT(usuario models.Usuario) string {
	contra := os.Getenv("JWT_CONTRA")

	var jwtKey = []byte(contra)
	claims := jwt.MapClaims{
		"id":       usuario.ID,
		"email":    usuario.Email,
		"nombre":   usuario.Nombre,
		"tipo":     usuario.TipoUsuario,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return ""
	}

	return tokenString
}

func Login(c *gin.Context, db *gorm.DB) {
	var loginRequest LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Buscar el usuario en la base de datos por email
	var usuario models.Usuario
	if err := db.Where("email = ?", loginRequest.Email).First(&usuario).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Verificar la contraseña
	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Contrasenia), []byte(loginRequest.Contrasenia)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
		return
	}

	// Generar el token JWT
	token := GenerarJWT(usuario)

	// Enviar el token como respuesta
	c.JSON(http.StatusOK, LoginResponse{Token: token})
}
