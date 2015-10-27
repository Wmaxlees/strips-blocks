package argtwo

import (
	"fmt"
)

type ArgTwo uint16

const (
	ArgTwoMask ArgTwo = 0x003F // 0 000 000000 111111
	BlockA     ArgTwo = 0x0001 // 0 000 000000 000001
	BlockB     ArgTwo = 0x0002 // 0 000 000000 000010
	BlockC     ArgTwo = 0x0003 // 0 000 000000 000011
	BlockD     ArgTwo = 0x0004 // 0 000 000000 000100
	BlockE     ArgTwo = 0x0005 // 0 000 000000 000101
	BlockF     ArgTwo = 0x0006 // 0 000 000000 000110
	BlockG     ArgTwo = 0x0007 // 0 000 000000 000111
	BlockH     ArgTwo = 0x0008 // 0 000 000000 001000
	BlockI     ArgTwo = 0x0009 // 0 000 000000 001001
	BlockJ     ArgTwo = 0x000A // 0 000 000000 001010
	BlockK     ArgTwo = 0x000B // 0 000 000000 001011
	BlockL     ArgTwo = 0x000C // 0 000 000000 001100
	BlockM     ArgTwo = 0x000D // 0 000 000000 001101
	BlockN     ArgTwo = 0x000E // 0 000 000000 001110
	BlockO     ArgTwo = 0x000F // 0 000 000000 001111
	BlockP     ArgTwo = 0x0010 // 0 000 000000 010000
	BlockQ     ArgTwo = 0x0011 // 0 000 000000 010001
	BlockR     ArgTwo = 0x0012 // 0 000 000000 010010
	BlockS     ArgTwo = 0x0013 // 0 000 000000 010011
	BlockT     ArgTwo = 0x0014 // 0 000 000000 010100
	BlockU     ArgTwo = 0x0015 // 0 000 000000 010101
	BlockV     ArgTwo = 0x0016 // 0 000 000000 010110
	BlockW     ArgTwo = 0x0017 // 0 000 000000 010111
	BlockX     ArgTwo = 0x0018 // 0 000 000000 011000
	BlockY     ArgTwo = 0x0019 // 0 000 000000 011001
	BlockZ     ArgTwo = 0x001A // 0 000 000000 011010
	Block27    ArgTwo = 0x001B // 0 000 000000 011011
	Block28    ArgTwo = 0x001C // 0 000 000000 011100
	Block29    ArgTwo = 0x001D // 0 000 000000 011101
	Block30    ArgTwo = 0x001E // 0 000 000000 011110
	Block31    ArgTwo = 0x001F // 0 000 000000 011111
	Block32    ArgTwo = 0x0020 // 0 000 000000 100000
	Block33    ArgTwo = 0x0021 // 0 000 000000 100001
	Block34    ArgTwo = 0x0022 // 0 000 000000 100010
	Block35    ArgTwo = 0x0023 // 0 000 000000 100011
	Block36    ArgTwo = 0x0024 // 0 000 000000 100100
	Block37    ArgTwo = 0x0025 // 0 000 000000 100101
	Block38    ArgTwo = 0x0026 // 0 000 000000 100110
	Block39    ArgTwo = 0x0027 // 0 000 000000 100111
	Block40    ArgTwo = 0x0028 // 0 000 000000 101000
	Block41    ArgTwo = 0x0029 // 0 000 000000 101001
	Block42    ArgTwo = 0x002A // 0 000 000000 101010
	Block43    ArgTwo = 0x002B // 0 000 000000 101011
	Block44    ArgTwo = 0x002C // 0 000 000000 101100
	Block45    ArgTwo = 0x002D // 0 000 000000 101101
	Block46    ArgTwo = 0x002E // 0 000 000000 101110
	Block47    ArgTwo = 0x002F // 0 000 000000 101111
	Block48    ArgTwo = 0x0030 // 0 000 000000 110000
	Block49    ArgTwo = 0x0031 // 0 000 000000 110001
	Block50    ArgTwo = 0x0032 // 0 000 000000 110010
	Block51    ArgTwo = 0x0033 // 0 000 000000 110011
	Block52    ArgTwo = 0x0034 // 0 000 000000 110100
	Block53    ArgTwo = 0x0035 // 0 000 000000 110101
	Block54    ArgTwo = 0x0036 // 0 000 000000 110110
	Block55    ArgTwo = 0x0037 // 0 000 000000 110111
	Block56    ArgTwo = 0x0038 // 0 000 000000 111000
	Block57    ArgTwo = 0x0039 // 0 000 000000 111001
	Block58    ArgTwo = 0x003A // 0 000 000000 111010
	Block59    ArgTwo = 0x003B // 0 000 000000 111011
	Block60    ArgTwo = 0x003C // 0 000 000000 111100
	Block61    ArgTwo = 0x003D // 0 000 000000 111101
	Block62    ArgTwo = 0x003E // 0 000 000000 111110
	Floor      ArgTwo = 0x0000 // 0 000 000000 111111
	Blank      ArgTwo = 0x0000 // 0 000 000000 000000
)

func GetBlockLabel(cmd uint16) string {
	// Generate arg 2
	arg2 := uint16(ArgTwoMask) & cmd

	if arg2 == uint16(Floor) {
		return "Floor"
	} else if arg2 < 27 {
		return fmt.Sprintf("[%c]", arg2+64)
	} else {
		return ""
	}
}
