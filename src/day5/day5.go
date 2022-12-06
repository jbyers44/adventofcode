package main

import (
	"adventofcode2022/utils"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file := utils.OpenFile("../../input/day5.txt")
	scanner := bufio.NewScanner(file)

	crates := make(map[int][]string)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		for c := 0; c < 9; c++ {
			a := c*4 + 1
			b := string(line[a])
			if _, err := strconv.Atoi(b); err != nil && b != " " {
				crates[c] = append(crates[c], string(line[a]))
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, " ")

		x, _ := strconv.Atoi(s[1])
		a, _ := strconv.Atoi(s[3])
		b, _ := strconv.Atoi(s[5])

		for i := 0; i < x; i++ {
			val := crates[a-1][0]
			crates[a-1] = crates[a-1][1:]
			crates[b-1] = append([]string{val}, crates[b-1]...)
		}
	}

	for i := 0; i < 9; i++ {
		fmt.Printf("%v", crates[i][0])
	}

	fmt.Printf("\n")
}

func part2() {
	file := utils.OpenFile("../../input/day5.txt")
	scanner := bufio.NewScanner(file)

	crates := make(map[int][]string)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		for c := 0; c < 9; c++ {
			a := c*4 + 1
			b := string(line[a])
			if _, err := strconv.Atoi(b); err != nil && b != " " {
				crates[c] = append(crates[c], string(line[a]))
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, " ")

		x, _ := strconv.Atoi(s[1])
		a, _ := strconv.Atoi(s[3])
		b, _ := strconv.Atoi(s[5])

		val := make([]string, x)
		copy(val, crates[a-1][0:x])
		crates[a-1] = crates[a-1][x:]
		crates[b-1] = append(val, crates[b-1]...)
	}

	for i := 0; i < 9; i++ {
		fmt.Printf("%v", crates[i][0])
	}
}
