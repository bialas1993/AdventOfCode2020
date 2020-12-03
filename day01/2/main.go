package main

import (
	"fmt"
	"strconv"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

func main() {
	numbers := []int{}

	for _, number := range utils.FileRows("/day01/1/input.txt") {
		i, err := strconv.Atoi(number)
		if err == nil {
			numbers = append(numbers, i)
		}
	}

	for _, i := range numbers {
		for _, j := range numbers[1:] {
			for _, k := range numbers[2:] {
				if i+j+k == 2020 {
					fmt.Printf("%d + %d + %d = 2020, %d * %d * %d = %d\n", i, j, k, i, j, k, j*i*k)
					break
				}
			}
		}
	}
}
