package tokenizer

import (
	"strings"
)


func Normalize(t Token) (Token) {
	t.value = strings.ToLower(t.value)
	return t
}
