package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/bialas1993/AdventOfCode2020/utils"
)

func main() {
	rules := makeRules(utils.FileBlob("/day07/1/input.txt"))
	c := resolve(rules, "shiny gold")

	fmt.Printf("result: %#v\n", c)
}

func resolve(rules map[string][]string, input string) int {
	bags := rules[input]
	if len(bags) == 0 {
		return 0
	}
	total := 0
	for _, b := range bags {
		name := b[2:]
		number, err := strconv.Atoi(b[:1])
		if err != nil {
			log.Fatal(err)
		}
		total += number + number*resolve(rules, name)
	}
	return total
}

func makeRules(data string) map[string][]string {
	data = strings.ReplaceAll(data, ".", "")
	data = strings.ReplaceAll(data, ",", "")
	data = strings.ReplaceAll(data, "contain ", "")
	data = strings.ReplaceAll(data, "bags", "")
	data = strings.ReplaceAll(data, "bag", "")
	re := regexp.MustCompile("(  )|(no other)")

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
