SECTIONS
{
    . = 0x0000;
    .text : { *(.text) }
    . = 0x8000;
    .data : { *(.data) }
    .bss : { *(.bss) }
    /DISCARD/ : {
        *(.reginfo)
        *(.MIPS.abiflags)
        *(.pdr)
        *(.comment)
        *(.gnu.attributes)
    }
}
