package main

import (
	"flag"
	"vm/codewriter"
	"vm/mparser"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fileName := args[0]

	parser, err := mparser.MakeParser(fileName)
	if err != nil {
		panic("paser cannot create. reason = " + err.Error())
	}

	codeWriter := codewriter.MakeCodeWriter()
	codeWriter.SetFileName(fileName)

	for parser.HasMoreCommands() {
		command := parser.Current()
		// fmt.Println(command)
		switch (command.CommandType()) {
		case mparser.Push:
			codeWriter.WritePush(command)
		case mparser.Pop:
			codeWriter.WritePop(command)
		case mparser.Arithmetic:
			codeWriter.WriteArithmetic(command)
		}

		
		parser.Advance()
	}

	codeWriter.Close()
}