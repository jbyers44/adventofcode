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
	file := utils.OpenFile("../../input/day5.txt")
	lines := utils.FileLines(file)

	setup, moves := utils.Split(lines, "")

	crates := setupCrates(setup)

	for _, move := range moves[1:] {
		n := utils.Ints(move)

		for i := 0; i < n[0]; i++ {
			val := crates[n[1]][0]
			crates[n[1]] = crates[n[1]][1:]
			crates[n[2]] = append([]string{val}, crates[n[2]]...)
		}
	}

	for i := 1; i < 10; i++ {
		fmt.Printf("%v", crates[i][0])
	}

	fmt.Printf("\n")
}

func part2() {
	file := utils.OpenFile("../../input/day5.txt")
	lines := utils.FileLines(file)

	setup, moves := utils.Split(lines, "")

	crates := setupCrates(setup)

	for _, move := range moves[1:] {
		n := utils.Ints(move)

		val := make([]string, n[0])
		copy(val, crates[n[1]][0:n[0]])
		crates[n[1]] = crates[n[1]][n[0]:]
		crates[n[2]] = append(val, crates[n[2]]...)
	}

	for i := 1; i < 10; i++ {
		fmt.Printf("%v", crates[i][0])
	}

	fmt.Printf("\n")
}

func setupCrates(s []string) map[int][]string {
	crates := make(map[int][]string, 9)

	for _, line := range s[:len(s)-1] {
		for c := 0; c < 9; c++ {
			i := c*4 + 1
			b := string(line[i])
			if b != " " {
				crates[c+1] = append(crates[c+1], b)
			}
		}
	}

	return crates
}
