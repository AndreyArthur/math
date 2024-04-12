package lib

import "bytes"

type AstNode interface {
	TokenLiteral() string
	Debug() string
}

type AstExpression interface {
	AstNode
	expressionNode()
}

type AstIntegerLiteral struct {
	Token Token
	Value int64
}

func (integerLiteral *AstIntegerLiteral) expressionNode() {}
func (integerLiteral *AstIntegerLiteral) TokenLiteral() string {
	return integerLiteral.Token.Literal
}
func (integerLiteral *AstIntegerLiteral) Debug() string {
	if integerLiteral != nil {
		return integerLiteral.Token.Literal
	}
	return "?"
}

type AstPrefixExpression struct {
	Token    Token
	Operator string
	Right    AstExpression
}

func (prefixExpression *AstPrefixExpression) expressionNode() {}
func (prefixExpression *AstPrefixExpression) TokenLiteral() string {
	return prefixExpression.Token.Literal
}
func (prefixExpression *AstPrefixExpression) Debug() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(prefixExpression.Operator)
	if prefixExpression.Right != nil {
		out.WriteString(prefixExpression.Right.Debug())
	} else {
		out.WriteString("?")
	}
	out.WriteString(")")

	return out.String()
}

type AstInfixExpression struct {
	Token    Token
	Left     AstExpression
	Operator string
	Right    AstExpression
}

func (infixExpression *AstInfixExpression) expressionNode() {}
func (infixExpression *AstInfixExpression) TokenLiteral() string {
	return infixExpression.Token.Literal
}
func (infixExpression *AstInfixExpression) Debug() string {
	var out bytes.Buffer

	out.WriteString("(")
	if infixExpression.Left != nil {
		out.WriteString(infixExpression.Left.Debug())
	} else {
		out.WriteString("?")
	}
	out.WriteString(" " + infixExpression.Operator + " ")
	if infixExpression.Right != nil {
		out.WriteString(infixExpression.Right.Debug())
	} else {
		out.WriteString("?")
	}
	out.WriteString(")")

	return out.String()
}
