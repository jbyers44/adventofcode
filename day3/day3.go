package main

import (
	"adventofcode2022/utils"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	items := "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	file := utils.OpenFile("../input/day3.txt")
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		runes := []rune(line)

		compartment1 := runes[:len(line)/2]
		compartment2 := runes[len(line)/2:]

		for _, item := range compartment2 {
			if contains(compartment1, item) {
				total += strings.Index(items, string(item)) + 1
				break
			}
		}
	}

	fmt.Printf("%v\n", total)
}

func part2() {
	items := "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	file := utils.OpenFile("../input/day3.txt")
	scanner := bufio.NewScanner(file)

	total := 0
	lines := 1
	elfgroup := 0
	elfgroups := make(map[int]map[int]string)
	for scanner.Scan() {
		if elfgroups[elfgroup] == nil {
			elfgroups[elfgroup] = make(map[int]string)
		}

		line := scanner.Text()

		elfgroups[elfgroup][lines] = line

		if (lines%3) == 0 && lines > 0 {
			runes := []rune(line)

			for _, item := range runes {
				if contains([]rune(elfgroups[elfgroup][lines-1]), item) && contains([]rune(elfgroups[elfgroup][lines-2]), item) {
					total += strings.Index(items, string(item)) + 1
					break
				}
			}

			elfgroup++
		}

		lines++
	}

	fmt.Printf("%v\n", total)
}

func contains(list []rune, char rune) bool {
	for _, v := range list {
		if v == char {
			return true
		}
	}

	return false
}
