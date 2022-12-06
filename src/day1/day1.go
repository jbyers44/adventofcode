package main

import (
	"adventofcode2022/utils"
	"fmt"
	"sort"
)

func main() {
	file := utils.OpenFile("../../input/day1.txt")
	lines := utils.FileLines(file)

	elfs := []int{0}

	for i, j := 0, 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			calories := utils.SumList(utils.StringsToInts(lines[j:i]))
			elfs = append(elfs, calories)
			j = i + 1
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfs)))

	p1 := elfs[0]
	p2 := utils.SumList(elfs[0:3])

	fmt.Printf("%v\n", p1)
	fmt.Printf("%v\n", p2)
}
