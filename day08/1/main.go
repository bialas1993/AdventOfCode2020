package main

import (
	"fmt"

	"github.com/bialas1993/AdventOfCode2020/day08/vm"
	"github.com/bialas1993/AdventOfCode2020/utils"
)

func main() {
	opcodes := utils.FileRows("/day08/1/input.txt")
	vm := vm.New(opcodes)
	for {
		if ok, _ := vm.Cycle(true); !ok {
			break
		}
	}

	fmt.Printf("\n\nacc: %#v", vm.Accumulator)
}
