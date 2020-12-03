package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

func parseRow(line string) (int, int, string, string) {
	parts := strings.Split(line, " ")
	occurrences := strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(occurrences[0])
	max, _ := strconv.Atoi(occurrences[1])
	char := parts[1][:1]
	pass := parts[2]

	return min, max, char, pass
}

func main() {
	correctPasswords := 0
	for _, row := range utils.FileRows("/day02/1/input.txt") {
		i, j, char, password := parseRow(row)

		if (2<<(strings.Index(password[i-1:i], char)*-1))<<(strings.Index(password[j-1:j], char)*-1) == 4 {
			correctPasswords++
		}
	}

	fmt.Printf("Correct passwords: %d\n", correctPasswords)
}
