package argtwo

import (
	"fmt"
)

const (
	ArgTwoMask uint16 = 0x003F // 0 000 000000 111111
	BlockA     uint16 = 0x0001 // 0 000 000000 000001
	BlockB     uint16 = 0x0002 // 0 000 000000 000010
	BlockC     uint16 = 0x0003 // 0 000 000000 000011
	BlockD     uint16 = 0x0004 // 0 000 000000 000100
	BlockE     uint16 = 0x0005 // 0 000 000000 000101
	BlockF     uint16 = 0x0006 // 0 000 000000 000110
	BlockG     uint16 = 0x0007 // 0 000 000000 000111
	BlockH     uint16 = 0x0008 // 0 000 000000 001000
	BlockI     uint16 = 0x0009 // 0 000 000000 001001
	BlockJ     uint16 = 0x000A // 0 000 000000 001010
	BlockK     uint16 = 0x000B // 0 000 000000 001011
	BlockL     uint16 = 0x000C // 0 000 000000 001100
	BlockM     uint16 = 0x000D // 0 000 000000 001101
	BlockN     uint16 = 0x000E // 0 000 000000 001110
	BlockO     uint16 = 0x000F // 0 000 000000 001111
	BlockP     uint16 = 0x0010 // 0 000 000000 010000
	BlockQ     uint16 = 0x0011 // 0 000 000000 010001
	BlockR     uint16 = 0x0012 // 0 000 000000 010010
	BlockS     uint16 = 0x0013 // 0 000 000000 010011
	BlockT     uint16 = 0x0014 // 0 000 000000 010100
	BlockU     uint16 = 0x0015 // 0 000 000000 010101
	BlockV     uint16 = 0x0016 // 0 000 000000 010110
	BlockW     uint16 = 0x0017 // 0 000 000000 010111
	BlockX     uint16 = 0x0018 // 0 000 000000 011000
	BlockY     uint16 = 0x0019 // 0 000 000000 011001
	BlockZ     uint16 = 0x001A // 0 000 000000 011010
	Block27    uint16 = 0x001B // 0 000 000000 011011
	Block28    uint16 = 0x001C // 0 000 000000 011100
	Block29    uint16 = 0x001D // 0 000 000000 011101
	Block30    uint16 = 0x001E // 0 000 000000 011110
	Block31    uint16 = 0x001F // 0 000 000000 011111
	Block32    uint16 = 0x0020 // 0 000 000000 100000
	Block33    uint16 = 0x0021 // 0 000 000000 100001
	Block34    uint16 = 0x0022 // 0 000 000000 100010
	Block35    uint16 = 0x0023 // 0 000 000000 100011
	Block36    uint16 = 0x0024 // 0 000 000000 100100
	Block37    uint16 = 0x0025 // 0 000 000000 100101
	Block38    uint16 = 0x0026 // 0 000 000000 100110
	Block39    uint16 = 0x0027 // 0 000 000000 100111
	Block40    uint16 = 0x0028 // 0 000 000000 101000
	Block41    uint16 = 0x0029 // 0 000 000000 101001
	Block42    uint16 = 0x002A // 0 000 000000 101010
	Block43    uint16 = 0x002B // 0 000 000000 101011
	Block44    uint16 = 0x002C // 0 000 000000 101100
	Block45    uint16 = 0x002D // 0 000 000000 101101
	Block46    uint16 = 0x002E // 0 000 000000 101110
	Block47    uint16 = 0x002F // 0 000 000000 101111
	Block48    uint16 = 0x0030 // 0 000 000000 110000
	Block49    uint16 = 0x0031 // 0 000 000000 110001
	Block50    uint16 = 0x0032 // 0 000 000000 110010
	Block51    uint16 = 0x0033 // 0 000 000000 110011
	Block52    uint16 = 0x0034 // 0 000 000000 110100
	Block53    uint16 = 0x0035 // 0 000 000000 110101
	Block54    uint16 = 0x0036 // 0 000 000000 110110
	Block55    uint16 = 0x0037 // 0 000 000000 110111
	Block56    uint16 = 0x0038 // 0 000 000000 111000
	Block57    uint16 = 0x0039 // 0 000 000000 111001
	Block58    uint16 = 0x003A // 0 000 000000 111010
	Block59    uint16 = 0x003B // 0 000 000000 111011
	Block60    uint16 = 0x003C // 0 000 000000 111100
	Block61    uint16 = 0x003D // 0 000 000000 111101
	Block62    uint16 = 0x003E // 0 000 000000 111110
	Floor      uint16 = 0x0000 // 0 000 000000 111111
	Blank      uint16 = 0x0000 // 0 000 000000 000000
)

func GetBlockLabel(cmd uint16) string {
	// Generate arg 2
	arg2 := ArgTwoMask & cmd

	if arg2 == Floor {
		return "Floor"
	} else if arg2 < 27 {
		return fmt.Sprintf("[%c]", arg2+64)
	} else {
		return ""
	}
}
