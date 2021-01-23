package mips

import "errors"

type Emulator struct {
	CPU              *CPU
	Memory           *Memory
	InstructionCount int
}

func NewEmulator(text []byte) (*Emulator, error) {
	if len(text) > MemorySize {
		return nil, errors.New("text is too big")
	}

	// store text on memory
	mem := &Memory{}
	for i, d := range text {
		mem[i] = d
	}

	return &Emulator{
		CPU:              NewCPU(mem),
		Memory:           mem,
		InstructionCount: len(text) / 4, // 32bit per an instruction
	}, nil
}

func (emu *Emulator) Run() {
	for i := 0; i < emu.InstructionCount; i++ {
		insData := emu.CPU.Fetch()
		ins, err := emu.CPU.Decode(insData)
		if err != nil {
			panic(err)
		}
		emu.CPU.Execute(ins)
	}
}
