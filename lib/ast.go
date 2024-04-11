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
	return integerLiteral.Token.Literal
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
	out.WriteString(prefixExpression.Right.Debug())
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
	out.WriteString(infixExpression.Left.Debug())
	out.WriteString(" " + infixExpression.Operator + " ")
	out.WriteString(infixExpression.Right.Debug())
	out.WriteString(")")

	return out.String()
}
