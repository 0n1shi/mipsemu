package mips

const GeneralPurposeRegisterCount = 32

type CPU struct {
	PC int32 // Program Counter

	Zero int32 // Always returns 0
	AT   int32 // (assembler temporary) Reserved for use by assembler
	V0   int32 // Value returned by subroutine
	V1   int32 // Value returned by subroutine
	A0   int32 // (arguments) First four parameters for a subroutine
	A1   int32 // (arguments) First four parameters for a subroutine
	A2   int32 // (arguments) First four parameters for a subroutine
	A3   int32 // (arguments) First four parameters for a subroutine
	T0   int32 // (temporaries) Subroutines can use without saving
	T1   int32 // (temporaries) Subroutines can use without saving
	T2   int32 // (temporaries) Subroutines can use without saving
	T3   int32 // (temporaries) Subroutines can use without saving
	T4   int32 // (temporaries) Subroutines can use without saving
	T5   int32 // (temporaries) Subroutines can use without saving
	T6   int32 // (temporaries) Subroutines can use without saving
	T7   int32 // (temporaries) Subroutines can use without saving
	S0   int32 // Subroutine register variables, must be restored before returning
	S1   int32 // Subroutine register variables, must be restored before returning
	S2   int32 // Subroutine register variables, must be restored before returning
	S3   int32 // Subroutine register variables, must be restored before returning
	S4   int32 // Subroutine register variables, must be restored before returning
	S5   int32 // Subroutine register variables, must be restored before returning
	S6   int32 // Subroutine register variables, must be restored before returning
	S7   int32 // Subroutine register variables, must be restored before returning
	T8   int32 // (temporaries) Subroutines can use without saving
	T9   int32 // (temporaries) Subroutines can use without saving
	K0   int32 // Reserved for use by interrupt/trap handler; may change under your feet
	K1   int32 // Reserved for use by interrupt/trap handler; may change under your feet
	GP   int32 // Global pointer; used to access "static" or "extern" variables
	SP   int32 // Stack pointer
	FP   int32 // Frame pointer or ninth subroutine variable
	RA   int32 // Return address for subroutine

	Registers [GeneralPurposeRegisterCount]*int32

	Memory *Memory
}

func NewCPU(mem *Memory) *CPU {
	cpu := CPU{
		PC:        0,
		Zero:      0,
		AT:        0,
		V0:        0,
		V1:        0,
		A0:        0,
		A1:        0,
		A2:        0,
		A3:        0,
		T0:        0,
		T1:        0,
		T2:        0,
		T3:        0,
		T4:        0,
		T5:        0,
		T6:        0,
		T7:        0,
		S0:        0,
		S1:        0,
		S2:        0,
		S3:        0,
		S4:        0,
		S5:        0,
		S6:        0,
		S7:        0,
		T8:        0,
		T9:        0,
		K0:        0,
		K1:        0,
		GP:        0,
		SP:        0,
		FP:        0,
		RA:        0,
		Registers: [32]*int32{},
		Memory:    mem,
	}
	cpu.Registers = [32]*int32{
		&cpu.Zero,
		&cpu.AT,
		&cpu.V0,
		&cpu.V1,
		&cpu.A0,
		&cpu.A1,
		&cpu.A2,
		&cpu.A3,
		&cpu.T0,
		&cpu.T1,
		&cpu.T2,
		&cpu.T3,
		&cpu.T4,
		&cpu.T5,
		&cpu.T6,
		&cpu.T7,
		&cpu.S0,
		&cpu.S1,
		&cpu.S2,
		&cpu.S3,
		&cpu.S4,
		&cpu.S5,
		&cpu.S6,
		&cpu.S7,
		&cpu.T8,
		&cpu.T9,
		&cpu.K0,
		&cpu.K1,
		&cpu.GP,
		&cpu.SP,
		&cpu.FP,
		&cpu.RA,
	}

	return &cpu
}

func (cpu *CPU) Fetch() uint32 {
	ins := uint32(0)
	for i := 0; i < 4; i++ {
		ins = (ins << 8) | uint32(cpu.Memory[cpu.PC])
	}
	return ins
}

func (cpu *CPU) Decode(ins uint32) (*Instruction, error) {
	opcode := byte(ins >> 26)
	opcodeType, err := getOpcodeType(opcode)
	if err != nil {
		panic(err)
	}

	switch opcodeType {
	case OpcodeTypeR:
		funcCode := byte(ins & 0x3F)
	case OpcodeTypeI:
	case OpcodeTypeJ:
	}

	return &Instruction{}, nil
}

func (cpu *CPU) Execute(ins *Instruction) {
	ins.Function(cpu, ins)
}
