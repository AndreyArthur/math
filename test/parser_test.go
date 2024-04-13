package test

import (
	"meth/lib"
	"testing"
)

func TestParserDefaultPrecedence(t *testing.T) {
	content := "5 - -2 / 7 ^ 3 != 70 % -5 + 8"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()

	expected := "((5 - ((-2) / (7 ^ 3))) != ((70 % (-5)) + 8))"
	got := root.Debug()

	if expected != got {
		t.Fatalf(
			"Expected ast: %q.\n               Got ast:      %q.",
			expected,
			got,
		)
	}
}

func TestParserEnforcedPrecedence(t *testing.T) {
	content := "(5 - -(-2)) / -7 ^ (3 + 2) != 70 / (-5 + 8)"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()

	expected := "(((5 - (-(-2))) / (-(7 ^ (3 + 2)))) != (70 / ((-5) + 8)))"
	got := root.Debug()

	if expected != got {
		t.Fatalf(
			"Expected ast: %q.\n               Got ast:      %q.",
			expected,
			got,
		)
	}
}

func TestParserEnforcedPrecedenceEdgeCases(t *testing.T) {
	content := "(5 - -(-2)) / (-7 ^ (3 + (2))) != (70 / (-5 + 8))"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()

	expected := "(((5 - (-(-2))) / (-(7 ^ (3 + 2)))) != (70 / ((-5) + 8)))"
	got := root.Debug()

	if expected != got {
		t.Fatalf(
			"Expected ast: %q.\n               Got ast:      %q.",
			expected,
			got,
		)
	}
}

func TestParserErrorHandling(t *testing.T) {
	content := "( + 2 (/ 7) *);"

	tokens := lib.NewLexer(content).Lex()
	parser := lib.NewParser(tokens)
	parser.Parse()

	expectedMessages := []string{
		`Opening parenthesis must be followed by numbers or minus sign. Found "+".`,
		`Opening parenthesis must be followed by numbers or minus sign. Found "/".`,
		`Closing parenthesis must be preceded by numbers. Found "*".`,
		`Found an illegal token ";".`,
		`Expected Token of type "Integer", but got token of type "Plus".`,
	}

	for index, message := range parser.GetErrors() {
		if message != expectedMessages[index] {
			t.Fatalf(
				"Expected message: %s.\n               But got:          %s",
				expectedMessages[index],
				message,
			)
		}
	}

}
