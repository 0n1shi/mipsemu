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

func (cpu *CPU) printCPU() {
	fmt.Println("============================ CPU registers ==========================")
	fmt.Printf("Zero:%10d | ", cpu.Zero)
	fmt.Printf("AT:  %10d | ", cpu.AT)
	fmt.Printf("V0:  %10d | ", cpu.V0)
	fmt.Printf("V1:  %10d\n", cpu.V1)
	fmt.Printf("A0:  %10d | ", cpu.A0)
	fmt.Printf("A1:  %10d | ", cpu.A1)
	fmt.Printf("A2:  %10d | ", cpu.A2)
	fmt.Printf("A3:  %10d\n", cpu.A3)
	fmt.Printf("T0:  %10d | ", cpu.T0)
	fmt.Printf("T1:  %10d | ", cpu.T1)
	fmt.Printf("T2:  %10d | ", cpu.T2)
	fmt.Printf("T3:  %10d\n", cpu.T3)
	fmt.Printf("T4:  %10d | ", cpu.T4)
	fmt.Printf("T5:  %10d | ", cpu.T5)
	fmt.Printf("T6:  %10d | ", cpu.T6)
	fmt.Printf("T7:  %10d\n", cpu.T7)
	fmt.Printf("S0:  %10d | ", cpu.S0)
	fmt.Printf("S1:  %10d | ", cpu.S1)
	fmt.Printf("S2:  %10d | ", cpu.S2)
	fmt.Printf("S3:  %10d\n", cpu.S3)
	fmt.Printf("S4:  %10d | ", cpu.S4)
	fmt.Printf("S5:  %10d | ", cpu.S5)
	fmt.Printf("S6:  %10d | ", cpu.S6)
	fmt.Printf("S7:  %10d\n", cpu.S7)
	fmt.Printf("T8:  %10d | ", cpu.T8)
	fmt.Printf("T9:  %10d | ", cpu.T9)
	fmt.Printf("K0:  %10d | ", cpu.K0)
	fmt.Printf("K1:  %10d\n", cpu.K1)
	fmt.Printf("GP:  %10d | ", cpu.GP)
	fmt.Printf("SP:      0x%04X | ", cpu.SP)
	fmt.Printf("FP:      0x%04X | ", cpu.FP)
	fmt.Printf("RA:  %10d\n", cpu.RA)
	fmt.Printf("PC:  0x%08X | ", cpu.PC)
	fmt.Printf("HI:  %10d | ", cpu.HI)
	fmt.Printf("LO:  %10d\n", cpu.LO)
}

func (cpu *CPU) printStack() {
	fmt.Println("==================== Stack on memory =====================")
	for pc := 0xFF00; pc <= 0xFFFF; pc++ {
		if pc%16 == 0 {
			fmt.Printf("0x%04X:    ", pc)
		}
		fmt.Printf("%02X ", cpu.Memory[pc])
		if pc%16 == 15 {
			fmt.Println()
		}
	}
}
