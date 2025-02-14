package main

import (
    "backendGimnasio/config" 
    "backendGimnasio/models" 
    "backendGimnasio/routes" 
    "github.com/gin-gonic/gin"
    "log"
    "github.com/gin-contrib/cors"

)

func main() {
    // Configuración de la base de datos
    db := config.ConexionBD()

    // Migración automática de modelos
    models.Migrate(db)

    // Configurar el enrutador de Gin
    r := gin.Default()

    // Configurar CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, 
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // Rutas
	routes.SetupRoutes(r, db)

    // Iniciar el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	} 
	
		
}
