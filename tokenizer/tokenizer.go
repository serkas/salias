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
			token := Token{part}
			token = Normalize(token)
			result = append(result, &token)
		}
	}
	return result
}

func TokenizeToStrings(str string) ([]string) {
	tokens := Tokenize(str)
	tokensStr := []string{}
	for _, token := range tokens {
		tokensStr = append(tokensStr, token.value)
	}
	return tokensStr
}

