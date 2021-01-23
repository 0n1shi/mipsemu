package mips

import "errors"

type OpcodeType byte

const (
	OpcodeTypeInvalid OpcodeType = iota
	OpcodeTypeR
	OpcodeTypeI
	OpcodeTypeJ
)

func getOpcodeType(opcode int) (OpcodeType, error) {
	if _, ok := FunctionTypeRMap[opcode]; ok {
		return OpcodeTypeR, nil
	}
	if _, ok := FunctionTypeIMap[opcode]; ok {
		return OpcodeTypeI, nil
	}
	if _, ok := FunctionTypeJMap[opcode]; ok {
		return OpcodeTypeJ, nil
	}
	return OpcodeTypeInvalid, errors.New("invalid opcode")
}
