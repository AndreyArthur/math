package test

import (
	"meth/lib"
	"testing"
)

func TestEvaluatorEvalNumbers(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / (16 % 6)"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 6

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalEqualsSuccess(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 = 6"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 1

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalEqualsFail(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 = 20"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 0

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalNotEqualsSuccess(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 != 7"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 1

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalNotEqualsFail(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 != 6"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 0

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalGreaterSuccess(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 > 5"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 1

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalGreaterFail(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 > 6"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 0

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalGreaterOrEqualSuccess(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 >= 6"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 1

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalGreaterOrEqualFail(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 >= 7"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 0

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalLessSuccess(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 < 7"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 1

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalLessFail(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 < 6"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 0

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalLessOrEqualSuccess(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 <= 6"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 1

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}

func TestEvaluatorEvalLessOrEqualFail(t *testing.T) {
	content := "(2 ^ (5 - 3) * 5 + 4) / 4 <= 5"

	tokens := lib.NewLexer(content).Lex()
	root := lib.NewParser(tokens).Parse()
	value := lib.NewEvaluator(root).Eval()
	expected := 0

	if value != int64(expected) {
		t.Fatalf(
			"Expected value to be: %d. Got value: %d.",
			expected,
			value,
		)
	}
}
