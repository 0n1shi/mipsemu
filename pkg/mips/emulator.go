package mips

import (
	"errors"

	errs "github.com/pkg/errors"
)

type Emulator struct {
	CPU              *CPU
	Memory           *Memory
	InstructionCount int

	DebugMode bool
}

func NewEmulator(text []byte, debug bool) (*Emulator, error) {
	if len(text) > MemorySize {
		return nil, errors.New("text is too big")
	}

	// store text on memory
	mem := &Memory{}
	for i, d := range text {
		mem[i] = d
	}

	return &Emulator{
		CPU:              NewCPU(mem, debug),
		Memory:           mem,
		InstructionCount: len(text) / 4, // 32bit per an instruction
		DebugMode:        debug,
	}, nil
}

func (emu *Emulator) Run() error {
	for i := 0; i < emu.InstructionCount; i++ {
		data, _ := emu.CPU.Fetch()
		ins, err := emu.CPU.Decode(data)
		if err != nil {
			return errs.WithStack(err)
		}
		err = emu.CPU.Execute(ins)
		if err != nil {
			return errs.WithStack(err)
		}
	}
	return nil
}

func (emu *Emulator) Dump() {
	emu.CPU.printCPU()
	emu.CPU.printStack()
}
