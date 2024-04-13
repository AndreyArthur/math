package lib

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	TOKEN_INTEGER          = "INTEGER"
	TOKEN_PLUS             = "+"
	TOKEN_MINUS            = "-"
	TOKEN_ASTERISK         = "*"
	TOKEN_SLASH            = "/"
	TOKEN_MODULUS          = "%"
	TOKEN_POWER            = "^"
	TOKEN_OPEN_PAREN       = "("
	TOKEN_CLOSE_PAREN      = ")"
	TOKEN_LESS             = "<"
	TOKEN_LESS_OR_EQUAL    = "<="
	TOKEN_GREATER          = ">"
	TOKEN_GREATER_OR_EQUAL = ">="
	TOKEN_EQUALS           = "="
	TOKEN_NOT_EQUALS       = "!="
	TOKEN_EOF              = "EOF"
	TOKEN_ILLEGAL          = "ILLEGAL"
)

func (token *Token) GetTypeString() string {
	types := map[TokenType]string{
		"INTEGER": "Integer",
		"+":       "Plus",
		"-":       "Minus",
		"*":       "Asterisk",
		"/":       "Slash",
		"%":       "Modulus",
		"^":       "Power",
		"(":       "Paren",
		")":       "Paren",
		"<":       "Less",
		"<=":      "Less or equal",
		">":       "Greater",
		">=":      "Greater or equal",
		"=":       "Equals",
		"!=":      "Not equals",
		"EOF":     "Eof",
		"ILLEGAL": "Illegal",
	}

	return types[token.Type]
}

func (token *Token) Debug() string {
	return fmt.Sprintf("%s %q", token.GetTypeString(), token.Literal)
}

func NewToken(tokenType TokenType, literal string) *Token {
	return &Token{
		Type:    tokenType,
		Literal: literal,
	}
}
