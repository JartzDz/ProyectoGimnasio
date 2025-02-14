package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"backendGimnasio/controllers"
	
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	
	r.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})

	r.POST("/registro", func(c *gin.Context) {
		controllers.CreateUsuario(c, db)
	})
	
}