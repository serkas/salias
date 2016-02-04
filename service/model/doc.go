package model

import "salias/core/tokenizer"

type Doc struct {
	Text string
	Tokens []*tokenizer.Token
}

func MakeDoc(text string) *Doc {
	tokens := tokenizer.Tokenize(text)
	return &Doc{text, tokens}
}