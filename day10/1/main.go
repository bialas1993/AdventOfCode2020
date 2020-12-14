package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

func main() {
	nums := []int{}

	for _, row := range utils.FileRows("/day10/1/input.txt") {
		n, err := strconv.Atoi(strings.Trim(row, " "))
		if err != nil {
			println("cannot parse number " + row)
		}
		nums = append(nums, n)
	}

	sort.Ints(nums)

	jolts := map[int]int{
		1: 0,
		2: 0,
		3: 1,
	}
	current := 0
	for _, n := range nums {
		if n-current > 0 && n-current <= 3 {
			jolts[(n-current)]++
			current = n
		} else {
			panic(fmt.Errorf("cannot find next adapter"))
		}
	}

	fmt.Printf("different: %d\n", jolts[1]*jolts[3])
}
