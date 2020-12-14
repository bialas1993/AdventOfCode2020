package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

func main() {
	rules := makeRules(utils.FileBlob("/day07/1/input.txt"))

	c := 0
	for k, v := range rules {
		if k != "shiny gold" {
			if resolve(rules, v) {
				c++
			}
		}
	}

	fmt.Printf("result: %#v\n", c)
}

func resolve(rules map[string][]string, input []string) bool {
	if len(input) == 0 {
		return false
	}
	new := []string{}
	for _, v := range input {
		if v == "shiny gold" {
			return true
		}
		for _, w := range rules[v] {
			new = append(new, w)
		}
	}
	return resolve(rules, new)
}

func makeRules(data string) map[string][]string {
	data = strings.ReplaceAll(data, ".", "")
	data = strings.ReplaceAll(data, ",", "")
	data = strings.ReplaceAll(data, "contain ", "")
	data = strings.ReplaceAll(data, "bags", "")
	data = strings.ReplaceAll(data, "bag", "")
	data = strings.ReplaceAll(data, "  ", " ")
	re := regexp.MustCompile("( [0-9] )|( no other)")
	rules := make(map[string][]string)

	for _, r := range strings.Split(data, "\n") {
		t := re.Split(strings.TrimSpace(r), -1)
		if t[1] != "" {
			rules[t[0]] = t[1:]
		} else {
			rules[t[0]] = []string{}
		}
	}

	return rules
}
