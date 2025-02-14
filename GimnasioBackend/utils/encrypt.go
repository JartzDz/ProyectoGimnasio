package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// EncriptarContrasenia encripta la contraseña usando bcrypt
func EncriptarContrasenia(contrasenia string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(contrasenia), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}


