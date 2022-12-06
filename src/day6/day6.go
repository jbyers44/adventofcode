package main

import (
	"adventofcode2022/utils"
	"bufio"
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	file := utils.OpenFile("../../input/day6.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	line := scanner.Text()

	index := utils.UniqueIndex([]rune(line), 4)

	fmt.Printf("%v\n", index+1)
}

func part2() {
	file := utils.OpenFile("../../input/day6.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	line := scanner.Text()

	index := utils.UniqueIndex([]rune(line), 14)

	fmt.Printf("%v\n", index+1)
}
