package tokenizer

import (
	"strings"
	"regexp"
)

func Tokenize(str string) ([]*Token) {
	result := []*Token{}

	//stringParts := strings.Split(str, " ")
	stringParts := regexp.MustCompile("[\\s-_]+|[[:punct:]]+").Split(str, -1)
	for _, part := range stringParts {
		part = strings.TrimSpace(part)
		if part != "" {
			token := new(Token)
			token.value = part
			result = append(result, token)
		}
	}
	return result
}

