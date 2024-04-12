package main

import (
	"bufio"
	"fmt"
	"io"
	"meth/lib"
	"os"
	"strings"
)

func Repl(in io.Reader, out io.Writer) {
	fmt.Println("Welcome to the Meth language REPL!")
	mode := "parse"
	fmt.Printf(
		"Select a mode by typing \"mode lex | parse\" (default: %s).\n",
		mode,
	)
	prompt := ">> "
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if strings.HasPrefix(line, "mode lex") {
			mode = "lex"
			fmt.Println("Switched to the lex mode.")
			continue
		}

		if strings.HasPrefix(line, "mode pars") {
			mode = "parse"
			fmt.Println("Switched to the parse mode.")
			continue
		}

		if mode == "lex" {
			lexer := lib.NewLexer(line)
			tokens := lexer.Lex()

			for _, token := range tokens {
				fmt.Println(token.Debug())
			}
		} else if mode == "parse" {
			lexer := lib.NewLexer(line)
			tokens := lexer.Lex()
			parser := lib.NewParser(tokens)
			root := parser.Parse()

			if len(parser.Errors) != 0 {
				for _, message := range parser.Errors {
					fmt.Println("Error:", message)
				}
			} else {
				fmt.Println(root.Debug())
			}
		}
	}
}

func main() {
	Repl(os.Stdin, os.Stdout)
}
