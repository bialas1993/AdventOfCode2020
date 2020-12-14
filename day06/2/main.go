package main

import (
	"fmt"
	"strings"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

func main() {
	answers := 0
	persons := 0
	groupAnswers := map[rune]int{}
	for _, row := range utils.FileRows("/day06/1/input.txt") {
		normalizeRow := strings.Trim(row, " ")
		if normalizeRow != "" {
			persons++
			for _, char := range normalizeRow {
				groupAnswers[char]++
			}
		} else {
			fmt.Printf("persons: %d, group: %#v\n", persons, groupAnswers)
			for _, row := range groupAnswers {
				if row == persons {
					answers++
				}
			}

			groupAnswers = map[rune]int{}
			persons = 0
		}
	}

	fmt.Printf("result: %#v\n", answers)
}
