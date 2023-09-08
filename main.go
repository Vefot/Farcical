package main

import (
	"farcical/evaluator"
	"farcical/lexer"
	"farcical/parser"
	"farcical/repl"
	"flag"
	"fmt"
	"io"
	"os"
	"os/user"
)

func main() {
	filepath := flag.String("file", "", "Path to file to interpret")
	flag.Parse()

	if *filepath != "" {
		code, err := os.ReadFile(*filepath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		codeString := string(code)
		l := lexer.New(codeString)
		p := parser.New(l)
		program := p.ParseProgram()
		evaluated := evaluator.Eval(program)
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
		return
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(` ______ 
|  ____|
| |__   
|  __|  
| |     
|_|
`)
	fmt.Printf("\nFarcical v0.0.0")
	fmt.Printf("\nREPL Session: %s\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
