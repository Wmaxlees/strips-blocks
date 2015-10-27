package bitcodes

const OpCodeMask uint16 = 0xF000 // 1 111 000000 000000
const Arg1Mask uint16 = 0x0FC0   // 0 000 111111 000000
const Arg2Mask uint16 = 0x003F   // 0 000 000000 111111

const StackOpCode uint16 = 0x1000   // 0 001 000000 000000
const UnstackOpCode uint16 = 0x2000 // 0 010 000000 000000
const PickupOpCode uint16 = 0x3000  // 0 011 000000 000000
const PutdownOpCode uint16 = 0x4000 // 0 100 000000 000000

const OnCode uint16 = 0x9000      // 1 001 000000 000000
const HoldingCode uint16 = 0xA000 // 1 010 000000 000000

const Block01Arg1 uint16 = 0x0040 // 0 000 000001 000000
const Block02Arg1 uint16 = 0x0080 // 0 000 000010 000000
const Block03Arg1 uint16 = 0x00C0 // 0 000 000011 000000
const Block04Arg1 uint16 = 0x0100 // 0 000 000100 000000
const Block05Arg1 uint16 = 0x0140 // 0 000 000101 000000
const Block06Arg1 uint16 = 0x0180 // 0 000 000110 000000
const Block07Arg1 uint16 = 0x01C0 // 0 000 000111 000000
const Block08Arg1 uint16 = 0x0200 // 0 000 001000 000000
const Block09Arg1 uint16 = 0x0240 // 0 000 001001 000000
const Block10Arg1 uint16 = 0x0280 // 0 000 001010 000000
const Block11Arg1 uint16 = 0x02C0 // 0 000 001011 000000
const Block12Arg1 uint16 = 0x0300 // 0 000 001100 000000
const Block13Arg1 uint16 = 0x0340 // 0 000 001101 000000
const Block14Arg1 uint16 = 0x0380 // 0 000 001110 000000
const Block15Arg1 uint16 = 0x03C0 // 0 000 001111 000000
const Block16Arg1 uint16 = 0x0400 // 0 000 010000 000000
const Block17Arg1 uint16 = 0x0440 // 0 000 010001 000000
const Block18Arg1 uint16 = 0x0480 // 0 000 010010 000000
const Block19Arg1 uint16 = 0x04C0 // 0 000 010011 000000
const Block20Arg1 uint16 = 0x0500 // 0 000 010100 000000
const Block21Arg1 uint16 = 0x0540 // 0 000 010101 000000
const Block22Arg1 uint16 = 0x0580 // 0 000 010110 000000
const Block23Arg1 uint16 = 0x05C0 // 0 000 010111 000000
const Block24Arg1 uint16 = 0x0600 // 0 000 011000 000000
const Block25Arg1 uint16 = 0x0640 // 0 000 011001 000000
const Block26Arg1 uint16 = 0x0680 // 0 000 011010 000000
const Block27Arg1 uint16 = 0x06C0 // 0 000 011011 000000
const Block28Arg1 uint16 = 0x0700 // 0 000 011100 000000
const Block29Arg1 uint16 = 0x0740 // 0 000 011101 000000
const Block30Arg1 uint16 = 0x0780 // 0 000 011110 000000
const Block31Arg1 uint16 = 0x07C0 // 0 000 011111 000000
const Block32Arg1 uint16 = 0x0800 // 0 000 100000 000000
const Block33Arg1 uint16 = 0x0840 // 0 000 100001 000000
const Block34Arg1 uint16 = 0x0880 // 0 000 100010 000000
const Block35Arg1 uint16 = 0x08C0 // 0 000 100011 000000
const Block36Arg1 uint16 = 0x0900 // 0 000 100100 000000
const Block37Arg1 uint16 = 0x0940 // 0 000 100101 000000
const Block38Arg1 uint16 = 0x0980 // 0 000 100110 000000
const Block39Arg1 uint16 = 0x09C0 // 0 000 100111 000000
const Block40Arg1 uint16 = 0x0A00 // 0 000 101000 000000
const Block41Arg1 uint16 = 0x0A40 // 0 000 101001 000000
const Block42Arg1 uint16 = 0x0A80 // 0 000 101010 000000
const Block43Arg1 uint16 = 0x0AC0 // 0 000 101011 000000
const Block44Arg1 uint16 = 0x0B00 // 0 000 101100 000000
const Block45Arg1 uint16 = 0x0B40 // 0 000 101101 000000
const Block46Arg1 uint16 = 0x0B80 // 0 000 101110 000000
const Block47Arg1 uint16 = 0x0BC0 // 0 000 101111 000000
const Block48Arg1 uint16 = 0x0C00 // 0 000 110000 000000
const Block49Arg1 uint16 = 0x0C40 // 0 000 110001 000000
const Block50Arg1 uint16 = 0x0C80 // 0 000 110010 000000
const Block51Arg1 uint16 = 0x0CC0 // 0 000 110011 000000
const Block52Arg1 uint16 = 0x0D00 // 0 000 110100 000000
const Block53Arg1 uint16 = 0x0D40 // 0 000 110101 000000
const Block54Arg1 uint16 = 0x0D80 // 0 000 110110 000000
const Block55Arg1 uint16 = 0x0DC0 // 0 000 110111 000000
const Block56Arg1 uint16 = 0x0E00 // 0 000 111000 000000
const Block57Arg1 uint16 = 0x0E40 // 0 000 111001 000000
const Block58Arg1 uint16 = 0x0E80 // 0 000 111010 000000
const Block59Arg1 uint16 = 0x0EC0 // 0 000 111011 000000
const Block60Arg1 uint16 = 0x0F00 // 0 000 111100 000000
const Block61Arg1 uint16 = 0x0F40 // 0 000 111101 000000
const Block62Arg1 uint16 = 0x0F80 // 0 000 111110 000000
const Block63Arg1 uint16 = 0x0FC0 // 0 000 111111 000000
const FloorArg1 uint16 = 0x0000   // 0 000 000000 000000

