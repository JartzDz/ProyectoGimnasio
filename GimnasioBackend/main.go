package main

import (
    "backendGimnasio/config" 
    "backendGimnasio/models" 
    "backendGimnasio/routes" 
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    // Configuración de la base de datos
    db := config.ConexionBD()

    // Migración automática de modelos
    models.Migrate(db)

    // Configurar el enrutador de Gin
    r := gin.Default()

    // Rutas
	routes.SetupRoutes(r, db)

    // Iniciar el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	} 
	
		
}
