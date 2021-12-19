package compiler

import "strings"

func LexicalAnalisys(rawText string) []*Token {
	lexemes := strings.Fields(rawText)

	var tokens []*Token
	for i, v := range lexemes {
		var token *Token

		if i == 0 {
			token = GetToken("", v)
		} else {
			token = GetToken(lexemes[i-1], v)
		}

		tokens = append(tokens, token)
	}

	return tokens
}
