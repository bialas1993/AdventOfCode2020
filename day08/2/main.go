package main

import (
	"fmt"
	"strings"

	"github.com/bialas1993/AdventOfCode2020/day08/vm"
	"github.com/bialas1993/AdventOfCode2020/utils"
)

func main() {
	ret := make(chan int, 1)
	opcodes := utils.FileRows("/day08/1/input.txt")
	for i := 0; i < len(opcodes); i++ {
		if strings.HasPrefix(opcodes[i], "jmp") {
			instrs := make([]string, len(opcodes))

			copy(instrs, opcodes)
			instrs[i] = strings.ReplaceAll(opcodes[i], "jmp", "nop")

			go func(opcodes []string) {
				v := vm.New(opcodes)
				for {
					if _, canceled := v.Cycle(true); canceled {
						ret <- v.Accumulator
						break
					}
				}
			}(instrs)
		}
	}

	fmt.Printf("\n\nacc: %#v", <-ret)
}
