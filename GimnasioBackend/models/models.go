package models

import (
	"gorm.io/gorm"
	"time"
)

type Usuario struct {
	gorm.Model
	Nombre       string    `json:"nombre" gorm:"size:100;not null"`
	Email        string    `json:"email" gorm:"size:100;not null;unique"`
	Contrasenia  string    `json:"contrasenia" gorm:"not null"`
	TarjetaID    string    `json:"tarjeta_id" gorm:"unique"`
	Saldo        float64   `json:"saldo" gorm:"type:decimal(10,2);default:0;check:saldo >= 0"`
	TipoUsuario  int       `json:"tipo_usuario" gorm:"not null"`  
	FechaRegistro time.Time `json:"fecha_registro" gorm:"default:current_timestamp"`
}

// Migrate realiza la migración de la tabla 'usuarios'
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Usuario{})
}

// CreateUsuario crea un nuevo usuario en la base de datos
func CreateUsuario(db *gorm.DB, usuario *Usuario) error {
	result := db.Create(&usuario)
	return result.Error
}
