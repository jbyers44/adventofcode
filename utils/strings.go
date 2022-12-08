package utils

import (
	"log"
	"regexp"
	"strconv"
)

func StringsToInts(s []string) (ints []int) {
	ints = make([]int, len(s))

	for n, a := range s {
		b, err := strconv.Atoi(a)
		if err != nil {
			log.Fatal(err)
		}
		ints[n] = b
	}

	return
}

func findInts(s string) []string {
	re := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(s, -1)

	return matches
}

func Ints(s string) (ints []int) {
	return StringsToInts(findInts(s))
}

func IndexMany(s string, res string) [][]int {
	re := regexp.MustCompile(res)

	matches := re.FindAllStringIndex(s, -1)

	return matches
}

func IndexManyLetters(s string) []int {
	re := regexp.MustCompile(`[a-zA-Z]`)

	matches := re.FindAllStringIndex(s, -1)
	i := make([]int, len(matches))
	for index, match := range matches {
		i[index] = match[0]
	}

	return i
}

func StringToInts(s string) []int {
	ints := []int{}

	for i := 0; i < len(s); i++ {
		v, _ := strconv.Atoi(string(s[i]))
		ints = append(ints, v)
	}

	return ints
}
