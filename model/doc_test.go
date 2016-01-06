package model

import "testing"

func TestMakeDoc(t *testing.T) {
	text := "The quick brown fox"
	doc := MakeDoc(text)

	if doc.Text != text {
		t.Error("Given text and DocText not equal")
	}

	if len(doc.Tokens) != 4 {
		t.Error("4 tokens expected for", text, "got", len(doc.Tokens))
	}
}

