package utils

func Mod(a, b int) int {
	return (a%b + b) % b
}

func SumList(l []int) int {
	total := 0

	for _, v := range l {
		total += v
	}

	return total
}
