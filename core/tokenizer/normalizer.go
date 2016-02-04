package tokenizer

import (
	"strings"
)


func Normalize(t Token) (Token) {
	t.Value = strings.ToLower(t.Value)
	return t
}
