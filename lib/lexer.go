package lib

type Lexer struct {
	content  string
	length   int
	position int
	current  byte
}

func (lexer *Lexer) advance() {
	if lexer.position >= 0 && lexer.position < lexer.length-1 {
		lexer.position += 1
		lexer.current = lexer.content[lexer.position]
	} else if lexer.position == lexer.length-1 {
		lexer.position += 1
		lexer.current = 0
	}
}

func (lexer *Lexer) peek() byte {
	if lexer.position+1 >= 0 && lexer.position+1 < lexer.length {
		return lexer.content[lexer.position+1]
	}
	return 0
}

func (lexer *Lexer) skipWhitespaces() {
	for lexer.current == ' ' || lexer.current == '\n' || lexer.current == '\t' {
		lexer.advance()
	}
}

func (lexer *Lexer) collectInteger() *Token {
	var integerLiteral string
	for lexer.current >= '0' && lexer.current <= '9' {
		integerLiteral = integerLiteral + string(lexer.current)
		lexer.advance()
	}
	return NewToken(TOKEN_INTEGER, integerLiteral)
}

func (lexer *Lexer) next() *Token {
	lexer.skipWhitespaces()

	switch lexer.current {
	case 0:
		token := NewToken(TOKEN_EOF, "")
		return token
	case '+':
		token := NewToken(TOKEN_PLUS, string(lexer.current))
		lexer.advance()
		return token
	case '-':
		token := NewToken(TOKEN_MINUS, string(lexer.current))
		lexer.advance()
		return token
	case '*':
		token := NewToken(TOKEN_ASTERISK, string(lexer.current))
		lexer.advance()
		return token
	case '/':
		token := NewToken(TOKEN_SLASH, string(lexer.current))
		lexer.advance()
		return token
	case '^':
		token := NewToken(TOKEN_POWER, string(lexer.current))
		lexer.advance()
		return token
	case '(':
		token := NewToken(TOKEN_OPEN_PAREN, string(lexer.current))
		lexer.advance()
		return token
	case ')':
		token := NewToken(TOKEN_CLOSE_PAREN, string(lexer.current))
		lexer.advance()
		return token
	case '<':
		if lexer.peek() == '=' {
			first := lexer.current
			lexer.advance()
			token := NewToken(TOKEN_LESS_OR_EQUAL, string(first)+string(lexer.current))
			lexer.advance()
			return token
		}
		token := NewToken(TOKEN_LESS, string(lexer.current))
		lexer.advance()
		return token
	case '>':
		if lexer.peek() == '=' {
			first := lexer.current
			lexer.advance()
			token := NewToken(TOKEN_GREATER_OR_EQUAL, string(first)+string(lexer.current))
			lexer.advance()
			return token
		}
		token := NewToken(TOKEN_GREATER, string(lexer.current))
		lexer.advance()
		return token
	case '=':
		token := NewToken(TOKEN_EQUALS, string(lexer.current))
		lexer.advance()
		return token
	default:
		if lexer.current == '!' && lexer.peek() == '=' {
			first := lexer.current
			lexer.advance()
			token := NewToken(TOKEN_NOT_EQUALS, string(first)+string(lexer.current))
			lexer.advance()
			return token
		}
		if lexer.current >= '0' && lexer.current <= '9' {
			return lexer.collectInteger()
		}
		token := NewToken(TOKEN_ILLEGAL, string(lexer.current))
		lexer.advance()
		return token
	}
}

func (lexer *Lexer) Lex() []*Token {
	var tokens []*Token
	token := lexer.next()
	tokens = append(tokens, token)

	for token.Type != TOKEN_EOF {
		token = lexer.next()
		tokens = append(tokens, token)
	}

	return tokens
}

func NewLexer(content string) *Lexer {
	lexer := &Lexer{
		content: content,
	}
	lexer.position = 0
	lexer.current = lexer.content[lexer.position]
	lexer.length = len(lexer.content)
	return lexer
}
