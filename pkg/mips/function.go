package mips

import (
	"errors"
	"fmt"
)

// R Type Instructions
type FunctionTypeR func(cpu *CPU, rs int, rt int, rd int, sa int) error

func Add(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"add\" not implemented)\n")
	return errors.New("not implemented: add")
}
func Addu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "addu")
		fmt.Printf("%s,%s,%s\n", registerNames[rd], registerNames[rs], registerNames[rt])
	}

	*cpu.Registers[rd] = *cpu.Registers[rs] + *cpu.Registers[rt]
	return nil
}
func And(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"and\" not implemented)\n")
	return errors.New("not implemented: and")
}
func Break(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"break\" not implemented)\n")
	return errors.New("not implemented: break")
}
func Div(cpu *CPU, rs int, rt int, rd int, sa int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "div")
		fmt.Printf("%s,%s\n", registerNames[rs], registerNames[rt])
	}

	cpu.LO = *cpu.Registers[rs] / *cpu.Registers[rt]
	cpu.HI = *cpu.Registers[rs] % *cpu.Registers[rt]

	return nil
}
func Divu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"divu\" not implemented)\n")
	return errors.New("not implemented: divu")

}
func Jalr(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"jalr\" not implemented)\n")
	return errors.New("not implemented: jalr")
}

// Jump Register
const ErrorMsgForEndOfMain = "end of main function"

func Jr(cpu *CPU, rs int, rt int, rd int, sa int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "jr")
		fmt.Printf("%s\n", registerNames[rs])
	}

	if rs == 31 && cpu.SP == stackBaseAddr { // register source == ra register
		return errors.New(ErrorMsgForEndOfMain)
	}

	cpu.PC = int(*cpu.Registers[rs])
	return nil
}

// Move From HI
func Mfhi(cpu *CPU, rs int, rt int, rd int, sa int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "mfhi")
		fmt.Printf("%s\n", registerNames[rd])
	}

	*cpu.Registers[rd] = cpu.HI

	return nil
}

// Move From LO
func Mflo(cpu *CPU, rs int, rt int, rd int, sa int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "mflo")
		fmt.Printf("%s\n", registerNames[rd])
	}

	*cpu.Registers[rd] = cpu.LO

	return nil
}
func Mthi(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"mthi\" not implemented)\n")
	return errors.New("not implemented: mthi")
}
func Mtlo(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"mtlo\" not implemented)\n")
	return errors.New("not implemented: mtlo")
}

// Multiply
func Mult(cpu *CPU, rs int, rt int, rd int, sa int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "mult")
		fmt.Printf("%s,%s\n", registerNames[rs], registerNames[rt])
	}

	result := int64(*cpu.Registers[rs]) * int64(*cpu.Registers[rt])
	cpu.HI = int32(result >> 32)
	cpu.LO = int32(result & 0xFFFFFFFF)

	return nil
}

func Multu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"multu\" not implemented)\n")
	return errors.New("not implemented: multu")
}
func Nor(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"nor\" not implemented)\n")
	return errors.New("not implemented: nor")
}

// Or
func Or(cpu *CPU, rs int, rt int, rd int, sa int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "or")
		fmt.Printf("%s,%s,%s\n", registerNames[rd], registerNames[rs], registerNames[rt])
	}

	*cpu.Registers[rd] = *cpu.Registers[rs] | *cpu.Registers[rt]

	return nil
}

// Shift Left Logical
func Sll(cpu *CPU, rs int, rt int, rd int, sa int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "sll")
		fmt.Printf("%s,%s,%d\n", registerNames[rd], registerNames[rt], int16(sa))
	}

	*cpu.Registers[rd] = *cpu.Registers[rt] << int16(sa)

	return nil
}

func Sllv(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"sllv\" not implemented)\n")
	return errors.New("not implemented: sllv")
}
func Slt(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"slt\" not implemented)\n")
	return errors.New("not implemented: slt")
}
func Sltu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"sltu\" not implemented)\n")
	return errors.New("not implemented: sltu")
}
func Sra(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"sra\" not implemented)\n")
	return errors.New("not implemented: sra")
}
func Srav(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"srav\" not implemented)\n")
	return errors.New("not implemented: srav")
}
func Srl(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"srl\" not implemented)\n")
	return errors.New("not implemented: srl")
}
func Srlv(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"srlv\" not implemented)\n")
	return errors.New("not implemented: srlv")
}
func Sub(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"sub\" not implemented)\n")
	return errors.New("not implemented: sub")
}

// Subtract Unsigned
func Subu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "subu")
		fmt.Printf("%s,%s,%s\n", registerNames[rd], registerNames[rs], registerNames[rt])
	}

	*cpu.Registers[rd] = *cpu.Registers[rs] - *cpu.Registers[rt]

	return nil
}
func Syscall(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"syscall\" not implemented)\n")
	return errors.New("not implemented: syscall")
}
func Xor(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("(\"xor\" not implemented)\n")
	return errors.New("not implemented: xor")
}

// I Type Instruction
type FunctionTypeI func(cpu *CPU, rs int, rt int, imm int) error

// Add Immediate
func Addi(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "addi")
		fmt.Printf("%s,%s,%d\n", registerNames[rt], registerNames[rs], int16(imm))
	}

	*cpu.Registers[rt] = *cpu.Registers[rs] + int32(int16(imm))

	return nil
}

