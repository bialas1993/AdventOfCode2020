package main

import (
	"fmt"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

const (
	TreeMark = '#'
)

func main() {
	trees := 0

	for i, row := range utils.FileRows("/day03/1/input.txt") {
		if row[(i*3)%(len(row))] == TreeMark {
			trees++
		}
	}

	fmt.Printf("Encountered trees: %d\n", trees)
}
