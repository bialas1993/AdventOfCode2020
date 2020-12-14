package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

func sumCombinations(number int, numbers []int, mem map[int]int) int {
	total := 0

	if len(numbers) == 0 {
		return 1
	}

	for _, next := range numbers {
		numbers = numbers[1:]
		if next-number > 3 {
			break
		}
		if _, ok := mem[next]; !ok {
			mem[next] = sumCombinations(next, numbers, mem)
		}
		total += mem[next]
	}
	return total
}

func main() {
	nums := []int{}

	for _, row := range utils.FileRows("/day10/1/input.txt") {
		n, err := strconv.Atoi(strings.Trim(row, " "))
		if err != nil {
			panic(fmt.Errorf("cannot parse number " + row))
		}
		nums = append(nums, n)
	}

	sort.Ints(nums)

	fmt.Printf("sum: %v\n", sumCombinations(0, nums, make(map[int]int)))
}
