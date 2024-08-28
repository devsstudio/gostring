package helpers

import (
	"strings"
)

func GenerateAlias(text string) string {
	text = normalizeChars(text)

	// Remover cualquier '-' del string y reemplazarlo con un espacio
	text = strings.ReplaceAll(text, "-", " ")
	text = strings.ReplaceAll(text, "_", " ")

	// Remover cualquier espacio en blanco duplicado, y asegurar que todos los caracteres sean alfanuméricos
	text = strings.ReplaceAll(text, " ", "-")
	text = strings.ReplaceAll(text, `[^A-Za-z0-9-]`, "")

	// Convertir a minúsculas y hacer trim
	text = strings.ToLower(strings.TrimSpace(text))

	return text
}

func normalizeChars(text string) string {
	replace := map[rune]string{
		'À': "A", 'Á': "A", 'Â': "A", 'Ã': "A", 'Ä': "Ae", 'Å': "A", 'Æ': "A", 'Ă': "A",
		'à': "a", 'á': "a", 'â': "a", 'ã': "a", 'ä': "ae", 'å': "a", 'ă': "a", 'æ': "ae",
		'þ': "b", 'Þ': "B", 'Ç': "C", 'ç': "c", 'È': "E", 'É': "E", 'Ê': "E", 'Ë': "E",
		'è': "e", 'é': "e", 'ê': "e", 'ë': "e", 'Ğ': "G", 'ğ': "g", 'Ì': "I", 'Í': "I",
		'Î': "I", 'Ï': "I", 'İ': "I", 'ı': "i", 'ì': "i", 'í': "i", 'î': "i", 'ï': "i",
		'Ñ': "N", 'Ò': "O", 'Ó': "O", 'Ô': "O", 'Õ': "O", 'Ö': "Oe", 'Ø': "O", 'ö': "oe",
		'ø': "o", 'ð': "o", 'ñ': "n", 'ò': "o", 'ó': "o", 'ô': "o", 'õ': "o", 'Š': "S",
		'š': "s", 'Ş': "S", 'ș': "s", 'Ș': "S", 'ş': "s", 'ß': "ss", 'ț': "t", 'Ț': "T",
		'Ù': "U", 'Ú': "U", 'Û': "U", 'Ü': "Ue", 'ù': "u", 'ú': "u", 'û': "u", 'ü': "ue",
		'Ý': "Y", 'ý': "y", 'ÿ': "y", 'Ž': "Z", 'ž': "z",
	}

	var normalized strings.Builder
	for _, char := range text {
		if replacement, exists := replace[char]; exists {
			normalized.WriteString(replacement)
		} else {
			normalized.WriteRune(char)
		}
	}

	return normalized.String()
}
