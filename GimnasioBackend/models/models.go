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
	PagoMensual bool   `json:"pago_mensual" gorm:"default:false"` 
	FechaExpiracion time.Time `json:"fecha_expiracion"`

}


// Migrate realiza la migración de la tabla 'usuarios'
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Usuario{})
}

func CreateUsuario(db *gorm.DB, usuario *Usuario) error {
	if usuario.PagoMensual {
		usuario.FechaExpiracion = usuario.FechaRegistro.AddDate(0, 1, 0) // Un mes después
	}
	result := db.Create(&usuario)
	return result.Error          
}

func RenovarMensualidad(db *gorm.DB, usuarioID uint) error {
	// Obtener el usuario por ID
	var usuario Usuario
	if err := db.First(&usuario, usuarioID).Error; err != nil {
		return err // Si no se encuentra el usuario, retorna el error
	}

	// Actualizar el estado del pago mensual y la fecha de expiración
	usuario.PagoMensual = true
	usuario.FechaExpiracion = time.Now().AddDate(0, 1, 0) // Un mes a partir de ahora

	// Guardar los cambios en la base de datos
	if err := db.Save(&usuario).Error; err != nil {
		return err // Si hay un error al guardar, retorna el error
	}

	return nil // Si todo sale bien, retorna nil
}