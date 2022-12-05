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
	file := utils.OpenFile("../../input/day4.txt")
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		elves := strings.Split(line, ",")

		z := strings.Split(elves[0]+"-"+elves[1], "-")

		a, _ := strconv.Atoi(z[0])
		b, _ := strconv.Atoi(z[1])
		c, _ := strconv.Atoi(z[2])
		d, _ := strconv.Atoi(z[3])

		if a <= c && c <= d && d <= b || c <= a && a <= b && b <= d {
			total++
		}
	}

	fmt.Printf("%v\n", total)
}

func part2() {
	file := utils.OpenFile("../../input/day4.txt")
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		elves := strings.Split(line, ",")

		z := strings.Split(elves[0]+"-"+elves[1], "-")

		a, _ := strconv.Atoi(z[0])
		b, _ := strconv.Atoi(z[1])
		c, _ := strconv.Atoi(z[2])
		d, _ := strconv.Atoi(z[3])

		if b >= c && d >= a {
			total++
		}
	}

	fmt.Printf("%v\n", total)
}