// Add Immediate Unsigned
func Addiu(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "addiu")
		fmt.Printf("%s,%s,%d\n", registerNames[rt], registerNames[rs], int16(imm))
	}

	// even when overflow, won't throw exception
	*cpu.Registers[rt] = *cpu.Registers[rs] + int32(int16(imm))

	return nil
}
func Andi(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"andi\" not implemented)\n")
	return errors.New("not implemented: andi")
}
func Beq(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"beq\" not implemented)\n")
	return errors.New("not implemented: beq")
}
func Bgez(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"bgez\" not implemented)\n")
	return errors.New("not implemented: bgez")
}
func Bgtz(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"bgtz\" not implemented)\n")
	return errors.New("not implemented: bgtz")
}
func Blez(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"blez\" not implemented)\n")
	return errors.New("not implemented: blez")
}
func Bltz(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"bltz\" not implemented)\n")
	return errors.New("not implemented: bltz")
}

// Branch On Not Equal
func Bne(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "bne")
		fmt.Printf("%s,%s,%d\n", registerNames[rs], registerNames[rt], int16(imm)<<2)
	}

	if *cpu.Registers[rs] != *cpu.Registers[rt] {
		cpu.PC = cpu.PC + int(int16(imm))<<2
	}

	return nil
}

// Load Byte
func Lb(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "lb")
		fmt.Printf("%s,%d(%s)\n", registerNames[rt], int16(imm), registerNames[rs])
	}

	*cpu.Registers[rt] = int32(cpu.Memory[int(*cpu.Registers[rs])+int(int16(imm))])

	return nil
}
func Lbu(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"lbu\" not implemented)\n")
	return errors.New("not implemented: lbu")
}
func Lh(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"lh\" not implemented)\n")
	return errors.New("not implemented: lh")
}
func Lhu(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"lhu\" not implemented)\n")
	return errors.New("not implemented: lhu")
}

// Load Upper Immediate
func Lui(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "lui")
		fmt.Printf("%s, %04X\n", registerNames[rt], int16(imm))
	}

	*cpu.Registers[rt] = int32(imm << 16)

	return nil
}

// Load Word
func Lw(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "lw")
		fmt.Printf("%s,%d(%s)\n", registerNames[rt], int16(imm), registerNames[rs])
	}

	value := int32(cpu.Memory[int(int16(imm))+int(*cpu.Registers[rs])])
	value |= int32(cpu.Memory[int(int16(imm))+int(*cpu.Registers[rs])+1]) << 8
	value |= int32(cpu.Memory[int(int16(imm))+int(*cpu.Registers[rs])+2]) << 16
	value |= int32(cpu.Memory[int(int16(imm))+int(*cpu.Registers[rs])+3]) << 24
	*cpu.Registers[rt] = value

	return nil
}
func Lwc1(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"lwc1\" not implemented)\n")
	return errors.New("not implemented: lwc1")
}
func Ori(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"ori\" not implemented)\n")
	return errors.New("not implemented: ori")
}

// Store Byte
func Sb(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "sb")
		fmt.Printf("%s,%d(%s)\n", registerNames[rt], int16(imm), registerNames[rs])
	}

	cpu.Memory[int(int16(imm))+int(*cpu.Registers[rs])] = byte(*cpu.Registers[rt])

	return nil
}

// 	Set On Less Than Immediate
func Slti(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "slti")
		fmt.Printf("%s,%s,%d\n", registerNames[rt], registerNames[rs], int16(imm))
	}

	if *cpu.Registers[rs] < int32(int16(imm)) {
		*cpu.Registers[rt] = 1
	} else {
		*cpu.Registers[rt] = 0
	}

	return nil
}
func Sltiu(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"sltiu\" not implemented)\n")
	return errors.New("not implemented: sltiu")
}
func Sh(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"sh\" not implemented)\n")
	return errors.New("not implemented: sh")
}

// Store Word
func Sw(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "sw")
		fmt.Printf("%s,%d(%s)\n", registerNames[rt], int16(imm), registerNames[rs])
	}

	addr := int(*cpu.Registers[rs]) + int(int16(imm))
	value := *cpu.Registers[rt]
	cpu.Memory[addr] = byte(value & 0xFF)
	cpu.Memory[addr+1] = byte((value >> 8) & 0xFF)
	cpu.Memory[addr+2] = byte((value >> 16) & 0xFF)
	cpu.Memory[addr+3] = byte((value >> 24) & 0xFF)

	return nil
}
func Swc1(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"swc1\" not implemented)\n")
	return errors.New("not implemented: swc1")
}
func Xori(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(\"xori\" not implemented)\n")
	return errors.New("not implemented: xori")
}
func DummyTypeI(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("(not implemented)\n")
	return nil
} // Dummy

// J Type Instructions
type FunctionTypeJ func(cpu *CPU, addr int) error

// Jump
func J(cpu *CPU, addr int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "j")
		fmt.Printf("0x%08X\n", addr<<2)
	}

	cpu.PC = (cpu.PC & 0xFFFF0000) | (addr << 2)

	return nil
}

// Jump And Link
func Jal(cpu *CPU, addr int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "jal")
		fmt.Printf("0x%08X\n", addr<<2)
	}

	cpu.RA = int32(cpu.PC)
	cpu.PC = addr << 2

	return nil
}
