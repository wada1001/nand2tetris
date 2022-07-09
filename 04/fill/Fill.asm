// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.


(LOOP)
    @KBD
    D=M

    @ON
    D;JGT

    @OFF
    0;JMP

(ON)
    @R0
    M=-1 // 1111111111111111
    @DRAW
    0;JMP

(OFF)
    @R0
    M=0
    @DRAW
    0;JMP

(DRAW)
    @8191 // 516 * 256 / 16
    D=A // D = 8191
    @R1
    M=D

(NEXT)
    @R1
    D=M
    @pos
    M=D
    @SCREEN
    D=A
    @pos
    M=M+D // 16384 + start pos

    @R0
    D=M
    @pos
    A=M
    M=D

    // shift 1 (16bit)
    @R1
    D=M-1
    M=D

    // 8191 - count >= 0.
    @NEXT
    D;JGE

@LOOP
0;JMP