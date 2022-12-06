package main

import (
	"adventofcode2022/utils"
	"fmt"
)

func main() {
	file := utils.OpenFile("../../input/day2.txt")
	lines := utils.FileLines(file)

	p1, p2 := 0, 0
	for _, line := range lines {
		l := []rune(line)

		x, y := int(l[0]-'A')+1, int(l[2]-'X')+1

		p1 += y
		p2 += utils.Mod(y+x, 3) + 1

		p1 += utils.Mod((y-x+1), 3) * 3
		p2 += (y - 1) * 3
	}

	fmt.Printf("%v\n", p1)
	fmt.Printf("%v\n", p2)
}
