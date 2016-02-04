package tokenizer

import "testing"


func TestAscii(t *testing.T) {
	normal := Token{"lower"}
	tested := Token{"LoWer"}
	tested = Normalize(tested)

	compare(t, normal, tested)
}

func TestCyrillic(t *testing.T) {
	normal := Token{"нижний"}
	tested := Token{"НиЖний"}
	tested = Normalize(tested)

	compare(t, normal, tested)
}

func compare(t *testing.T, normal Token, tested Token)  {

	if normal.Value != tested.Value {
		t.Error("Expected", normal.Value, "got", tested.Value)
	}
}

