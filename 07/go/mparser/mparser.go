package mparser

import (
	"bufio"
	"errors"
	"os"
	"regexp"
)

type Parser struct {
	commands []*Command
	pointer int
}

func MakeParser(filePath string) (*Parser, error) {
	vmfile, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("file not exists. path = " + filePath)
	}

	scanner := bufio.NewScanner(vmfile)
	commands := []*Command {}
	commentRegex := regexp.MustCompile(`/.*`)
	for scanner.Scan() {
		// Read from vmfile once line.
		exComment := commentRegex.ReplaceAllString(scanner.Text(), "$1")
		if exComment == "" {
			continue;
		}

		c, err := MakeCommand(exComment)
		if err != nil {
			return nil, errors.New("invalid command format")
		}
		commands = append(commands, c)
	}
	
	return &Parser{
		commands: commands,
		pointer: 0,
	}, nil
}

func (p *Parser) HasMoreCommands() bool {
	return len(p.commands) > p.pointer
}

func (p *Parser) Advance()  {
	p.pointer++
}

func (p *Parser) Current() *Command {
	return p.commands[p.pointer]
}

func (p *Parser) CommandType() CommandType {
	return p.Current().CommandType()
}

func (p *Parser) Arg1() string {
	return p.Current().arg1
}

func (p *Parser) Arg2() string {
	return p.Current().arg2
}