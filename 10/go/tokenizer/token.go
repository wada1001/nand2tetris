package tokenizer

import "fmt"

type TokenType int

const (
	KeyWord TokenType = iota
	Symbol
	Identifier
	IntConst
	StringConst
)

type KeyWordType int

const (
	Class KeyWordType = iota
	Method
	Function
	Constructor
	Int
	Boolean
	Char
	Void
	Var
	Static
	Field
	Let
	Do
	If
	Else
	While
	Return
	True
	False
	Null
	This
)

var KEYWORD = []string {
	"class",
	"constructor",
	"function",
	"method",
	"int",
	"char",
	"boolean",
	"void",
	"true",
	"false",
	"null",
	"this",
	"let",
	"do",
	"if",
	"else",
	"while",
	"return",
}

var SYMBOL = []string {
	"{","}","(",")","[","]",
	".",",",";","+","-","*",
	"/","&","|","<",">","=","~",
}

type Token struct {
	token string
	tokenType TokenType
}

func MakeToken(token string) *Token {
	fmt.Println(token)
	return &Token{token: token}
}

func (t *Token) TokenType() TokenType {
	return StringConst
}

func (t *Token) Identifier() string {
	return ""
}

func (t *Token) IntVal() int {
	return 0
}

func (t *Token) StringVal() string {
	return ""
}

