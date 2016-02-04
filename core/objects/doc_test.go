package objects

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

func TestCountTokens(t *testing.T) {
	test := "aa bb cc bb"
	doc := MakeDoc(test)
	counts := doc.CountTokens()

	if counts["bb"] != 2 {
		t.Error("There are 2 `bb` tokens in string `", test, "`, got", counts["bb"])
	}

}