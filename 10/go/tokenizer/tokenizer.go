package tokenizer

import (
	"bufio"
	"os"
)

type Tokenizer struct {
	tokens []*Token
	pointer int
}

func MakeTokenizer(file *os.File) *Tokenizer {
	scanner := bufio.NewScanner(file)

	tokens := []*Token {}
	for scanner.Scan() {
		current := scanner.Text()
		
		for i:= 0; i < len(current); i++ {
			switch (current[i]) {
			case ' ', '\t', '\n': // skip
				continue;
			case '/':
				i = len(current)
				break;
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				buff := ""
				for current[i] >= 48 && current[i] <= 57 {
					buff = buff + string(current[i])
					i++
				}
				tokens = append(tokens, MakeToken(buff))
				continue;
			case '{','}','(',')','[',']','.',',',';','+','-','*','&','|','<','>','=', '~':
				tokens = append(tokens, MakeToken(string(current[i])))
			}

			done := false
			buff := ""
			for i < len(current) && !done {
				buff = buff + string(current[i])
				for _, v := range KEYWORD {
					if v != buff {
						continue
					}
					done = true
				}
				i++
			}
			tokens = append(tokens, MakeToken(buff))
		}
	}

	return &Tokenizer{tokens: tokens, pointer: 0}
}

func (t *Tokenizer) HasMoreToken() bool {
	return false
}

func (t *Tokenizer) Advance() {
	
}

func (t *Tokenizer) Current() *Token {
	return t.tokens[t.pointer]
}

func (t *Tokenizer) TokenType() TokenType {
	return t.Current().TokenType()
}

func (t *Tokenizer) Identifier() string {
	return t.Current().Identifier()
}

func (t *Tokenizer) IntVal() int {
	return t.Current().IntVal()
}

func (t *Tokenizer) StringVal() string {
	return t.Current().StringVal()
}

