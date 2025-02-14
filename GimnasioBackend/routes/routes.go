package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"backendGimnasio/controllers"
	
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/usuarios", func(c *gin.Context) {
		controllers.CreateUsuario(c, db)
	})


	// Ruta para renovar la mensualidad de un usuario
	r.POST("/usuarios/:id/renovar", func(c *gin.Context) {
		controllers.RenovarMensualidad(c,db)
	})

	r.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})
	
}