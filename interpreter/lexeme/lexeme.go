package lexeme

import (
	"strings"

	"github.com/Guilherme-De-Marchi/orwells-language/interpreter/token"
)

func GetLexemes(rawText string) []string {
	return strings.Fields(rawText) // Temporary solution
}

func TokenizeLexemes(lexemes []string) []*token.Token {
	var tkArray []*token.Token

	for i, v := range lexemes {
		var tk *token.Token

		var prevLex string
		if i > 0 {
			prevLex = lexemes[i-1]
		}

		tk = token.GetToken(v, prevLex)
		tkArray = append(tkArray, tk)
	}

	return tkArray
}
