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
	file := utils.OpenFile("../../input/day4.txt")
	lines := utils.FileLines(file)

	total := 0
	for _, line := range lines {
		n := utils.Ints(line)

		if n[0] <= n[2] && n[2] <= n[3] && n[3] <= n[1] || n[2] <= n[0] && n[0] <= n[1] && n[1] <= n[3] {
			total++
		}
	}

	fmt.Printf("%v\n", total)
}

func part2() {
	file := utils.OpenFile("../../input/day4.txt")
	lines := utils.FileLines(file)

	total := 0
	for _, line := range lines {
		n := utils.Ints(line)

		if n[1] >= n[2] && n[3] >= n[0] {
			total++
		}
	}

	fmt.Printf("%v\n", total)
}
