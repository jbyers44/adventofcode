package utils

func Contains(list []rune, char rune) bool {
	for _, v := range list {
		if v == char {
			return true
		}
	}

	return false
}
