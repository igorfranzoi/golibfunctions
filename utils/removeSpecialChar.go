package utils

import "regexp"

// RemoveSpecialChars - remove barras e caracteres especiais de uma string
func RemoveSpecialChars(strConversion string) string {
	// Definir expressão regular para encontrar barras e caracteres especiais
	regExp := regexp.MustCompile(`[\/\\<>:"|?*]`)

	strRet := regExp.ReplaceAllString(strConversion, "")

	return strRet
}
