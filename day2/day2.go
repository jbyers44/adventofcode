package main

import (
	"adventofcode2022/utils"
	"bufio"
	"fmt"
)

func main() {
	p1_outcomes := [3][3]int{
		// A, B ,C
		{4, 1, 7}, // X, rock
		{8, 5, 2}, // Y, paper
		{3, 9, 6}, // Z, scissors
	}

	p2_outcomes := [3][3]int{
		// A, B ,C
		{3, 1, 2}, // X, lose
		{4, 5, 6}, // Y, draw
		{8, 9, 7}, // Z, win
	}

	choices := map[rune]int{
		'A': 0,
		'B': 1,
		'C': 2,
		'X': 0,
		'Y': 1,
		'Z': 2,
	}

	file := utils.OpenFile("../input/day2.txt")
	scanner := bufio.NewScanner(file)

	p1_total := 0
	p2_total := 0
	for scanner.Scan() {
		line := scanner.Text()

		chars := []rune(line)

		p1_total += p1_outcomes[choices[chars[2]]][choices[chars[0]]]
		p2_total += p2_outcomes[choices[chars[2]]][choices[chars[0]]]
	}

	fmt.Printf("%v, %v\n", p1_total, p2_total)
}
