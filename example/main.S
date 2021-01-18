ADDI $s0, $zero, 10
ADDI $s1, $zero, 9
SB $s0, -10($sp)
SB $s1, -9($sp)
LB $s2, -9($sp)
ADDI $sp, $sp, -1
ADDI $sp, $sp, -2
ADDI $sp, $sp, -3
ADDI $sp, $sp, -5
ADDI $sp, $sp, -9
ADDI $sp, $sp, -14
ADDI $sp, $sp, -23
ADDI $sp, $sp, -37
LB $s3, 0($sp)
