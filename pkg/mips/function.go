package mips

import "fmt"

// R Type Instructions
type FunctionTypeR func(cpu *CPU, rs int, rt int, rd int, sa int) error

func Add(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Addu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func And(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Break(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Div(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Divu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Jalr(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Jr(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Mfhi(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Mflo(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Mthi(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Mtlo(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Mult(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Multu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Nor(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Or(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Sll(cpu *CPU, rs int, rt int, rd int, sa int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "sll")
		fmt.Printf("%s,%s,%d\n", registerNames[rd], registerNames[rt], int16(sa))
	}

	*cpu.Registers[rd] = *cpu.Registers[rt] << int16(sa)
	return nil
}
func Sllv(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Slt(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Sltu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Sra(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Srav(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Srl(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Srlv(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Sub(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Subu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Syscall(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Xor(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}

// I Type Instruction
type FunctionTypeI func(cpu *CPU, rs int, rt int, imm int) error

func Addi(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "addi")
		fmt.Printf("%s,%s,%d\n", registerNames[rt], registerNames[rs], int16(imm))
	}

	*cpu.Registers[rt] = *cpu.Registers[rs] + int32(int16(imm))
	return nil
}
func Addiu(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Andi(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Beq(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Bgez(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Bgtz(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Blez(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Bltz(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Bne(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Lb(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "lb")
		fmt.Printf("%s,%d(%s)\n", registerNames[rt], int16(imm), registerNames[rs])
	}
	*cpu.Registers[rt] = int32(cpu.Memory[int(*cpu.Registers[rs])+int(int16(imm))])
	return nil
}
func Lbu(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Lh(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Lhu(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Lui(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Lw(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Lwc1(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Ori(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Sb(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "sb")
		fmt.Printf("%s,%d(%s)\n", registerNames[rt], int16(imm), registerNames[rs])
	}
	cpu.Memory[int(*cpu.Registers[rs])+int(int16(imm))] = byte(*cpu.Registers[rt])
	return nil
}
func Slti(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Sltiu(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Sh(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Sw(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Swc1(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Xori(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func DummyTypeI(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
} // Dummy

// J Type Instructions
type FunctionTypeJ func(cpu *CPU, addr int) error

func J(cpu *CPU, addr int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
func Jal(cpu *CPU, addr int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
}
