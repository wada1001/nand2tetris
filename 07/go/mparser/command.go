package mparser

import (
	"errors"
	"strings"
)

type CommandType int 

const (
	Arithmetic CommandType = iota
	Push
	Pop
	Label
	Goto
	If
	Function
	Return
	Call 
)

type Command struct {
	command string
	arg1 string
	arg2 string
}

func MakeCommand(commandStr string) (*Command, error) {
	if commandStr == "" {
		return nil, errors.New("empty command.")
	}

	arr := strings.Split(commandStr, ` `)
	command := ""
	arg1 := ""
	arg2 := ""
	if len(arr) >= 3 {
		arg2 = strings.TrimSpace(arr[2])
	}
	if len(arr) >= 2 {
		arg1 = strings.TrimSpace(arr[1])
	}
	if len(arr) >= 1 {
		command = strings.TrimSpace(arr[0])
	}

	return &Command{
		command: command,
		arg1: arg1,
		arg2: arg2,
	}, nil
}

func (c *Command) CommandType() CommandType {
	switch (c.command) {
	// case "add":
	// case "sub":
	// case "neg":
	// case "eq":
	// case "gt":
	// case "lt":
	// case "and":
	// case "or":
	// case "not":
	// 	return Arithmetic
	case "push":
		return Push
	case "pop":
		return Pop
	case "label":
		return Label
	case "goto":
		return Goto
	case "if-goto":
		return If
	case "function":
		return Function
	case "return":
		return Return
	case "call":
		return Call
	default:
		return Arithmetic
	}
}

func (c *Command) Command() string {
	return c.command
}

func (c *Command) Arg1() string {
	return c.arg1
}

func (c *Command) Arg2() string {
	return c.arg2
}
