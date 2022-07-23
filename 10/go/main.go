package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"translator/analyzer"
)


func main()  {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		panic("no args")
	}

	dir := args[0]
	files := getFiles(dir)
	if len(files) == 0 {
		panic("no files = " + dir)
	}

	for _, f := range files {
		if !strings.Contains(f, "jack") {
			continue;
		}
		analyzer := analyzer.MakeAnalyzer(f)
		tokenizer := analyzer.GetTokenizer()
		defer analyzer.Close()

		return
		fmt.Println(tokenizer)
	}
}

func getFiles(dir string) []string {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        panic(err)
    }

    var paths []string
    for _, file := range files {
        if file.IsDir() {
            paths = append(paths, getFiles(filepath.Join(dir, file.Name()))...)
            continue
        }
        paths = append(paths, filepath.Join(dir, file.Name()))
    }

    return paths
}