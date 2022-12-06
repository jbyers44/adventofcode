package utils

func Contains(list []rune, char rune) bool {
	for _, v := range list {
		if v == char {
			return true
		}
	}

	return false
}

func IsUnique(buf []rune) bool {
	a := []rune{}

	for _, c := range buf {
		if Contains(a, c) {
			return false
		}
		a = append(a, c)
	}

	return true
}

func UniqueSlice(data []rune, len int) (unique []rune, index int) {
	buf := make([]rune, len)
	index = -1

	for i, c := range data {
		buf = append(buf, c)
		buf = buf[1:]
		if i > len && IsUnique(buf) {
			index = i
			break
		}
	}

	return
}
