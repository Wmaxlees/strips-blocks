package stripsstack

import (
	"fmt"
	"github.com/Wmaxlees/strips-blocks/argone"
	"github.com/Wmaxlees/strips-blocks/argtwo"
	"github.com/Wmaxlees/strips-blocks/opcodes"
)

type StripsStack [][]uint16

func (stack *StripsStack) Push(item []uint16) {
	*stack = append(*stack, item)
}

func (stack *StripsStack) Append(other [][]uint16) {
	*stack = append(*stack, other...)
}

func (stack *StripsStack) Pop() {
	*stack = (*stack)[:len(*stack)-1]
}

func (stack *StripsStack) Peek() []uint16 {
	return (*stack)[len(*stack)-1]
}

func (stack *StripsStack) GetLength() int {
	return len(*stack)
}

func (stack *StripsStack) Print() {
	for i := len(*stack) - 1; i >= 0; i-- {
		for _, item := range (*stack)[i] {
			opCode, _, _ := opcodes.GetComponents(uint16(item))
			if opCode == opcodes.OnOpCode || opCode == opcodes.StackOpCode || opCode == opcodes.UnstackOpCode {
				fmt.Printf("%s(%s, %s) ", opcodes.GetOpCode(item), argone.GetBlockLabel(item), argtwo.GetBlockLabel(item))
			} else {
				fmt.Printf("%s(%s) ", opcodes.GetOpCode(item), argone.GetBlockLabel(item))
			}
		}
		fmt.Println()
	}
}
