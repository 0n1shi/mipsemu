package mips

import "errors"

type OpcodeType byte

const (
	OpcodeTypeR OpcodeType = iota
	OpcodeTypeI
	OpcodeTypeJ
)

var OpcodeMapTypeR = map[byte]Function{
	0b000000: func(cpu *CPU, ins *Instruction) error { return nil }, // dummy function
}
var OpcodeMapTypeI = map[byte]Function{
	0b001000: Addi,
	0b001001: Addiu,
	0b001100: Andi,
	0b000100: Beq,
	0b000001: Bgez, // or Bltz
	0b000111: Bgtz,
	0b000110: Blez,
	// 0b000001: Bltz // or Bgez,
	0b000101: Bne,
	0b100000: Lb,
	0b100100: Lbu,
	0b100001: Lh,
	0b100101: Lhu,
	0b001111: Lui,
	0b100011: Lw,
	0b110001: Lwc1,
	0b001101: Ori,
	0b101000: Sb,
	0b001010: Slti,
	0b001011: Sltiu,
	0b101001: Sh,
	0b101011: Sw,
	0b111001: Swc1,
	0b001110: Xori,
}
var OpcodeMapTypeJ = map[byte]Function{
	0b000010: J,
	0b000011: Jal,
}

func getOpcodeType(opcode byte) (OpcodeType, error) {
	if _, ok := OpcodeMapTypeR[opcode]; ok {
		return OpcodeTypeR, nil
	}
	if _, ok := OpcodeMapTypeI[opcode]; ok {
		return OpcodeTypeI, nil
	}
	if _, ok := OpcodeMapTypeJ[opcode]; ok {
		return OpcodeTypeJ, nil
	}
	return 0, errors.New("invalid opcode")
}
