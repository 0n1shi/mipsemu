# MIPS

## Registers

It consists of the 32-bit wide program counter (PC), and a bank of 32 general-purpose registers called r0..r31, each of which is 32-bit wide. All general-purpose registers can be used as the target registers and data sources for all logical, arithmetical, memory access, and control-flow instructions. Only r0 is special because it is internally hardwired to zero. Reading r0 always returns the value 0x00000000, and a value written to r0 is ignored and lost.

| register | assembly name | Comment                      |
| -------- | ------------- | ---------------------------- |
| r0       | $zero         | Always 0                     |
| r1       | $at           | Reserved for assembler       |
| r2-r3    | $v0-$v1       | Stores results               |
| r4-r7    | $a0-$a3       | Stores arguments             |
| r8-r15   | $t0-$t7       | Temporaries, not saved       |
| r16-r23  | $s0-$s7       | Contents saved for use later |
| r24-r25  | $t8-$t9       | More temporaries, not saved  |
| r26-r27  | $k0-$k1       | Reserved by operating system |
| r28      | $gp           | Global pointer               |
| r29      | $sp           | Stack pointer                |
| r30      | $fp           | Frame pointer                |
| r31      | $ra           | Return address               |
