# Dump of ELF binary

## Create binary and dump it.

```asm
$ make example
docker run --rm -v /Users/kaz/ghq/github.com/0n1shi/mipsemu:/workdir multiarch/crossbuild /usr/mipsel-linux-gnu/bin/objdump -D bin/main.elf

bin/main.elf:     file format elf32-tradlittlemips

Disassembly of section .text:

00000000 <main>:
   0:   27bdffd0        addiu   sp,sp,-48
   4:   afbf002c        sw      ra,44(sp)
   8:   afbe0028        sw      s8,40(sp)
   c:   03a0f021        move    s8,sp
  10:   2402000a        li      v0,10
  14:   afc20018        sw      v0,24(s8)
  18:   24020003        li      v0,3
  1c:   afc2001c        sw      v0,28(s8)
  20:   8fc40018        lw      a0,24(s8)
  24:   8fc5001c        lw      a1,28(s8)
  28:   0c000014        jal     50 <plus>
  2c:   00200825        move    at,at
  30:   afc20020        sw      v0,32(s8)
  34:   00001021        move    v0,zero
  38:   03c0e821        move    sp,s8
  3c:   8fbf002c        lw      ra,44(sp)
  40:   8fbe0028        lw      s8,40(sp)
  44:   27bd0030        addiu   sp,sp,48
  48:   03e00008        jr      ra
  4c:   00200825        move    at,at

00000050 <plus>:
  50:   27bdfff8        addiu   sp,sp,-8
  54:   afbe0004        sw      s8,4(sp)
  58:   03a0f021        move    s8,sp
  5c:   afc40008        sw      a0,8(s8)
  60:   afc5000c        sw      a1,12(s8)
  64:   8fc30008        lw      v1,8(s8)
  68:   8fc2000c        lw      v0,12(s8)
  6c:   00200825        move    at,at
  70:   00621021        addu    v0,v1,v0
  74:   03c0e821        move    sp,s8
  78:   8fbe0004        lw      s8,4(sp)
  7c:   27bd0008        addiu   sp,sp,8
  80:   03e00008        jr      ra
  84:   00200825        move    at,at
  88:   00200825        move    at,at
  8c:   00200825        move    at,at

00000090 <a>:
  90:   0000000a        0xa
        ...
```
