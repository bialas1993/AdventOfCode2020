package main

import (
	"fmt"
	"strings"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

func main() {
	answers := 0
	groupAnswers := map[rune]struct{}{}
	for _, row := range utils.FileRows("/day06/1/input.txt") {
		normalizeRow := strings.Trim(row, " ")
		if normalizeRow != "" {
			for _, char := range normalizeRow {
				if _, ok := groupAnswers[char]; !ok {
					groupAnswers[char] = struct{}{}
				}
			}
		} else {
			answers = answers + len(groupAnswers)
			groupAnswers = map[rune]struct{}{}
		}
	}

	fmt.Printf("result: %#v\n", answers)
}
