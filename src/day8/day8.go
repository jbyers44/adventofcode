package main

import (
	"adventofcode2022/utils"
	"fmt"
	"sort"
)

func main() {
	file := utils.OpenFile("../../input/day8.txt")
	lines := utils.FileLines(file)

	trees := [][]int{}
	for _, line := range lines {
		trees = append(trees, utils.StringToInts(line))
	}

	total := 0
	scores := []int{}
	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees[y]); x++ {
			visible, score := check(trees, x, y)
			if visible {
				total++
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	fmt.Printf("%v\n", total)
	fmt.Printf("%v\n", scores[len(scores)-1])
}

func check(trees [][]int, x int, y int) (bool, int) {
	up, udist := checkup(trees, x, y)
	down, ddist := checkdown(trees, x, y)
	left, ldist := checkleft(trees, x, y)
	right, rdist := checkright(trees, x, y)

	return up || down || left || right, udist * ddist * ldist * rdist
}

func checkup(trees [][]int, x int, y int) (bool, int) {
	for i := y - 1; i >= 0; i-- {
		if trees[i][x] >= trees[y][x] {
			return false, (y - i)
		}
	}
	return true, y
}

func checkdown(trees [][]int, x int, y int) (bool, int) {
	for i := y + 1; i < len(trees); i++ {
		if trees[i][x] >= trees[y][x] {
			return false, (i - y)
		}
	}
	return true, len(trees) - y - 1
}

func checkleft(trees [][]int, x int, y int) (bool, int) {
	for i := x - 1; i >= 0; i-- {
		if trees[y][i] >= trees[y][x] {
			return false, (x - i)
		}
	}
	return true, x
}

func checkright(trees [][]int, x int, y int) (bool, int) {
	for i := x + 1; i < len(trees[y]); i++ {
		if trees[y][i] >= trees[y][x] {
			return false, (i - x)
		}
	}
	return true, len(trees) - x - 1
}
