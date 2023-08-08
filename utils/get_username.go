package utils

import "strings"

// getUserNameFromEmail obtiene la primera parte (antes del símbolo '@') de una dirección de correo electrónico.
// Si la dirección de correo electrónico es válida, devuelve el nombre de usuario. Si no, devuelve una cadena vacía.
func GetUserNameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return parts[0]
	}
	return ""
}
