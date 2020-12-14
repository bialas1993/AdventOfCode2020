package vm

import (
	"strconv"
	"strings"
)

type Instruction func(*VM)

type VM struct {
	Program        []Instruction
	Accumulator    int
	Executed       map[int]bool
	ProgramCounter int
}

func (vm *VM) Cycle(terminateRepeat bool) (bool, bool) {
	if vm.Executed[vm.ProgramCounter] && terminateRepeat {
		return false, false
	}

	if vm.ProgramCounter == len(vm.Program) {
		return false, true
	}

	vm.Program[vm.ProgramCounter](vm)
	vm.Executed[vm.ProgramCounter] = true
	vm.ProgramCounter++

	return true, false
}

func New(opcodes []string) *VM {
	var cmd Instruction
	vm := &VM{
		Executed: make(map[int]bool),
	}

	for _, opcode := range opcodes {
		instr := strings.Fields(opcode)
		operation := instr[0]
		sign := string(instr[1][:1])
		v, err := strconv.Atoi(instr[1][1:])
		if err != nil {
			panic(err)
		}

		switch operation {
		case "jmp":
			cmd = func(vm *VM) {
				// fmt.Printf("%s %v%v\n", operation, sign, v)
				vm.ProgramCounter--
				if sign == "-" {
					vm.ProgramCounter -= v
					return
				}

				vm.ProgramCounter += v
			}
		case "acc":
			cmd = func(vm *VM) {
				// fmt.Printf("%s %v%v\n", operation, sign, v)

				if sign == "-" {
					vm.Accumulator -= v
					return
				}
				vm.Accumulator += v
			}
		default:
			cmd = func(vm *VM) {
				// fmt.Printf("%s %v%v\n", operation, sign, v)
			}
		}

		vm.Program = append(vm.Program, cmd)
	}

	return vm
}
