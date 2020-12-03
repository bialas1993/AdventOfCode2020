package main

import (
	"fmt"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

const (
	TreeMark = '#'
)

func checkTrees(downStep, rightStep int, rows []string) int {
	fmt.Printf("Down step: %d, right step: %d\n", downStep, rightStep)

	trees := 0
	step := 0
	for i := 0; i < len(rows); i = i + downStep {
		if rows[i][(step*rightStep)%(len(rows[i]))] == TreeMark {
			trees++
		}
		step++
	}

	return trees
}

func main() {
	rows := utils.FileRows("/day03/1/input.txt")

	t1 := checkTrees(1, 1, rows)
	t2 := checkTrees(1, 3, rows)
	t3 := checkTrees(1, 5, rows)
	t4 := checkTrees(1, 7, rows)
	t5 := checkTrees(2, 1, rows)
	fmt.Printf("Encountered trees step %d x %d = %d\n", 1, 1, t1)
	fmt.Printf("Encountered trees step %d x %d = %d\n", 1, 3, t2)
	fmt.Printf("Encountered trees step %d x %d = %d\n", 1, 5, t3)
	fmt.Printf("Encountered trees step %d x %d = %d\n", 1, 7, t4)
	fmt.Printf("Encountered trees step %d x %d = %d\n", 2, 1, t5)

	fmt.Printf("Encountered trees: %d\n", t1*t2*t3*t4*t5)
}
