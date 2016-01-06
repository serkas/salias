package tokenizer

import "testing"

func TestEmpty(t *testing.T) {
	var tokens = Tokenize("")
	if len(tokens) != 0 {
		t.Error("Expected empty tokens array, got", tokens)
	}
}

func TestOne(t *testing.T) {
	var tokens = Tokenize("easy")

	if len(tokens) != 1 {
		t.Error("Expected one token array, got", tokens)
	}
}


func TestSpaces(t *testing.T) {
	var tokens = Tokenize(" токен    ")

	if len(tokens) != 1 {
		t.Error("Expected one token array, got", len(tokens))
		return
	}

	testToken := tokens[0]
	if testToken.value != "токен" {
		t.Error("Expected token `токен`, got", testToken.value)
	}
}

func TestPunctuation(t *testing.T) {
	test := "(проект). Однако, этап - решение!"
	var tokens = Tokenize(test)
	
	if len(tokens) != 4 {
		t.Error("Expected 4 tokens array for `", test, "` got", len(tokens))
		return
	}

	need := []*Token{&Token{"проект"}, &Token{"Однако"}, &Token{"этап"}, &Token{"решение"}}
	if !compareTokens(t, tokens, need) {
		t.Error("Tokens missmatch")
	}
}

func compareTokens(t *testing.T, tested []*Token, need []*Token) bool {
	if len(tested) != len(need) {
		return false
	}
	for index, token := range tested {
		if token.value != need[index].value {
			t.Error(token.value, "!=", need[index].value)
			return false
		}
	}
	return true
}