package token

import (
	"strconv"
)

type Token struct {
	Id    int
	Value interface{}
}

func NewToken(reference *Token) *Token {
	tk := *reference
	return &tk
}

func GetToken(lex, prevLex string) *Token {
	// Keyword
	if ref, ok := KeywordMap[lex]; ok {
		return NewToken(ref)
	}

	// Literal integer
	if v, err := strconv.ParseInt(lex, 10, 64); err == nil {
		tk := NewToken(INTEGER_LITERAL_TOKEN)
		tk.Value = v
		return tk
	}

	// Literal floating
	if v, err := strconv.ParseFloat(lex, 64); err == nil {
		tk := NewToken(FLOATING_LITERAL_TOKEN)
		tk.Value = v
		return tk
	}

	// Declared identifier
	if prevLex == VAR_KEYWORD {
		tk := NewToken(IDENTIFIER_TOKEN)
		tk.Value = lex
		return tk
	}

	// Referenced identifier
	if lex[:1] == IDENTIFIER_PREFIX {
		tk := NewToken(IDENTIFIER_TOKEN)
		tk.Value = lex[1:]
		return tk
	}

	// String literal
	tk := NewToken(STRING_LITERAL_TOKEN)
	tk.Value = lex
	return tk
}
