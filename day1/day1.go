package main

import (
	"adventofcode2022/utils"
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	file := utils.OpenFile("../input/day1.txt")
	scanner := bufio.NewScanner(file)

	elfs := []int{0}
	cur_elf := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elfs = append(elfs, 0)
			cur_elf++
		} else {
			calories, _ := strconv.Atoi(line)
			elfs[cur_elf] += calories
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfs)))

	fmt.Printf("%v\n", elfs[0])

	total := 0
	for _, calories := range elfs[0:3] {
		total += calories
	}

	fmt.Printf("%v\n", total)
}
