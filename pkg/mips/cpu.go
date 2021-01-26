package mips

import (
	"fmt"

	"github.com/pkg/errors"
)

const GeneralPurposeRegisterCount = 32

type CPU struct {
	PC int // Program Counter

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

	DebugMode bool
}

func NewCPU(mem *Memory, debugMode bool) *CPU {
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
		SP:        0xFFFF,
		FP:        0,
		RA:        0,
		Registers: [32]*int32{},
		Memory:    mem,
		DebugMode: debugMode,
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

func (cpu *CPU) Fetch() (int, int) {
	ins := 0

	cpu.printAddr() // if debug mode

	for i := cpu.PC + 3; i >= cpu.PC; i-- {
		ins = (ins << 8) | int(cpu.Memory[i])
	}

	cpu.printRawData(ins) // if debug mode

	pc := cpu.PC
	cpu.PC += 4
	return ins, pc
}

func (cpu *CPU) Decode(insData int) (*Instruction, error) {
	opcode := insData >> 26
	opcodeType, err := getOpcodeType(opcode)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ins := &Instruction{}
	ins.Opcode = opcode
	ins.OpcodeType = opcodeType

	switch opcodeType {
	case OpcodeTypeR:
		insTypeR := &InstructionTypeR{}
		insTypeR.FuncCode = insData & 0x3F
		insTypeR.ShiftAmount = (insData >> 6) & 0x1F
		insTypeR.DestinationRegister = (insData >> 11) & 0x1F
		insTypeR.TargetRegister = (insData >> 16) & 0x1F
		insTypeR.SourceRegister = (insData >> 21) & 0x1F

		// t type instruction, opcode is always 0
		f, ok := FunctionTypeRMap[insTypeR.FuncCode]
		if !ok {
			return nil, fmt.Errorf("invalid func code: 0b%06b", insTypeR.FuncCode)
		}
		insTypeR.Function = f

		ins.TypeR = insTypeR
	case OpcodeTypeI:
		insTypeI := &InstructionTypeI{}
		insTypeI.Immediate = insData & 0xFFFF
		insTypeI.TargetRegister = (insData >> 16) & 0x1F
		insTypeI.SourceRegister = (insData >> 21) & 0x1F

		f, ok := FunctionTypeIMap[opcode]
		if !ok {
			return nil, fmt.Errorf("invalid opcode: 0b%06b", opcode)
		}
		if opcode == 0b000001 {
			f, ok = FunctionTypeITargetAddressMap[insTypeI.TargetRegister]
			if !ok {
				return nil, fmt.Errorf("invalid target register: %d", insTypeI.TargetRegister)
			}
		}
		insTypeI.Function = f

		ins.TypeI = insTypeI
	case OpcodeTypeJ:
		insTypeJ := &InstructionTypeJ{}
		insTypeJ.TargetAddress = insData & 0x3FFFFFF

		f, ok := FunctionTypeJMap[opcode]
		if !ok {
			return nil, fmt.Errorf("invalid opcode: 0b%06b", opcode)
		}
		insTypeJ.Function = f

		ins.TypeJ = insTypeJ
	}

	//cpu.printInstruction(ins) // debug

	return ins, nil
}

func (cpu *CPU) Execute(ins *Instruction) error {
	var err error
	switch ins.OpcodeType {
	case OpcodeTypeR:
		r := ins.TypeR
		err = r.Function(cpu, r.SourceRegister, r.TargetRegister, r.DestinationRegister, r.ShiftAmount)
	case OpcodeTypeI:
		i := ins.TypeI
		err = i.Function(cpu, i.SourceRegister, i.TargetRegister, i.Immediate)
	case OpcodeTypeJ:
		j := ins.TypeJ
		err = j.Function(cpu, j.TargetAddress)
	}
	return errors.WithStack(err)
}