const Block01Arg2 uint16 = 0x0001 // 0 000 000000 000001
const Block02Arg2 uint16 = 0x0002 // 0 000 000000 000010
const Block03Arg2 uint16 = 0x0003 // 0 000 000000 000011
const Block04Arg2 uint16 = 0x0004 // 0 000 000000 000100
const Block05Arg2 uint16 = 0x0005 // 0 000 000000 000101
const Block06Arg2 uint16 = 0x0006 // 0 000 000000 000110
const Block07Arg2 uint16 = 0x0007 // 0 000 000000 000111
const Block08Arg2 uint16 = 0x0008 // 0 000 000000 001000
const Block09Arg2 uint16 = 0x0009 // 0 000 000000 001001
const Block10Arg2 uint16 = 0x000A // 0 000 000000 001010
const Block11Arg2 uint16 = 0x000B // 0 000 000000 001011
const Block12Arg2 uint16 = 0x000C // 0 000 000000 001100
const Block13Arg2 uint16 = 0x000D // 0 000 000000 001101
const Block14Arg2 uint16 = 0x000E // 0 000 000000 001110
const Block15Arg2 uint16 = 0x000F // 0 000 000000 001111
const Block16Arg2 uint16 = 0x0010 // 0 000 000000 010000
const Block17Arg2 uint16 = 0x0011 // 0 000 000000 010001
const Block18Arg2 uint16 = 0x0012 // 0 000 000000 010010
const Block19Arg2 uint16 = 0x0013 // 0 000 000000 010011
const Block20Arg2 uint16 = 0x0014 // 0 000 000000 010100
const Block21Arg2 uint16 = 0x0015 // 0 000 000000 010101
const Block22Arg2 uint16 = 0x0016 // 0 000 000000 010110
const Block23Arg2 uint16 = 0x0017 // 0 000 000000 010111
const Block24Arg2 uint16 = 0x0018 // 0 000 000000 011000
const Block25Arg2 uint16 = 0x0019 // 0 000 000000 011001
const Block26Arg2 uint16 = 0x001A // 0 000 000000 011010
const Block27Arg2 uint16 = 0x001B // 0 000 000000 011011
const Block28Arg2 uint16 = 0x001C // 0 000 000000 011100
const Block29Arg2 uint16 = 0x001D // 0 000 000000 011101
const Block30Arg2 uint16 = 0x001E // 0 000 000000 011110
const Block31Arg2 uint16 = 0x001F // 0 000 000000 011111
const Block32Arg2 uint16 = 0x0020 // 0 000 000000 100000
const Block33Arg2 uint16 = 0x0021 // 0 000 000000 100001
const Block34Arg2 uint16 = 0x0022 // 0 000 000000 100010
const Block35Arg2 uint16 = 0x0023 // 0 000 000000 100011
const Block36Arg2 uint16 = 0x0024 // 0 000 000000 100100
const Block37Arg2 uint16 = 0x0025 // 0 000 000000 100101
const Block38Arg2 uint16 = 0x0026 // 0 000 000000 100110
const Block39Arg2 uint16 = 0x0027 // 0 000 000000 100111
const Block40Arg2 uint16 = 0x0028 // 0 000 000000 101000
const Block41Arg2 uint16 = 0x0029 // 0 000 000000 101001
const Block42Arg2 uint16 = 0x002A // 0 000 000000 101010
const Block43Arg2 uint16 = 0x002B // 0 000 000000 101011
const Block44Arg2 uint16 = 0x002C // 0 000 000000 101100
const Block45Arg2 uint16 = 0x002D // 0 000 000000 101101
const Block46Arg2 uint16 = 0x002E // 0 000 000000 101110
const Block47Arg2 uint16 = 0x002F // 0 000 000000 101111
const Block48Arg2 uint16 = 0x0030 // 0 000 000000 110000
const Block49Arg2 uint16 = 0x0031 // 0 000 000000 110001
const Block50Arg2 uint16 = 0x0032 // 0 000 000000 110010
const Block51Arg2 uint16 = 0x0033 // 0 000 000000 110011
const Block52Arg2 uint16 = 0x0034 // 0 000 000000 110100
const Block53Arg2 uint16 = 0x0035 // 0 000 000000 110101
const Block54Arg2 uint16 = 0x0036 // 0 000 000000 110110
const Block55Arg2 uint16 = 0x0037 // 0 000 000000 110111
const Block56Arg2 uint16 = 0x0038 // 0 000 000000 111000
const Block57Arg2 uint16 = 0x0039 // 0 000 000000 111001
const Block58Arg2 uint16 = 0x003A // 0 000 000000 111010
const Block59Arg2 uint16 = 0x003B // 0 000 000000 111011
const Block60Arg2 uint16 = 0x003C // 0 000 000000 111100
const Block61Arg2 uint16 = 0x003D // 0 000 000000 111101
const Block62Arg2 uint16 = 0x003E // 0 000 000000 111110
const Block63Arg2 uint16 = 0x003F // 0 000 000000 111111
const FloorArg2 uint16 = 0x0000   // 0 000 000000 000000