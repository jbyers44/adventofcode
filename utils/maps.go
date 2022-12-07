package utils

import (
	"sort"
)

type Ordered interface {
	int | float64 | ~string
}

func SortMap[T, V Ordered](m map[T]V) []T {
	sorted := make([]T, 0, len(m))
	for t := range m {
		sorted = append(sorted, t)
	}

	sort.Slice(sorted, func(a, b int) bool {
		return m[sorted[a]] > m[sorted[b]]
	})

	return sorted
}
