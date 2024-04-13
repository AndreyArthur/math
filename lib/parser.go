package lib

import (
	"fmt"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESS_GREATER
	SUM
	PRODUCT
	PREFIX
	POWER
)

var precedences = map[TokenType]int{
	TOKEN_EQUALS:           EQUALS,
	TOKEN_NOT_EQUALS:       EQUALS,
	TOKEN_LESS:             LESS_GREATER,
	TOKEN_LESS_OR_EQUAL:    LESS_GREATER,
	TOKEN_GREATER:          LESS_GREATER,
	TOKEN_GREATER_OR_EQUAL: LESS_GREATER,
	TOKEN_PLUS:             SUM,
	TOKEN_MINUS:            SUM,
	TOKEN_ASTERISK:         PRODUCT,
	TOKEN_SLASH:            PRODUCT,
	TOKEN_MODULUS:          PRODUCT,
	TOKEN_POWER:            POWER,
}

type Parser struct {
	tokens       []*Token
	position     int
	current      *Token
	errors       []string
	isComparison bool
}

func (parser *Parser) next() {
	if parser.current.Type != TOKEN_EOF {
		parser.position += 1
		parser.current = parser.tokens[parser.position]
	}
}

func (parser *Parser) peek(offset int) *Token {
	if parser.position+offset >= 0 &&
		parser.position+offset < len(parser.tokens) {
		return parser.tokens[parser.position+offset]
	}
	return nil
}

func (parser *Parser) expect(offset int, expectedType TokenType) bool {
	token := parser.peek(offset)
	if token.Type != expectedType {
		message := fmt.Sprintf(
			"Expected Token of type %q, but got token of type %q.",
			NewToken(expectedType, "").GetTypeString(),
			token.GetTypeString(),
		)
		parser.errors = append(parser.errors, message)
		return false
	}
	return true
}

func (parser *Parser) parseIntegerLiteral() *AstIntegerLiteral {
	if parser.expect(0, TOKEN_INTEGER) {
		value, err := strconv.ParseInt(parser.current.Literal, 0, 64)
		if err != nil {
			msg := fmt.Sprintf(
				"Could not parse %q as integer.",
				parser.current.Literal,
			)
			parser.errors = append(parser.errors, msg)
			return nil
		}
		token := *parser.current
		parser.next()
		return &AstIntegerLiteral{
			Token: token,
			Value: value,
		}
	}
	return nil
}

func (parser *Parser) parsePrefixExpression() *AstPrefixExpression {
	if parser.expect(0, TOKEN_MINUS) {
		token := *parser.current
		parser.next()
		return &AstPrefixExpression{
			Token:    token,
			Operator: token.Literal,
			Right:    parser.parseExpression(PREFIX),
		}
	}
	return nil
}

func (parser *Parser) parseInfixExpression(
	left AstExpression,
) *AstInfixExpression {
	expression := &AstInfixExpression{
		Token:    *parser.current,
		Operator: parser.current.Literal,
		Left:     left,
	}

	precedence := precedences[parser.current.Type]
	parser.next()
	expression.Right = parser.parseExpression(precedence)

	return expression
}

func (parser *Parser) parseEnforcedPrecedenceExpression() AstExpression {
	if parser.expect(0, TOKEN_OPEN_PAREN) {
		parser.next()
		expression := parser.parseExpression(LOWEST)
		parser.next()
		return expression
	}
	return nil
}

func (parser *Parser) parseExpression(precedence int) AstExpression {
	var left AstExpression

	switch parser.current.Type {
	case TOKEN_OPEN_PAREN:
		left = parser.parseEnforcedPrecedenceExpression()
	case TOKEN_MINUS:
		left = parser.parsePrefixExpression()
	default:
		left = parser.parseIntegerLiteral()
	}

	for parser.current.Type != TOKEN_EOF &&
		precedence < precedences[parser.current.Type] {
		left = parser.parseInfixExpression(left)
	}

	return left
}

func (parser *Parser) pre() {
	comparisonTokenCount := 0

	for index, token := range parser.tokens {
		if token.Literal == "=" ||
			token.Literal == "!=" ||
			token.Literal == "<" ||
			token.Literal == ">" ||
			token.Literal == "<=" ||
			token.Literal == ">=" {
			comparisonTokenCount += 1
		}

		if token.Type == TOKEN_ILLEGAL {
			message := fmt.Sprintf(
				"Found an illegal token %q.",
				token.Literal,
			)
			parser.errors = append(parser.errors, message)
		}

		if token.Type == TOKEN_OPEN_PAREN &&
			(parser.peek(index+1).Type != TOKEN_INTEGER &&
				parser.peek(index+1).Type != TOKEN_MINUS &&
				parser.peek(index+1).Type != TOKEN_OPEN_PAREN) {
			message := fmt.Sprintf(
				"Opening parenthesis must be followed by numbers or minus sign. Found %q.",
				parser.peek(index+1).Literal,
			)
			parser.errors = append(parser.errors, message)
		}

		if token.Type == TOKEN_CLOSE_PAREN &&
			(parser.peek(index-1).Type != TOKEN_INTEGER &&
				parser.peek(index-1).Type != TOKEN_CLOSE_PAREN) {
			message := fmt.Sprintf(
				"Closing parenthesis must be preceded by numbers. Found %q.",
				parser.peek(index-1).Literal,
			)
			parser.errors = append(parser.errors, message)
		}

		if token.Type == TOKEN_INTEGER &&
			parser.peek(index+1).Type == TOKEN_INTEGER {
			message := fmt.Sprintf(
				"Found a number %q followed by another number %q.",
				token.Literal,
				parser.peek(index+1).Literal,
			)
			parser.errors = append(parser.errors, message)
		}
	}

	if comparisonTokenCount > 0 {
		parser.isComparison = true
	}

	if comparisonTokenCount > 1 {
		message := "Cannot make more than one comparision per expression."
		parser.errors = append(parser.errors, message)
	}
}

func (parser *Parser) GetErrors() []string {
	return parser.errors
}

func (parser *Parser) IsComparison() bool {
	return parser.isComparison
}

func (parser *Parser) Parse() AstExpression {
	parser.pre()
	root := parser.parseExpression(LOWEST)
	return root
}

func NewParser(tokens []*Token) *Parser {
	parser := &Parser{
		tokens:   tokens,
		position: 0,
		errors:   []string{},
	}
	parser.current = parser.tokens[parser.position]
	return parser
}
