package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	fileName := args[0]

	asmfile, err := os.Open(fileName)
	if err != nil {
		panic("file cannot open")
	}

	// Create .hack
	binFileName := strings.Replace(fileName, "asm", "hack", 1)
	binFile, err := os.Create(binFileName)
	if err != nil {
		panic("cannot create file ")
	}

	defer binFile.Close()
	defer asmfile.Close()

	parser := MakeParser(asmfile)
	symbolTable := MakeSymbolTable()

	pc := 0
	for parser.HasMoreCommands() {
		if parser.CommadType() != L {
			parser.Advance()
			pc++
			continue;
		}
		symbol := parser.Symbol()
		if symbolTable.Contains(symbol) {
			parser.Advance()
			pc++
			continue;
		}
		symbolTable.AddEntry(parser.Symbol(), pc)
		parser.Advance()
		// L symbol addr is next pc (SYMBOL)
	}

	// Reset Parser
	parser.Reset()
	// not l symbol addr start at 16
	currentNotLSymbolCount := 0 + 16

	for parser.HasMoreCommands() {
		switch parser.CommadType() {
		case A:
			a := 0b0000000000000000
			symbol := parser.Symbol()
			i, err := strconv.Atoi(symbol)

			if err != nil {
				// not const val
				if symbolTable.Contains(symbol) {
					addr := symbolTable.GetAddress(symbol)
					a = a | addr
				} else {
					a = a | currentNotLSymbolCount
					symbolTable.AddEntry(parser.Symbol(), currentNotLSymbolCount)
					currentNotLSymbolCount++
				}

				_, err := binFile.Write([]byte(fmt.Sprintf("%016b\n", a)))
				if err != nil {
					panic(err)
				}
			} else {
				// const val
				a = a | i
				_, err := binFile.Write([]byte(fmt.Sprintf("%016b\n", a)))
				if err != nil {
					panic(err)
				}
			}
		case C:
			ret := Comp(parser.Comp())
			ret = ret | Dest(parser.Dest())
			ret = ret | Jump(parser.Jump())
			_, err := binFile.Write([]byte(fmt.Sprintf("%016b\n", ret)))
			if err != nil {
				panic(err)
			}
		default:
		}
		parser.Advance()
	}
}

/////////////////////////Parser//////////////////////////

type CommandType int8

const (
    A CommandType = iota // 0xxxxxxxxxxxxxxx
    C CommandType = iota // 1xxxxxxxxxxxxxxx
    L CommandType = iota // (Xxx)
)

type Parser struct {
	commands []string
	pointer int
}

func MakeParser(asmFile *os.File) *Parser {
	commands := []string {}
	commentRegex := regexp.MustCompile(`/.*`)
	scanner := bufio.NewScanner(asmFile)

	for scanner.Scan() {
		result := strings.TrimSpace(scanner.Text())
		result = commentRegex.ReplaceAllString(result, "$1")
		if result == "" {
			continue;
		}
		commands = append(commands, result)
	}
	return &Parser{commands: commands, pointer: 0}
}

func (p *Parser) HasMoreCommands() bool {
	return len(p.commands) > p.pointer 
}

func (p *Parser) Advance () {
	p.pointer++;
}

func (p *Parser) Current () string {
	return p.commands[p.pointer]
}

func (p *Parser) CommadType() CommandType {
	prefix := p.Current()[0]
	switch prefix {
	case '@':
		return A
	case '(':
		return L
	default:
		return C
	}
}

func (p *Parser) Symbol() string {
	if p.CommadType() == L {
		res := regexp.MustCompile(`\(.+?\)`).Find([]byte(p.Current()))
		return string(res[1:len(res) - 1])
	}
	if p.CommadType() == A {
		r := regexp.MustCompile("@.+")
		return string(r.Find([]byte(p.Current()))[1:])
	}
	return ""
}

func (p *Parser) Dest() string {
	c := p.Current()
	tmp := strings.Split(c, "=")
	if len(tmp) == 1{
		return ""
	}
	return strings.TrimSpace(tmp[0])
}

func (p *Parser) Comp() string {
	c := p.Current()
	ret := ""
	if strings.Contains(c, ";") {
		// contain jump
		tmp := strings.Split(c, ";")
		ret = tmp[0]
	} else {
		// not jump
		tmp := strings.Split(c, "=")
		ret = tmp[len(tmp) - 1]
	}
	return strings.TrimSpace(ret)
}

func (p *Parser) Jump() string {
	c := p.Current()
	tmp := strings.Split(c, ";")
	if len(tmp) == 1 {
		return ""
	}
	return strings.TrimSpace(tmp[1])
}

func (p *Parser) Reset(){
	p.pointer = 0
}

/////////////////////////Code//////////////////////////

var COMP = map[string]int{
	"0":   0b1110101010000000,
	"1":   0b1110111111000000,
	"-1":  0b1110111010000000,
	"D":   0b1110001100000000,
	"A":   0b1110110000000000,
	"!D":  0b1110001101000000,
	"!A":  0b1110110001000000,
	"-D":  0b1110001111000000,
	"-A":  0b1110110011000000,
	"D+1": 0b1110011111000000,
	"A+1": 0b1110110111000000,
	"D-1": 0b1110001110000000,
	"A-1": 0b1110110010000000,
	"D+A": 0b1110000010000000,
	"D-A": 0b1110010011000000,
	"A-D": 0b1110000111000000,
	"D&A": 0b1110000000000000,
	"D|A": 0b1110010101000000,
	"M":   0b1111110000000000,
	"!M":  0b1111110001000000,
	"-M":  0b1111110011000000,
	"M+1": 0b1111110111000000,
	"M-1": 0b1111110010000000,
	"D+M": 0b1111000010000000,
	"D-M": 0b1111010011000000,
	"M-D": 0b1111000111000000,
	"D&M": 0b1111000000000000,
	"D|M": 0b1111010101000000,
}

func Dest(str string) int {
	ret := 0b000000

	if strings.Contains(str, "A") {
		ret = ret | 0b100000
	}
	if strings.Contains(str, "M") {
		ret = ret | 0b001000
	}
	if strings.Contains(str, "D") {
		ret = ret | 0b010000
	}
	return ret
}

func Comp(str string) int {
	val, ok := COMP[str]
	if !ok {
		fmt.Println(str + " not found")
        return 0b0000000000000000
    }
	return val
}

func Jump(str string) int {
	switch(str) {
	case "JGT":
		return 0b000001
	case "JEQ":
		return 0b000010
	case "JGE":
		return 0b000011
	case "JLT":
		return 0b000100
	case "JNE":
		return 0b000101
	case "JLE":
		return 0b000110
	case "JMP":
		return 0b000111
	}
	return 0b000000
}

/////////////////////////SymbolTable//////////////////////////

type SymbolTable struct {
	symbolMap map[string]int
}

func MakeSymbolTable() *SymbolTable {
	symbolMap := map[string]int{"SP": 0, "LCL": 1, "ARG": 2, "THIS": 3, "THAT": 4, "SCREEN": 16384, "KBD": 24576}
	for i := 0; i < 16; i++ {
		symbolMap[fmt.Sprintf("R%d", i)] = i
	}
	return &SymbolTable{symbolMap: symbolMap}
}

func (s *SymbolTable) AddEntry(symbol string, address int) {
	s.symbolMap[symbol] = address
}

func (s *SymbolTable) Contains(symbol string) bool {
	_, ok := s.symbolMap[symbol]
	return ok
}

func (s *SymbolTable) GetAddress(symbol string) int {
	address, ok := s.symbolMap[symbol]
	if ok {
		return address
	}
	return -1
}

