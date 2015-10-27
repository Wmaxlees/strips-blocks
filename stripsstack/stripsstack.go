package stripsstack

import (
	"fmt"
	"github.com/Wmaxlees/strips-blocks/argone"
	"github.com/Wmaxlees/strips-blocks/argtwo"
	"github.com/Wmaxlees/strips-blocks/opcodes"
)

type StripsStack struct {
	stack [][]uint16
}

func (stack *StripsStack) Push(item []uint16) {
	stack.stack = append(stack.stack, item)
}

func (stack *StripsStack) Append(other [][]uint16) {
	stack.stack = append(stack.stack, other...)
}

func (stack *StripsStack) Pop() {
	stack.stack = stack.stack[:len(stack.stack)-1]
}

func (stack *StripsStack) Peek() []uint16 {
	return stack.stack[len(stack.stack)-1]
}

func (stack *StripsStack) GetLength() int {
	return len(stack.stack)
}

func (stack *StripsStack) Print() {
	for i := len(stack.stack) - 1; i >= 0; i-- {
		for _, item := range stack.stack[i] {
			opCode, _, _ := opcodes.GetComponents(uint16(item))
			if opCode == uint16(opcodes.OnOpCode) || opCode == uint16(opcodes.StackOpCode) || opCode == uint16(opcodes.UnstackOpCode) {
				fmt.Printf("%s(%s, %s) ", opcodes.GetOpCode(item), argone.GetBlockLabel(item), argtwo.GetBlockLabel(item))
			} else {
				fmt.Printf("%s(%s) ", opcodes.GetOpCode(item), argone.GetBlockLabel(item))
			}
		}
		fmt.Println()
	}
}
