SECTIONS
{
    . = 0x0000;
    .text : {
        *(.text)
        *(.data)
        *(.bss)
    }
    /DISCARD/ : {
        *(.reginfo)
        *(.MIPS.abiflags)
        *(.pdr)
        *(.comment)
        *(.gnu.attributes)
    }
}
