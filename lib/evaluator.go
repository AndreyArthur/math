package lib

type Evaluator struct {
	ast   AstExpression
	value int64
}

func isIntegerLiteral(node AstExpression) *AstIntegerLiteral {
	value, ok := interface{}(node).(*AstIntegerLiteral)
	if !ok {
		return nil
	}
	return value
}

func isPrefix(node AstExpression) *AstPrefixExpression {
	value, ok := interface{}(node).(*AstPrefixExpression)
	if !ok {
		return nil
	}
	return value
}

func isInfix(node AstExpression) *AstInfixExpression {
	value, ok := interface{}(node).(*AstInfixExpression)
	if !ok {
		return nil
	}
	return value
}

func eval(astNode AstExpression) int64 {
	integerLiteralNode := isIntegerLiteral(astNode)
	if integerLiteralNode != nil {
		return integerLiteralNode.Value
	}

	prefixNode := isPrefix(astNode)
	if prefixNode != nil {
		return -eval(prefixNode.Right)
	}

	infixNode := isInfix(astNode)
	if infixNode != nil {
		if infixNode.Operator == "+" {
			return eval(infixNode.Left) + eval(infixNode.Right)
		}
		if infixNode.Operator == "-" {
			return eval(infixNode.Left) - eval(infixNode.Right)
		}
		if infixNode.Operator == "*" {
			return eval(infixNode.Left) * eval(infixNode.Right)
		}
		if infixNode.Operator == "/" {
			return eval(infixNode.Left) / eval(infixNode.Right)
		}
		if infixNode.Operator == "^" {
			base := eval(infixNode.Left)
			power := eval(infixNode.Right)
			result := base
			var i int64
			for i = 1; i < power; i++ {
				result = result * base
			}
			return result
		}
		if infixNode.Operator == "=" {
			if eval(infixNode.Left) == eval(infixNode.Right) {
				return 1
			}
			return 0
		}
		if infixNode.Operator == "!=" {
			if eval(infixNode.Left) != eval(infixNode.Right) {
				return 1
			}
			return 0
		}
		if infixNode.Operator == ">" {
			if eval(infixNode.Left) > eval(infixNode.Right) {
				return 1
			}
			return 0
		}
		if infixNode.Operator == ">=" {
			if eval(infixNode.Left) >= eval(infixNode.Right) {
				return 1
			}
			return 0
		}
		if infixNode.Operator == "<" {
			if eval(infixNode.Left) < eval(infixNode.Right) {
				return 1
			}
			return 0
		}
		if infixNode.Operator == "<=" {
			if eval(infixNode.Left) <= eval(infixNode.Right) {
				return 1
			}
			return 0
		}
	}
	return 0
}

func (evaluator *Evaluator) Eval() int64 {
	return eval(evaluator.ast)
}

func NewEvaluator(ast AstExpression) *Evaluator {
	return &Evaluator{
		ast:   ast,
		value: 0,
	}
}
