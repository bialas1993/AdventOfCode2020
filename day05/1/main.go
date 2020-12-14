package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

func main() {
	results := []int{}

	for _, row := range utils.FileRows("/day05/1/input.txt") {
		minRow := 0.0
		maxRow := 127.0
		minCell := 0.0
		maxCell := 7.0

		for _, char := range row[:7] {
			if char == 'F' {
				maxRow -= math.Ceil((maxRow - minRow) / 2)
			} else {
				minRow += math.Ceil((maxRow - minRow) / 2)
			}
		}

		for _, char := range row[7:] {
			if char == 'L' {
				maxCell -= math.Ceil((maxCell - minCell) / 2)
			} else if char == 'R' {
				minCell += math.Ceil((maxCell - minCell) / 2)
			}
		}

		results = append(results, int(minRow*8+minCell))
	}

	sort.Ints(results)
	fmt.Printf("result: %#v\n", results[len(results)-1])
}
