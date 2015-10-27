package stripsstack

import (
	"fmt"
	"github.com/Wmaxlees/strips-blocks/argone"
	"github.com/Wmaxlees/strips-blocks/argtwo"
	"github.com/Wmaxlees/strips-blocks/opcodes"
)

type StripsStack struct {
	stack []uint16
}

func (stack *StripsStack) Push(item uint16) {
	stack.stack = append(stack.stack, item)
}

func (stack *StripsStack) Append(other []uint16) {
	stack.stack = append(stack.stack, other...)
}

func (stack *StripsStack) Pop() {
	stack.stack = stack.stack[:len(stack.stack)-1]
}

func (stack *StripsStack) Peek() uint16 {
	return stack.stack[len(stack.stack)-1]
}

func (stack *StripsStack) GetLength() int {
	return len(stack.stack)
}

func (stack *StripsStack) Print() {
	for i := len(stack.stack) - 1; i >= 0; i-- {
		opCode, _, _ := opcodes.GetComponents(stack.stack[i])
		if opCode == uint16(opcodes.OnOpCode) || opCode == uint16(opcodes.StackOpCode) || opCode == uint16(opcodes.UnstackOpCode) {
			fmt.Printf("%s(%s, %s)\n", opcodes.GetOpCode(stack.stack[i]), argone.GetBlockLabel(stack.stack[i]), argtwo.GetBlockLabel(stack.stack[i]))
		} else {
			fmt.Printf("%s(%s)\n", opcodes.GetOpCode(stack.stack[i]), argone.GetBlockLabel(stack.stack[i]))
		}
	}
}
