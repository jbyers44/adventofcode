package main

import (
	"adventofcode2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	file := utils.OpenFile("../../input/day7.txt")
	lines := utils.FileLines(file)

	cur := []string{"/"}
	sizes := make(map[string]int)
	for _, line := range lines {
		words := strings.Split(line, " ")

		switch words[0] {
		case "$":
			if words[1] == "cd" {
				switch words[2] {
				case "/":
					cur = []string{"/"}
				case "..":
					cur = cur[:len(cur)-1]
				default:
					cur = append(cur, words[2])
				}
			} else {
				continue
			}
		case "dir":
			continue
		default:
			val, _ := strconv.Atoi(words[0])
			for i := 0; i < len(cur); i++ {
				path := strings.Join(cur[:i+1], ".")
				sizes[path] += val
			}
		}
	}

	l := utils.SortMap(sizes)

	p1 := 0
	p2 := 0
	for _, d := range l {
		if sizes[d] <= 100000 {
			p1 += sizes[d]
		}

		if sizes[d] >= (sizes["/"] - 40000000) {
			p2 = sizes[d]
		}
	}

	fmt.Printf("%v\n", p1)
	fmt.Printf("%v\n", p2)
}
