package codewriter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"vm/mparser"
)

var ADDRESS = map[string]string {
	"local"   : "@LCL",
	"argument": "@ARG",
	"this"    : "@THIS",
	"that"    : "@THAT", 
}

type CodeWriter struct {
	fileName string
	file *os.File
	boolCount int
}

func MakeCodeWriter() *CodeWriter {
	return &CodeWriter{boolCount: 0}
}

func (c *CodeWriter) SetFileName(name string) {
	// TODO fix me. trim filename
	c.fileName = "temp"
	binFileName := strings.Replace(name, "vm", "asm", 1)
	binFile, err := os.Create(binFileName)
	if err != nil {
		panic("cannot create file ")
	}
	c.file = binFile
}

func (c *CodeWriter) WriteArithmetic(cm *mparser.Command) {
	c.writeToFile("// " + cm.Command() + " " + cm.Arg1() + " " + cm.Arg2())
	if cm.Command() != "not" && cm.Command() != "neg" {
		c.popStackToD()
	}
	c.decrementSP()
	c.setAToStack()

	switch(cm.Command()) {
	case "add":
		c.writeToFile("M=M+D")
	case "sub":
		c.writeToFile("M=M-D")
	case "or":
		c.writeToFile("M=M|D")
	case "and":
		c.writeToFile("M=M&D")
	case "eq", "gt", "lt":
		c.writeToFile("D=M-D")
		c.writeToFile(fmt.Sprintf("@BOOL%d", c.boolCount))

		switch (cm.Command()) {
		case "eq":
			c.writeToFile("D;JEQ")
		case "gt":
			c.writeToFile("D;JGT")
		case "lt":
			c.writeToFile("D;JLT")
		}

		// false
		c.setAToStack()
		c.writeToFile("M=0")
		c.writeToFile(fmt.Sprintf("@ENDBOOL%d", c.boolCount))
		c.writeToFile("0;JMP")
		
		// true
		c.writeToFile(fmt.Sprintf("(BOOL%d)", c.boolCount))
		c.setAToStack()
		c.writeToFile("M=-1")

		c.writeToFile(fmt.Sprintf("(ENDBOOL%d)", c.boolCount))
		c.boolCount++
	case "not":
		// affect top of stack
		c.writeToFile("M=-M")
	case "neg":
		// affect top of stack
		c.writeToFile("M=!M")
	}	
	c.incrementSP()
}

func (c *CodeWriter) WritePush(cm *mparser.Command) {
	c.writeToFile("// push " + cm.Arg1() + " " + cm.Arg2())
	c.setAddress(cm)
	if cm.Arg1() == "constant" {
		c.writeToFile("D=A")
	} else {
		c.writeToFile("D=M")
	}
	c.pushDToStack()
 }

func (c *CodeWriter) WritePop(cm *mparser.Command) {
	c.writeToFile("// pop " + cm.Arg1() + " " + cm.Arg2())
	c.setAddress(cm)

	// MEMORY[@R13] = @address 
	c.writeToFile("D=A")
	c.writeToFile("@R13")
	c.writeToFile("M=D")
	
	c.popStackToD()

	// MEMORY[@address] = D(from stack)
	c.writeToFile("@R13")
	c.writeToFile("A=M")
	c.writeToFile("M=D")
}

func (c *CodeWriter) Close() {
	c.file.Close()
}

func (c *CodeWriter) writeToFile(line string) {
	_, err := c.file.Write([]byte(line + "\n"))
	if err != nil {
		panic("write file failed = " + line)
	}
}

func (c *CodeWriter) setAddress(cm *mparser.Command) {
	address, _ := ADDRESS[cm.Arg1()]
	i, _ := strconv.Atoi(cm.Arg2())
	switch (cm.Arg1()) {
	case "constant":
		c.writeToFile("@" + cm.Arg2())
	case "static":
		c.writeToFile("@" + c.fileName + "." + cm.Arg2())
	case "temp":
		if i >= 8 {
			panic("temp compile err")
		}
		c.writeToFile("@R" + strconv.Itoa(5 + i))
	case "pointer":
		if i > 1{
			panic("pointer compile err")
		}
		c.writeToFile("@R" + strconv.Itoa(3 + i))
	case "local", "argment", "this", "that":
		c.writeToFile(address)
		c.writeToFile("D=M")
		c.writeToFile("@" + strconv.Itoa(i))
		// M = MEMORY[@address + x]
		c.writeToFile("A=D+A")
	}
}

func (c *CodeWriter) pushDToStack() {
	c.writeToFile("@SP")
	c.writeToFile("A=M")
	c.writeToFile("M=D")
	c.incrementSP()
}

func (c *CodeWriter) popStackToD() {
	c.decrementSP()
	c.writeToFile("A=M")
	c.writeToFile("D=M")	
}

func (c *CodeWriter) decrementSP() {
	c.writeToFile("@SP")
	c.writeToFile("M=M-1")
}

func (c *CodeWriter) incrementSP() {
	c.writeToFile("@SP")
	c.writeToFile("M=M+1")
}

func (c *CodeWriter) setAToStack() {
	c.writeToFile("@SP")
	c.writeToFile("A=M")
}