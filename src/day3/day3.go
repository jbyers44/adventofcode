package main

import (
	"adventofcode2022/utils"
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	file := utils.OpenFile("../../input/day3.txt")
	lines := utils.FileLines(file)

	total := 0
	for _, line := range lines {
		runes := []rune(line)

		a := runes[:len(line)/2]
		b := runes[len(line)/2:]

		total += utils.RuneToInt(utils.Intersect(a, b)[0])
	}

	fmt.Printf("%v\n", total)
}

func part2() {
	file := utils.OpenFile("../../input/day3.txt")
	lines := utils.FileLines(file)

	total := 0
	for i := 0; i < len(lines); i += 3 {
		a := []rune(lines[i])
		b := []rune(lines[i+1])
		c := []rune(lines[i+2])
		total += utils.RuneToInt(utils.IntersectMany([][]rune{a, b, c})[0])
	}

	fmt.Printf("%v\n", total)
}
