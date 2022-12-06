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
	file := utils.OpenFile("../../input/day6.txt")
	lines := utils.FileLines(file)

	a := 4

	index := utils.UniqueIndex([]rune(lines[0]), a)

	fmt.Printf("%v\n", index+a)
}

func part2() {
	file := utils.OpenFile("../../input/day6.txt")
	lines := utils.FileLines(file)

	a := 14

	index := utils.UniqueIndex([]rune(lines[0]), a)

	fmt.Printf("%v\n", index+a)
}
