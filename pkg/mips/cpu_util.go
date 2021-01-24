package mips

import "fmt"

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
		fmt.Println()
	}
}
