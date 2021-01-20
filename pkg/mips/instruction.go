package mips

var FunctionCodeMap = map[byte]InstructionFunction{}

var OpcodeMap = map[byte]InstructionFunction{}

type InstructionFunction func(cpu *CPU, ins *Instruction) error

type Instruction struct {
	DestinationRegister int
	SourceRegister      int
	TargetRegister      int
	ShiftAmount         int
	Immediate           int
	Function            InstructionFunction
}
