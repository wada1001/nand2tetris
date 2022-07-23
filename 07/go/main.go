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
	codeWriter.WriteInit()

	for parser.HasMoreCommands() {
		command := parser.Current()
		codeWriter.WriteComment(command)

		switch (command.CommandType()) {
		case mparser.Push:
			codeWriter.WritePush(command)
		case mparser.Pop:
			codeWriter.WritePop(command)
		case mparser.Arithmetic:
			codeWriter.WriteArithmetic(command)
		case mparser.Label:
			codeWriter.WriteLabel(command)
		case mparser.If:
			codeWriter.WriteLabel(command)
		case mparser.Call:
			codeWriter.WriteCall(command)
		case mparser.Goto:
			codeWriter.WriteGoto(command)
		case mparser.Function:
			codeWriter.WriteFunction(command)
		case mparser.Return:
			codeWriter.WriteReturn(command)
		}

		parser.Advance()
	}

	codeWriter.Close()
}