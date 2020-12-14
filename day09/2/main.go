package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bialas1993/AdventOfCode2020/day09/numbers"
	"github.com/bialas1993/AdventOfCode2020/utils"
)

func main() {
	nums := []int{}

	for _, row := range utils.FileRows("/day09/1/input.txt") {
		n, _ := strconv.Atoi(strings.Trim(row, " "))
		nums = append(nums, n)
	}

	weakness := numbers.FindNumber(nums)

	fmt.Printf("%v\n", numbers.FindWeakness(weakness, nums))
}
