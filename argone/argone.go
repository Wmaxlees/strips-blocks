package argone

import (
	"fmt"
)

const (
	ArgOneMask uint16 = 0x0FC0 // 0 000 111111 000000
	BlockA     uint16 = 0x0040 // 0 000 000001 000000
	BlockB     uint16 = 0x0080 // 0 000 000010 000000
	BlockC     uint16 = 0x00C0 // 0 000 000011 000000
	BlockD     uint16 = 0x0100 // 0 000 000100 000000
	BlockE     uint16 = 0x0140 // 0 000 000101 000000
	BlockF     uint16 = 0x0180 // 0 000 000110 000000
	BlockG     uint16 = 0x01C0 // 0 000 000111 000000
	BlockH     uint16 = 0x0200 // 0 000 001000 000000
	BlockI     uint16 = 0x0240 // 0 000 001001 000000
	BlockJ     uint16 = 0x0280 // 0 000 001010 000000
	BlockK     uint16 = 0x02C0 // 0 000 001011 000000
	BlockL     uint16 = 0x0300 // 0 000 001100 000000
	BlockM     uint16 = 0x0340 // 0 000 001101 000000
	BlockN     uint16 = 0x0380 // 0 000 001110 000000
	BlockO     uint16 = 0x03C0 // 0 000 001111 000000
	BlockP     uint16 = 0x0400 // 0 000 010000 000000
	BlockQ     uint16 = 0x0440 // 0 000 010001 000000
	BlockR     uint16 = 0x0480 // 0 000 010010 000000
	BlockS     uint16 = 0x04C0 // 0 000 010011 000000
	BlockT     uint16 = 0x0500 // 0 000 010100 000000
	BlockU     uint16 = 0x0540 // 0 000 010101 000000
	BlockV     uint16 = 0x0580 // 0 000 010110 000000
	BlockW     uint16 = 0x05C0 // 0 000 010111 000000
	BlockX     uint16 = 0x0600 // 0 000 011000 000000
	BlockY     uint16 = 0x0640 // 0 000 011001 000000
	BlockZ     uint16 = 0x0680 // 0 000 011010 000000
	Block27    uint16 = 0x06C0 // 0 000 011011 000000
	Block28    uint16 = 0x0700 // 0 000 011100 000000
	Block29    uint16 = 0x0740 // 0 000 011101 000000
	Block30    uint16 = 0x0780 // 0 000 011110 000000
	Block31    uint16 = 0x07C0 // 0 000 011111 000000
	Block32    uint16 = 0x0800 // 0 000 100000 000000
	Block33    uint16 = 0x0840 // 0 000 100001 000000
	Block34    uint16 = 0x0880 // 0 000 100010 000000
	Block35    uint16 = 0x08C0 // 0 000 100011 000000
	Block36    uint16 = 0x0900 // 0 000 100100 000000
	Block37    uint16 = 0x0940 // 0 000 100101 000000
	Block38    uint16 = 0x0980 // 0 000 100110 000000
	Block39    uint16 = 0x09C0 // 0 000 100111 000000
	Block40    uint16 = 0x0A00 // 0 000 101000 000000
	Block41    uint16 = 0x0A40 // 0 000 101001 000000
	Block42    uint16 = 0x0A80 // 0 000 101010 000000
	Block43    uint16 = 0x0AC0 // 0 000 101011 000000
	Block44    uint16 = 0x0B00 // 0 000 101100 000000
	Block45    uint16 = 0x0B40 // 0 000 101101 000000
	Block46    uint16 = 0x0B80 // 0 000 101110 000000
	Block47    uint16 = 0x0BC0 // 0 000 101111 000000
	Block48    uint16 = 0x0C00 // 0 000 110000 000000
	Block49    uint16 = 0x0C40 // 0 000 110001 000000
	Block50    uint16 = 0x0C80 // 0 000 110010 000000
	Block51    uint16 = 0x0CC0 // 0 000 110011 000000
	Block52    uint16 = 0x0D00 // 0 000 110100 000000
	Block53    uint16 = 0x0D40 // 0 000 110101 000000
	Block54    uint16 = 0x0D80 // 0 000 110110 000000
	Block55    uint16 = 0x0DC0 // 0 000 110111 000000
	Block56    uint16 = 0x0E00 // 0 000 111000 000000
	Block57    uint16 = 0x0E40 // 0 000 111001 000000
	Block58    uint16 = 0x0E80 // 0 000 111010 000000
	Block59    uint16 = 0x0EC0 // 0 000 111011 000000
	Block60    uint16 = 0x0F00 // 0 000 111100 000000
	Block61    uint16 = 0x0F40 // 0 000 111101 000000
	Block62    uint16 = 0x0F80 // 0 000 111110 000000
	Floor      uint16 = 0x0000 // 0 000 111111 000000
	Blank      uint16 = 0x0000 // 0 000 000000 000000
)

func GetBlockLabel(cmd uint16) string {
	// Generate arg 1
	arg1 := ArgOneMask & cmd
	arg1 = arg1 >> 6

	if arg1 == Blank {
		return "nil"
	} else if arg1 < 27 {
		return fmt.Sprintf("[%c]", arg1+64)
	} else {
		return ""
	}
}
