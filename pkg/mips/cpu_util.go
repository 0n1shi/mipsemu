package mips

import "fmt"

var registerNames = []string{
	"zero",
	"at",
	"v0",
	"v1",
	"a0",
	"a1",
	"a2",
	"a3",
	"t0",
	"t1",
	"t2",
	"t3",
	"t4",
	"t5",
	"t6",
	"t7",
	"s0",
	"s1",
	"s2",
	"s3",
	"s4",
	"s5",
	"s6",
	"s7",
	"t8",
	"t9",
	"k0",
	"k1",
	"gp",
	"sp",
	"fp",
	"ra",
}

func (cpu *CPU) printAddr() {
	if cpu.DebugMode {
		fmt.Printf("0x%08X:\t", cpu.PC)
	}
}

func (cpu *CPU) printRawData(data int) {
	if cpu.DebugMode {
		for i := 0; i < 4; i++ {
			fmt.Printf("%02X ", (data&(0xFF000000>>(i*8)))>>((3-i)*8))
		}
	}
	fmt.Print("\t\t")
}

// func (cpu *CPU) printInstruction(ins *Instruction) {
// 	switch ins.OpcodeType {
// 	case OpcodeTypeR:
// 		r := ins.TypeR
// 	case OpcodeTypeI:
// 		i := ins.TypeI
// 	case OpcodeTypeJ:
// 		j := ins.TypeJ
// 	}
// }
