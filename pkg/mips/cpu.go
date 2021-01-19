package mips

const GeneralPurposeRegisterCount = 32

var Registers []string = []string{
	// Always returns 0
	"zero",

	// (assembler temporary) Reserved for use by assembler
	"at",

	// Value returned by subroutine
	"v0", "v1",

	// (arguments) First four parameters for a subroutine
	"a0", "a1", "a2", "a3",

	// (temporaries) Subroutines can use without saving
	"t0", "t1", "t2", "t3", "t4", "t5", "t6", "t7",

	// Subroutine register variables, must be restored before returning
	"s0", "s1", "s2", "s3", "s4", "s5", "s6", "s7",

	// (temporaries) Subroutines can use without saving
	"t8", "t9",

	// Reserved for use by interrupt/trap handler; may change under your feet
	"k0", "k1",

	// Global pointer; used to access "static" or "extern" variables
	"gp",

	// Stack pointer
	"sp",

	// Frame pointer or ninth subroutine variable
	"fp",

	// Return address for subroutine
	"ra",
}

const InstructionSize = 8 // byte

type CPU struct {
	PC int32

	Registers [GeneralPurposeRegisterCount]int32
}

func NewCPU() *CPU {
	cpu := &CPU{}
	cpu.Registers[29] = MemorySize
	return &CPU{}
}

func (cpu *CPU) Fetch() {

}
