package mips

import "fmt"

const funcCodeBitMask = 0x3F

func getFuncByFuncCode(funcCode byte) (*Function, error) {
	if f, ok := FunctionCodeMap[funcCode]; ok {
		return &f, nil
	}
	return nil, fmt.Errorf("invalid func code: 0x%06X", funcCode)
}

func getFuncByOpcode(opcode byte) (*Function, error) {
	if f, ok := OpcodeMap[opcode]; ok {
		return &f, nil
	}
	return nil, fmt.Errorf("invalid opcode: 0x%06X", opcode)
}

func getInstruction(insData uint32)
