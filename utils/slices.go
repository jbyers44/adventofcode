package utils

// Contains returns a bool indicating whether list contains char.
func Contains[T comparable](list []T, t T) bool {
	for _, v := range list {
		if v == t {
			return true
		}
	}

	return false
}

// IndexOf returns the index of t in list, if any.
func IndexOf[T comparable](list []T, t T) int {
	for i, v := range list {
		if v == t {
			return i
		}
	}

	return -1
}

// Split returns two slices based on the index of t in list.
func Split[T comparable](list []T, t T) (a []T, b []T) {
	i := IndexOf(list, t)
	return list[:i], list[i:]
}

// IsUnique returns a bool indicating whether or not every rune in buf is unique.
func IsUnique[T comparable](list []T) bool {
	a := []T{}

	for _, val := range list {
		if Contains(a, val) {
			return false
		}
		a = append(a, val)
	}

	return true
}

// UniqueIndex finds the first index at which the following len characters are unique.
func UniqueIndex[T comparable](list []T, size int) (index int) {
	index = -1

	for i := 0; i < len(list)-size; i++ {
		if IsUnique(list[i : i+size]) {
			index = i
			break
		}
	}

	return
}

func Intersect[T comparable](la []T, lb []T) []T {
	l := []T{}
	hash := make(map[T]struct{})

	for _, v := range la {
		hash[v] = struct{}{}
	}

	for _, v := range lb {
		if _, ok := hash[v]; ok {
			l = append(l, v)
		}
	}

	return l
}

func IntersectMany[T comparable](a [][]T) []T {
	l := a[0]

	for i := 1; i < len(a); i++ {
		l = Intersect(l, a[i])
	}

	return l
}
