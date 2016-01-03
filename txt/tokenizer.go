package tokenizer

import "strings"

func Tokenize(str string) ([]*Token) {
	result  := []*Token{}
	if str != "" {
		stringParts := strings.Split(str, " ")
		for _, part := range stringParts{
			token := new(Token)
			token.value = part
			result = append(result, token)
		}
	}
	return result
}
