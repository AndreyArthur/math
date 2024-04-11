package test

import (
	"meth/lib"
	"testing"
)

func TestAst(t *testing.T) {
	ast := &lib.AstInfixExpression{
		Token: *lib.NewToken(lib.TOKEN_NOT_EQUALS, "!="),
		Left: &lib.AstInfixExpression{
			Token: *lib.NewToken(lib.TOKEN_SLASH, "/"),
			Left: &lib.AstInfixExpression{
				Token: *lib.NewToken(lib.TOKEN_MINUS, "-"),
				Left: &lib.AstIntegerLiteral{
					Token: *lib.NewToken(lib.TOKEN_INTEGER, "5"),
					Value: 5,
				},
				Operator: "-",
				Right: &lib.AstPrefixExpression{
					Token:    *lib.NewToken(lib.TOKEN_MINUS, "-"),
					Operator: "-",
					Right: &lib.AstIntegerLiteral{
						Token: *lib.NewToken(lib.TOKEN_INTEGER, "2"),
						Value: 2,
					},
				},
			},
			Operator: "/",
			Right: &lib.AstInfixExpression{
				Token: *lib.NewToken(lib.TOKEN_POWER, "^"),
				Left: &lib.AstIntegerLiteral{
					Token: *lib.NewToken(lib.TOKEN_INTEGER, "7"),
					Value: 7,
				},
				Operator: "^",
				Right: &lib.AstIntegerLiteral{
					Token: *lib.NewToken(lib.TOKEN_INTEGER, "3"),
					Value: 3,
				},
			},
		},
		Operator: "!=",
		Right: &lib.AstInfixExpression{
			Token: *lib.NewToken(lib.TOKEN_PLUS, "+"),
			Left: &lib.AstInfixExpression{
				Token: *lib.NewToken(lib.TOKEN_SLASH, "/"),
				Left: &lib.AstIntegerLiteral{
					Token: *lib.NewToken(lib.TOKEN_INTEGER, "70"),
					Value: 70,
				},
				Operator: "/",
				Right: &lib.AstPrefixExpression{
					Token:    *lib.NewToken(lib.TOKEN_MINUS, "-"),
					Operator: "-",
					Right: &lib.AstIntegerLiteral{
						Token: *lib.NewToken(lib.TOKEN_INTEGER, "5"),
						Value: 5,
					},
				},
			},
			Operator: "+",
			Right: &lib.AstIntegerLiteral{
				Token: *lib.NewToken(lib.TOKEN_INTEGER, "8"),
				Value: 8,
			},
		},
	}

	expected := "(((5 - (-2)) / (7 ^ 3)) != ((70 / (-5)) + 8))"
	got := ast.Debug()

	if expected != got {
		t.Fatalf(
			"Expected ast: %q.\n            Got ast:      %q.",
			expected,
			got,
		)
	}
}
