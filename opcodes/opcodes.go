package opcodes

import (
	"fmt"
	"github.com/Wmaxlees/strips-blocks/argone"
	"github.com/Wmaxlees/strips-blocks/argtwo"
)

const (
	OpCodeMask    uint16 = 0xF000 // 1 111 000000 000000
	NonActionMask uint16 = 0x8000 // 1 000 000000 000000
	StackOpCode   uint16 = 0x1000 // 0 001 000000 000000
	UnstackOpCode uint16 = 0x2000 // 0 010 000000 000000
	PickupOpCode  uint16 = 0x3000 // 0 011 000000 000000
	PutdownOpCode uint16 = 0x4000 // 0 100 000000 000000
	OnOpCode      uint16 = 0x9000 // 1 001 000000 000000
	HoldingOpCode uint16 = 0xA000 // 1 010 000000 000000
	ClearOpCode   uint16 = 0xB000 // 1 011 000000 000000
)

func GetOpCode(cmd uint16) string {
	switch cmd & OpCodeMask {
	case StackOpCode:
		return "Stack"
	case UnstackOpCode:
		return "Unstack"
	case PickupOpCode:
		return "Pickup"
	case OnOpCode:
		return "On"
	case HoldingOpCode:
		return "Holding"
	case ClearOpCode:
		return "Clear"
	default:
		return ""
	}
}

func PrintPredicate(cmd uint16) {
	fmt.Printf("%s(%s, %s)\n", GetOpCode(cmd), argone.GetBlockLabel(cmd), argtwo.GetBlockLabel(cmd))
}

func GetComponents(cmd uint16) (uint16, uint16, uint16) {
	opCode := cmd & OpCodeMask
	argOne := cmd & argone.ArgOneMask
	argTwo := cmd & argtwo.ArgTwoMask

	return opCode, argOne, argTwo
}
