package mips

// func code
var FunctionTypeRMap = map[int]FunctionTypeR{
	0b100000: Add,
	0b100001: Addu,
	0b100100: And,
	0b001101: Break,
	0b011010: Div,
	0b011011: Divu,
	0b001001: Jalr,
	0b001000: Jr,
	0b010000: Mfhi,
	0b010010: Mflo,
	0b010001: Mthi,
	0b010011: Mtlo,
	0b011000: Mult,
	0b011001: Multu,
	0b100111: Nor,
	0b100101: Or,
	0b000000: Sll,
	0b000100: Sllv,
	0b101010: Slt,
	0b101011: Sltu,
	0b000011: Sra,
	0b000111: Srav,
	0b000010: Srl,
	0b000110: Srlv,
	0b100010: Sub,
	0b100011: Subu,
	0b001100: Syscall,
	0b100110: Xor,
}

// Opcode
var FunctionTypeIMap = map[int]FunctionTypeI{
	// 0b000000: Dummy, // R type
	0b001000: Addi,
	0b001001: Addiu,
	0b001100: Andi,
	0b000100: Beq,
	0b000001: DummyTypeI, // bgez or bltz
	0b000111: Bgtz,
	0b000110: Blez,
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

var FunctionTypeITargetAddressMap = map[int]FunctionTypeI{
	0b00001: Bgez,
	0b00000: Bltz,
}

// Opcode
var FunctionTypeJMap = map[int]FunctionTypeJ{
	0b000010: J,
	0b000011: Jal,
}

type Instruction struct {
	Opcode     int
	OpcodeType OpcodeType
	TypeR      *InstructionTypeR
	TypeI      *InstructionTypeI
	TypeJ      *InstructionTypeJ
}

type InstructionTypeR struct {
	SourceRegister      int
	TargetRegister      int
	DestinationRegister int
	ShiftAmount         int
	FuncCode            int

	Function func(cpu *CPU, rs int, rt int, rd int, sa int) error
}

type InstructionTypeI struct {
	SourceRegister int
	TargetRegister int
	Immediate      int

	Function func(cpu *CPU, rs int, rt int, imm int) error
}

type InstructionTypeJ struct {
	TargetAddress int

	Function func(cpu *CPU, addr int) error
}
