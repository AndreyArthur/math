package test

import (
	"meth/lib"
	"testing"
)

func TestExpressionLexing(t *testing.T) {
	input := "(5 \t- 2) * 7 + 5 \n / 2 ^ 3"

	table := []struct {
		tokenType lib.TokenType
		literal   string
	}{
		{lib.TOKEN_OPEN_PAREN, "("},
		{lib.TOKEN_INTEGER, "5"},
		{lib.TOKEN_MINUS, "-"},
		{lib.TOKEN_INTEGER, "2"},
		{lib.TOKEN_CLOSE_PAREN, ")"},
		{lib.TOKEN_ASTERISK, "*"},
		{lib.TOKEN_INTEGER, "7"},
		{lib.TOKEN_PLUS, "+"},
		{lib.TOKEN_INTEGER, "5"},
		{lib.TOKEN_SLASH, "/"},
		{lib.TOKEN_INTEGER, "2"},
		{lib.TOKEN_POWER, "^"},
		{lib.TOKEN_INTEGER, "3"},
		{lib.TOKEN_EOF, ""},
	}

	lexer := lib.NewLexer(input)
	tokens := lexer.Lex()

	for index, expected := range table {
		token := tokens[index]

		if token.Type != expected.tokenType {
			t.Fatalf(
				"Token[%d] - Literal wrong. Expected %q, got %q.",
				index,
				expected.tokenType,
				token.Type,
			)
		}

		if token.Literal != expected.literal {
			t.Fatalf(
				"Token[%d] - Literal wrong. Expected %q, got %q.",
				index,
				expected.literal,
				token.Literal,
			)
		}
	}
}

func TestComparisonLexing(t *testing.T) {
	input := `5 + 2 > 3
  7 * 2 >= 14
  8 / 4 < 1 * 3
  8 / 2 <= 2 * 2
  3 ^ 3 = 9 * 3
  6 - 5 != -1`

	table := []struct {
		tokenType lib.TokenType
		literal   string
	}{
		{lib.TOKEN_INTEGER, "5"},
		{lib.TOKEN_PLUS, "+"},
		{lib.TOKEN_INTEGER, "2"},
		{lib.TOKEN_GREATER, ">"},
		{lib.TOKEN_INTEGER, "3"},
		{lib.TOKEN_INTEGER, "7"},
		{lib.TOKEN_ASTERISK, "*"},
		{lib.TOKEN_INTEGER, "2"},
		{lib.TOKEN_GREATER_OR_EQUAL, ">="},
		{lib.TOKEN_INTEGER, "14"},
		{lib.TOKEN_INTEGER, "8"},
		{lib.TOKEN_SLASH, "/"},
		{lib.TOKEN_INTEGER, "4"},
		{lib.TOKEN_LESS, "<"},
		{lib.TOKEN_INTEGER, "1"},
		{lib.TOKEN_ASTERISK, "*"},
		{lib.TOKEN_INTEGER, "3"},
		{lib.TOKEN_INTEGER, "8"},
		{lib.TOKEN_SLASH, "/"},
		{lib.TOKEN_INTEGER, "2"},
		{lib.TOKEN_LESS_OR_EQUAL, "<="},
		{lib.TOKEN_INTEGER, "2"},
		{lib.TOKEN_ASTERISK, "*"},
		{lib.TOKEN_INTEGER, "2"},
		{lib.TOKEN_INTEGER, "3"},
		{lib.TOKEN_POWER, "^"},
		{lib.TOKEN_INTEGER, "3"},
		{lib.TOKEN_EQUALS, "="},
		{lib.TOKEN_INTEGER, "9"},
		{lib.TOKEN_ASTERISK, "*"},
		{lib.TOKEN_INTEGER, "3"},
		{lib.TOKEN_INTEGER, "6"},
		{lib.TOKEN_MINUS, "-"},
		{lib.TOKEN_INTEGER, "5"},
		{lib.TOKEN_NOT_EQUALS, "!="},
		{lib.TOKEN_MINUS, "-"},
		{lib.TOKEN_INTEGER, "1"},
		{lib.TOKEN_EOF, ""},
	}

	lexer := lib.NewLexer(input)
	tokens := lexer.Lex()

	for index, expected := range table {
		token := tokens[index]

		if token.Type != expected.tokenType {
			t.Fatalf(
				"Token[%d] - Literal wrong. Expected %q, got %q.",
				index,
				expected.tokenType,
				token.Type,
			)
		}

		if token.Literal != expected.literal {
			t.Fatalf(
				"Token[%d] - Literal wrong. Expected %q, got %q.",
				index,
				expected.literal,
				token.Literal,
			)
		}
	}
}

func TestIllegalLexing(t *testing.T) {
	input := `8a / 7b;`

	table := []struct {
		tokenType lib.TokenType
		literal   string
	}{
		{lib.TOKEN_INTEGER, "8"},
		{lib.TOKEN_ILLEGAL, "a"},
		{lib.TOKEN_SLASH, "/"},
		{lib.TOKEN_INTEGER, "7"},
		{lib.TOKEN_ILLEGAL, "b"},
		{lib.TOKEN_ILLEGAL, ";"},
		{lib.TOKEN_EOF, ""},
	}

	lexer := lib.NewLexer(input)
	tokens := lexer.Lex()

	for index, expected := range table {
		token := tokens[index]

		if token.Type != expected.tokenType {
			t.Fatalf(
				"Token[%d] - Literal wrong. Expected %q, got %q.",
				index,
				expected.tokenType,
				token.Type,
			)
		}

		if token.Literal != expected.literal {
			t.Fatalf(
				"Token[%d] - Literal wrong. Expected %q, got %q.",
				index,
				expected.literal,
				token.Literal,
			)
		}
	}
}
