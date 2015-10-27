package opcodes

import (
	"fmt"
	"github.com/Wmaxlees/strips-blocks/argone"
	"github.com/Wmaxlees/strips-blocks/argtwo"
)

type Mask uint16

const (
	OpCodeMask    Mask = 0xF000 // 1 111 000000 000000
	NonActionMask Mask = 0x8000 // 1 000 000000 000000
)

type OpCode uint16

const (
	StackOpCode   OpCode = 0x1000 // 0 001 000000 000000
	UnstackOpCode OpCode = 0x2000 // 0 010 000000 000000
	PickupOpCode  OpCode = 0x3000 // 0 011 000000 000000
	PutdownOpCode OpCode = 0x4000 // 0 100 000000 000000
	OnOpCode      OpCode = 0x9000 // 1 001 000000 000000
	HoldingOpCode OpCode = 0xA000 // 1 010 000000 000000
	ClearOpCode   OpCode = 0xB000 // 1 011 000000 000000
)

func GetOpCode(cmd uint16) string {
	switch cmd & uint16(OpCodeMask) {
	case uint16(StackOpCode):
		return "Stack"
	case uint16(UnstackOpCode):
		return "Unstack"
	case uint16(PickupOpCode):
		return "Pickup"
	case uint16(OnOpCode):
		return "On"
	case uint16(HoldingOpCode):
		return "Holding"
	case uint16(ClearOpCode):
		return "Clear"
	default:
		return ""
	}
}

func PrintPredicate(cmd uint16) {
	fmt.Printf("%s(%s, %s)\n", GetOpCode(cmd), argone.GetBlockLabel(cmd), argtwo.GetBlockLabel(cmd))
}

func GetComponents(cmd uint16) (uint16, uint16, uint16) {
	opCode := cmd & uint16(OpCodeMask)
	argOne := cmd & uint16(argone.ArgOneMask)
	argTwo := cmd & uint16(argtwo.ArgTwoMask)

	return opCode, argOne, argTwo
}
