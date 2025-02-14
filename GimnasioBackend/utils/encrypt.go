package utils

import (
	"golang.org/x/crypto/bcrypt"
	"crypto/rand"
	"encoding/base64"
	"net/smtp"
	"fmt"
)

// EncriptarContrasenia encripta la contraseña usando bcrypt
func EncriptarContrasenia(contrasenia string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(contrasenia), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerarContrasenia(n int) string {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "defaultPass123" // En caso de error, retorna una por defecto
	}
	return base64.URLEncoding.EncodeToString(bytes)[:n] // Recortar a la longitud deseada
}


func EnviarCorreo(destinatario, asunto, cuerpo string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	remitente := "julian.lojanojara@gmail.com"
	contrasenia := "ugzqvbkrmoheicpx" 

	mensaje := "Subject: " + asunto + "\n\n" + cuerpo

	// Autenticación
	auth := smtp.PlainAuth("", remitente, contrasenia, smtpHost)

	// Enviar el correo
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, remitente, []string{destinatario}, []byte(mensaje))
	if err != nil {
		fmt.Println("Error enviando el correo:", err)
		return err
	}

	fmt.Println("Correo enviado exitosamente a:", destinatario)
	return nil
}
