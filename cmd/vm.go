package main

import "fmt"

var NilObj = NewNilValue()
var TrueObj = NewBoolValue(true)
var FalseObj = NewBoolValue(false)

// VM Operations
const (
	OpNil = iota
	OpTrue
	OpFalse
	OpNumber
	OpString

	OpAdd
	OpSub
	OpMul
	OpDiv
	OpMod
	OpNot
	OpNegate
)

type Chunk = []int

func DebugChunk(chunk Chunk) {
	for i := 0; i < len(chunk); i++ {
		ins := chunk[i]

		switch ins {
		case OpNil:
			fmt.Println("OpNil", i)
		case OpTrue:
			fmt.Println("OpTrue", i)
		case OpFalse:
			fmt.Println("OpFalse", i)
		}
	}
}

type VM struct {
	chunk Chunk
	stack *Stack
	ip    int
}

func NewVM(chunk Chunk) *VM {
	return &VM{chunk: chunk, stack: NewStack(), ip: 0}
}

func (vm *VM) Run() {
	for vm.ip <= len(vm.chunk) {
		op := vm.chunk[vm.ip]
		switch op {
		case OpTrue:
			fallthrough

		default:
			panic("Run() is WIP, aborting")
		}
	}
}

func (vm *VM) Write(ops ...int) {
	vm.chunk = append(vm.chunk, ops...)
}
