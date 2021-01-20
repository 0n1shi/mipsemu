package mips

// R Type Instructions
func Add(cpu *CPU, ins *Instruction) error     { return nil }
func Addu(cpu *CPU, ins *Instruction) error    { return nil }
func And(cpu *CPU, ins *Instruction) error     { return nil }
func Break(cpu *CPU, ins *Instruction) error   { return nil }
func Div(cpu *CPU, ins *Instruction) error     { return nil }
func Divu(cpu *CPU, ins *Instruction) error    { return nil }
func Jalr(cpu *CPU, ins *Instruction) error    { return nil }
func Jr(cpu *CPU, ins *Instruction) error      { return nil }
func Mfhi(cpu *CPU, ins *Instruction) error    { return nil }
func Mflo(cpu *CPU, ins *Instruction) error    { return nil }
func Mthi(cpu *CPU, ins *Instruction) error    { return nil }
func Mtlo(cpu *CPU, ins *Instruction) error    { return nil }
func Mult(cpu *CPU, ins *Instruction) error    { return nil }
func Multu(cpu *CPU, ins *Instruction) error   { return nil }
func Nor(cpu *CPU, ins *Instruction) error     { return nil }
func Or(cpu *CPU, ins *Instruction) error      { return nil }
func Sll(cpu *CPU, ins *Instruction) error     { return nil }
func Sllv(cpu *CPU, ins *Instruction) error    { return nil }
func Slt(cpu *CPU, ins *Instruction) error     { return nil }
func Sltu(cpu *CPU, ins *Instruction) error    { return nil }
func Sra(cpu *CPU, ins *Instruction) error     { return nil }
func Srav(cpu *CPU, ins *Instruction) error    { return nil }
func Srl(cpu *CPU, ins *Instruction) error     { return nil }
func Srlv(cpu *CPU, ins *Instruction) error    { return nil }
func Sub(cpu *CPU, ins *Instruction) error     { return nil }
func Subu(cpu *CPU, ins *Instruction) error    { return nil }
func Syscall(cpu *CPU, ins *Instruction) error { return nil }
func Xor(cpu *CPU, ins *Instruction) error     { return nil }

// I Type Instruction
func Addi(cpu *CPU, ins *Instruction) error  { return nil }
func Addiu(cpu *CPU, ins *Instruction) error { return nil }
func Andi(cpu *CPU, ins *Instruction) error  { return nil }
func Beq(cpu *CPU, ins *Instruction) error   { return nil }
func Bgez(cpu *CPU, ins *Instruction) error  { return nil }
func Bgtz(cpu *CPU, ins *Instruction) error  { return nil }
func Blez(cpu *CPU, ins *Instruction) error  { return nil }
func Bltz(cpu *CPU, ins *Instruction) error  { return nil }
func Bne(cpu *CPU, ins *Instruction) error   { return nil }
func Lb(cpu *CPU, ins *Instruction) error    { return nil }
func Lbu(cpu *CPU, ins *Instruction) error   { return nil }
func Lh(cpu *CPU, ins *Instruction) error    { return nil }
func Lhu(cpu *CPU, ins *Instruction) error   { return nil }
func Lui(cpu *CPU, ins *Instruction) error   { return nil }
func Lw(cpu *CPU, ins *Instruction) error    { return nil }
func Lwc1(cpu *CPU, ins *Instruction) error  { return nil }
func Ori(cpu *CPU, ins *Instruction) error   { return nil }
func Sb(cpu *CPU, ins *Instruction) error    { return nil }
func Slti(cpu *CPU, ins *Instruction) error  { return nil }
func Sltiu(cpu *CPU, ins *Instruction) error { return nil }
func Sh(cpu *CPU, ins *Instruction) error    { return nil }
func Sw(cpu *CPU, ins *Instruction) error    { return nil }
func Swc1(cpu *CPU, ins *Instruction) error  { return nil }
func Xori(cpu *CPU, ins *Instruction) error  { return nil }

// J Type Instructions
func J(cpu *CPU, ins *Instruction) error   { return nil }
func Jal(cpu *CPU, ins *Instruction) error { return nil }
