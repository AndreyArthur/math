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
	fmt.Println("Select a mode by typing \"mode lex\" (default: lex).")
	prompt := ">> "
	scanner := bufio.NewScanner(in)
	mode := "lex"

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

		if mode == "lex" {
			lexer := lib.NewLexer(line)
			tokens := lexer.Lex()

			for _, token := range tokens {
				token.Debug()
			}
		}
	}
}

func main() {
	Repl(os.Stdin, os.Stdout)
}
