package mips

import "fmt"

// R Type Instructions
type FunctionTypeR func(cpu *CPU, rs int, rt int, rd int, sa int) error

func Add(cpu *CPU, rs int, rt int, rd int, sa int) error     { return nil }
func Addu(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func And(cpu *CPU, rs int, rt int, rd int, sa int) error     { return nil }
func Break(cpu *CPU, rs int, rt int, rd int, sa int) error   { return nil }
func Div(cpu *CPU, rs int, rt int, rd int, sa int) error     { return nil }
func Divu(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Jalr(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Jr(cpu *CPU, rs int, rt int, rd int, sa int) error      { return nil }
func Mfhi(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Mflo(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Mthi(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Mtlo(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Mult(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Multu(cpu *CPU, rs int, rt int, rd int, sa int) error   { return nil }
func Nor(cpu *CPU, rs int, rt int, rd int, sa int) error     { return nil }
func Or(cpu *CPU, rs int, rt int, rd int, sa int) error      { return nil }
func Sll(cpu *CPU, rs int, rt int, rd int, sa int) error     { return nil }
func Sllv(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Slt(cpu *CPU, rs int, rt int, rd int, sa int) error     { return nil }
func Sltu(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Sra(cpu *CPU, rs int, rt int, rd int, sa int) error     { return nil }
func Srav(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Srl(cpu *CPU, rs int, rt int, rd int, sa int) error     { return nil }
func Srlv(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Sub(cpu *CPU, rs int, rt int, rd int, sa int) error     { return nil }
func Subu(cpu *CPU, rs int, rt int, rd int, sa int) error    { return nil }
func Syscall(cpu *CPU, rs int, rt int, rd int, sa int) error { return nil }
func Xor(cpu *CPU, rs int, rt int, rd int, sa int) error     { return nil }

// I Type Instruction
type FunctionTypeI func(cpu *CPU, rs int, rt int, imm int) error

func Addi(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "addi")
		fmt.Printf("%s,%s,%d\n", registerNames[rt], registerNames[rs], imm)
	}

	*cpu.Registers[rt] = *cpu.Registers[rs] + int32(imm)
	return nil
}
func Addiu(cpu *CPU, rs int, rt int, imm int) error { return nil }
func Andi(cpu *CPU, rs int, rt int, imm int) error  { return nil }
func Beq(cpu *CPU, rs int, rt int, imm int) error   { return nil }
func Bgez(cpu *CPU, rs int, rt int, imm int) error  { return nil }
func Bgtz(cpu *CPU, rs int, rt int, imm int) error  { return nil }
func Blez(cpu *CPU, rs int, rt int, imm int) error  { return nil }
func Bltz(cpu *CPU, rs int, rt int, imm int) error  { return nil }
func Bne(cpu *CPU, rs int, rt int, imm int) error   { return nil }
func Lb(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "lb")
		fmt.Printf("%s,%d(%s)\n", registerNames[rt], int16(imm), registerNames[rs])
	}
	*cpu.Registers[rt] = int32(cpu.Memory[int(*cpu.Registers[rs])+int(int16(imm))])
	return nil
}
func Lbu(cpu *CPU, rs int, rt int, imm int) error  { return nil }
func Lh(cpu *CPU, rs int, rt int, imm int) error   { return nil }
func Lhu(cpu *CPU, rs int, rt int, imm int) error  { return nil }
func Lui(cpu *CPU, rs int, rt int, imm int) error  { return nil }
func Lw(cpu *CPU, rs int, rt int, imm int) error   { return nil }
func Lwc1(cpu *CPU, rs int, rt int, imm int) error { return nil }
func Ori(cpu *CPU, rs int, rt int, imm int) error  { return nil }
func Sb(cpu *CPU, rs int, rt int, imm int) error {
	if cpu.DebugMode {
		fmt.Printf("%-7s ", "sb")
		fmt.Printf("%s,%d(%s)\n", registerNames[rt], int16(imm), registerNames[rs])
	}
	cpu.Memory[int(*cpu.Registers[rs])+int(int16(imm))] = byte(*cpu.Registers[rt])
	return nil
}
func Slti(cpu *CPU, rs int, rt int, imm int) error       { return nil }
func Sltiu(cpu *CPU, rs int, rt int, imm int) error      { return nil }
func Sh(cpu *CPU, rs int, rt int, imm int) error         { return nil }
func Sw(cpu *CPU, rs int, rt int, imm int) error         { return nil }
func Swc1(cpu *CPU, rs int, rt int, imm int) error       { return nil }
func Xori(cpu *CPU, rs int, rt int, imm int) error       { return nil }
func DummyTypeI(cpu *CPU, rs int, rt int, imm int) error { return nil } // Dummy

// J Type Instructions
type FunctionTypeJ func(cpu *CPU, addr int) error

func J(cpu *CPU, addr int) error   { return nil }
func Jal(cpu *CPU, addr int) error { return nil }
