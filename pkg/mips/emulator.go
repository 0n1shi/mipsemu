package mips

import (
	"errors"
	"fmt"

	errs "github.com/pkg/errors"
)

type Emulator struct {
	CPU              *CPU
	Memory           *Memory
	InstructionCount int

	DebugMode bool
	TextSize  int
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
		TextSize:         len(text),
	}, nil
}

func (emu *Emulator) Run() error {
	if emu.CPU.DebugMode {
		fmt.Println("# trace")
		fmt.Println("------------------------------------------------------------")
	}
	for {
		data, _ := emu.CPU.Fetch()
		if data == 0x00 {
			fmt.Println("(invalid opcode)")
			break
		}
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
	emu.CPU.printData()
	emu.CPU.printStack()
}
