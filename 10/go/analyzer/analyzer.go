package analyzer

import (
	"os"
	"strings"
	"translator/tokenizer"
)

type Analyzer struct {
	source *os.File
	output *os.File
}

func MakeAnalyzer(filepath string) *Analyzer {
	file, err := os.Open(filepath)
	if err != nil {
		panic("file not exists = " + filepath)
	}

	// ex MainT1.xml
	outFileName := strings.Replace(filepath, ".jack", "T1.xml", 1)
	output, err := os.Create(outFileName)
	if err != nil {
		panic("file cannot create = " + outFileName)
	}

	return &Analyzer{source: file, output: output}
}

func (a *Analyzer) GetTokenizer() *tokenizer.Tokenizer {
	return tokenizer.MakeTokenizer(a.source)
}

func (a *Analyzer) Close() {
	a.source.Close()
	a.output.Close()
}