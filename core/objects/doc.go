package objects

import "salias/core/tokenizer"

type Doc struct {
	Text string
	Tokens []*tokenizer.Token
}

func MakeDoc(text string) *Doc {
	tokens := tokenizer.Tokenize(text)
	return &Doc{text, tokens}
}

func (doc *Doc) CountTokens() map[string]int  {
	counts := map[string]int{}

	for _, t := range doc.Tokens {
		counts[t.Value] += 1
	}
	return counts
}