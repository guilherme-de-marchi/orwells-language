package compiler

import (
	"strconv"
)

type Token struct {
	Type, Value interface{}
	Reference   *Token
}

func NewToken(reference *Token) *Token {
	return &Token{
		Type:      reference.Type,
		Reference: reference,
	}
}

func GetToken(prevLexeme, lexeme string) *Token {
	token, ok := reservedTokens[lexeme]
	if ok {
		return NewToken(token)
	}

	vint, err := strconv.ParseInt(lexeme, 10, 64)
	if err == nil {
		token := NewToken(INTEGER_LITERAL_TOKEN)
		token.Value = vint
		return token
	}

	vfloat, err := strconv.ParseFloat(lexeme, 64)
	if err == nil {
		token := NewToken(FLOATING_LITERAL_TOKEN)
		token.Value = vfloat
		return token
	}

	if prevLexeme == VAR_KEYWORD || prevLexeme == VALUE_SIGN {
		token := NewToken(REFERENCER_TOKEN)
		token.Value = lexeme
		return token
	}

	strToken := NewToken(STRING_LITERAL_TOKEN)
	strToken.Value = lexeme
	return strToken
}

const (
	STRING_LITERAL = iota + 1
	INTEGER_LITERAL
	FLOATING_LITERAL

	REFERENCER

	// TOKEN PREDETERMINED VALUES
	RIGHT_SCOPE_DELIMITER = "}"
	LEFT_SCOPE_DELIMITER  = "{"
	INSTRUCTION_DELIMITER = ";"

	ASSIGNMENT_SIGN  = "="
	VALUE_SIGN       = "$"
	EQUAL_SIGN       = "=="
	SUM_SIGN         = "+"
	SUBTRACTION_SIGN = "-"

	IF_KEYWORD   = "if"
	VAR_KEYWORD  = "var"
	EXEC_KEYWORD = "exec"
)

var (
	// TOKENS WITH PREDETERMINED VALUES
	IF_TOKEN   = &Token{Type: IF_KEYWORD}
	VAR_TOKEN  = &Token{Type: VAR_KEYWORD}
	EXEC_TOKEN = &Token{Type: EXEC_KEYWORD}

	RIGHT_SCOPE_TOKEN           = &Token{Type: RIGHT_SCOPE_DELIMITER}
	LEFT_SCOPE_TOKEN            = &Token{Type: LEFT_SCOPE_DELIMITER}
	INSTRUCTION_DELIMITER_TOKEN = &Token{Type: INSTRUCTION_DELIMITER}

	ASSIGNMENT_TOKEN  = &Token{Type: ASSIGNMENT_SIGN}
	VALUE_TOKEN       = &Token{Type: VALUE_SIGN}
	EQUAL_TOKEN       = &Token{Type: EQUAL_SIGN}
	SUM_TOKEN         = &Token{Type: SUM_SIGN}
	SUBTRACTION_TOKEN = &Token{Type: SUBTRACTION_SIGN}

	// TOKENS WITH NOT PREDETERMINED VALUES
	STRING_LITERAL_TOKEN   = &Token{Type: STRING_LITERAL}
	INTEGER_LITERAL_TOKEN  = &Token{Type: INTEGER_LITERAL}
	FLOATING_LITERAL_TOKEN = &Token{Type: FLOATING_LITERAL}

	REFERENCER_TOKEN = &Token{Type: REFERENCER}

	reservedTokens = map[string]*Token{
		IF_KEYWORD:   IF_TOKEN,
		VAR_KEYWORD:  VAR_TOKEN,
		EXEC_KEYWORD: EXEC_TOKEN,

		RIGHT_SCOPE_DELIMITER: RIGHT_SCOPE_TOKEN,
		LEFT_SCOPE_DELIMITER:  LEFT_SCOPE_TOKEN,
		INSTRUCTION_DELIMITER: INSTRUCTION_DELIMITER_TOKEN,

		ASSIGNMENT_SIGN:  ASSIGNMENT_TOKEN,
		VALUE_SIGN:       VALUE_TOKEN,
		EQUAL_SIGN:       EQUAL_TOKEN,
		SUM_SIGN:         SUM_TOKEN,
		SUBTRACTION_SIGN: SUBTRACTION_TOKEN,
	}
)
