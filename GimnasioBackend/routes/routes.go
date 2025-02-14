package routes

import ( 
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"backendGimnasio/models" 
	"backendGimnasio/utils"

)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/usuarios", func(c *gin.Context){
		var usuario models.Usuario

		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(400,gin.H{"error":"Datos invalidos"})
			return
		}
		
		hashedPassword, err := utils.EncriptarContrasenia(usuario.Contrasenia)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al encriptar la contraseña"})
			return
		}
		usuario.Contrasenia = hashedPassword

		if err := models.CreateUsuario(db, &usuario); err != nil {
			c.JSON(500,gin.H{"error":"Error al crear el usuario"})
			return
		}

		c.JSON(201,usuario)
	})
}