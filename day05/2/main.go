package main

import (
	"fmt"
	"math"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

func main() {
	results := []int{}
	result := 0
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

	min, max := results[0], results[0]
	list := make(map[int]bool)

	for _, v := range results {
		list[v] = true
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	for i := min; i < max; i++ {
		if !list[i] {
			result = i
			break
		}
	}

	fmt.Printf("result: %#v\n", result)
}
