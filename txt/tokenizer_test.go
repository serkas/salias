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

	if len(tokens) !=1 {
		t.Error("Expected one token array, got", tokens)
	}
}
