package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"meth/lib"
	"os"
	"strings"
)

func Repl(in io.Reader, out io.Writer) {
	fmt.Println("Welcome to the Meth language REPL!")
	mode := "eval"
	fmt.Printf(
		"Select a mode by typing \"mode lex | parse | eval\" (default: %s).\n",
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

		if strings.HasPrefix(line, "mode ev") {
			mode = "eval"
			fmt.Println("Switched to the eval mode.")
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
			errors := parser.GetErrors()

			if len(errors) > 0 {
				for _, message := range errors {
					fmt.Println("Error:", message)
				}
			} else {
				fmt.Println(root.Debug())
			}
		} else {
			lexer := lib.NewLexer(line)
			tokens := lexer.Lex()
			parser := lib.NewParser(tokens)
			root := parser.Parse()
			evaluator := lib.NewEvaluator(root)
			errors := parser.GetErrors()

			if len(errors) > 0 {
				for _, message := range errors {
					fmt.Println("Error:", message)
				}
			} else {
				value := evaluator.Eval()

				if parser.IsComparison() {
					if value == 1 {
						fmt.Println("true")
					} else {
						fmt.Println("false")
					}
				} else if value == int64(math.Inf(1)) {
					fmt.Println(value, "(Infinity)")
				} else {
					fmt.Println(value)
				}
			}
		}
	}
}

func main() {
	Repl(os.Stdin, os.Stdout)
}
