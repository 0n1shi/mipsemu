package mips

import (
	"errors"
	"fmt"
)

// R Type Instructions
type FunctionTypeR func(cpu *CPU, rs int, rt int, rd int, sa int) error

func Add(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: add")
}
func Addu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: addu")
}
func And(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: and")
}
func Break(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: break")
}
func Div(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: div")
}
func Divu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: divu")

}
func Jalr(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: jalr")
}
func Jr(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: jr")
}
func Mfhi(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: mfhi")
}
func Mflo(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: mflo")
}
func Mthi(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: mthi")
}
func Mtlo(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: mtlo")
}
func Mult(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: mult")
}
func Multu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: multu")
}
func Nor(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: nor")
}
func Or(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: or")
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
	return errors.New("not implemented: sllv")
}
func Slt(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: slt")
}
func Sltu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: sltu")
}
func Sra(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: sra")
}
func Srav(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: srav")
}
func Srl(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: srl")
}
func Srlv(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: srlv")
}
func Sub(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: sub")
}
func Subu(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: subu")
}
func Syscall(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: syscall")
}
func Xor(cpu *CPU, rs int, rt int, rd int, sa int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: xor")
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
	return errors.New("not implemented: addiu")
}
func Andi(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: andi")
}
func Beq(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: beq")
}
func Bgez(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: bgez")
}
func Bgtz(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: bgtz")
}
func Blez(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: blez")
}
func Bltz(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: bltz")
}
func Bne(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: bne")
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
	return errors.New("not implemented: lbu")
}
func Lh(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: lh")
}
func Lhu(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: lhu")
}
func Lui(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: lui")
}
func Lw(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: lw")
}
func Lwc1(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: lwc1")
}
func Ori(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: ori")
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
	return errors.New("not implemented: slti")
}
func Sltiu(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: sltiu")
}
func Sh(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: sh")
}
func Sw(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: sw")
}
func Swc1(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: swc1")
}
func Xori(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: xori")
}
func DummyTypeI(cpu *CPU, rs int, rt int, imm int) error {
	fmt.Printf("*(not implemented)\n")
	return nil
} // Dummy

// J Type Instructions
type FunctionTypeJ func(cpu *CPU, addr int) error

func J(cpu *CPU, addr int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: j")
}
func Jal(cpu *CPU, addr int) error {
	fmt.Printf("*(not implemented)\n")
	return errors.New("not implemented: jal")
}
