package utils

import (
	"unicode"
)

func RuneToInt(char rune) int {
	if unicode.IsUpper(char) {
		return int(char-'A') + 27
	} else {
		return int(char-'a') + 1
	}
}
