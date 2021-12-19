package compiler

import (
	"strconv"
	"strings"
)

type Token struct {
	Group, Value interface{}
}

func GetToken(prevLexeme, lexeme string) *Token {
	token, ok := reservedTokens[lexeme]
	if ok {
		return token
	}

	vint, err := strconv.ParseInt(lexeme, 10, 64)
	if err == nil {
		token := INTEGER_LITERAL_TOKEN
		token.Value = vint
		return token
	}

	vfloat, err := strconv.ParseFloat(lexeme, 64)
	if err == nil {
		token := FLOATING_LITERAL_TOKEN
		token.Value = vfloat
		return token
	}

	vFields := strings.Fields(lexeme)

	if len(vFields) == 1 && prevLexeme == VAR_KEYWORD {
		token := REFERENCER_TOKEN
		token.Value = lexeme
		return token
	}

	strToken := STRING_LITERAL_TOKEN
	strToken.Value = lexeme
	return strToken
}

const (
	// TOKEN GROUPS
	KEYWORD = iota + 1

	STRING_LITERAL
	INTEGER_LITERAL
	FLOATING_LITERAL

	REFERENCER

	LOGIC_SIGN
	MATH_SIGN

	ASSIGNMENT
	DELIMITER

	// TOKEN PREDETERMINED VALUES
	BRACKET_DELIMITER     = "\""
	RIGHT_SCOPE_DELIMITER = "}"
	LEFT_SCOPE_DELIMITER  = "{"

	ASSIGNMENT_SIGN  = "="
	EQUAL_SIGN       = "=="
	SUM_SIGN         = "+"
	SUBTRACTION_SIGN = "-"

	IF_KEYWORD   = "if"
	VAR_KEYWORD  = "var"
	EXEC_KEYWORD = "exec"
)

var (
	// TOKENS WITH PREDETERMINED VALUES
	IF_TOKEN   = &Token{KEYWORD, IF_KEYWORD}
	VAR_TOKEN  = &Token{KEYWORD, VAR_KEYWORD}
	EXEC_TOKEN = &Token{KEYWORD, EXEC_KEYWORD}

	BRACKET_TOKEN     = &Token{DELIMITER, BRACKET_DELIMITER}
	RIGHT_SCOPE_TOKEN = &Token{DELIMITER, RIGHT_SCOPE_DELIMITER}
	LEFT_SCOPE_TOKEN  = &Token{DELIMITER, LEFT_SCOPE_DELIMITER}

	ASSIGNMENT_TOKEN  = &Token{ASSIGNMENT, ASSIGNMENT_SIGN}
	EQUAL_TOKEN       = &Token{LOGIC_SIGN, EQUAL_SIGN}
	SUM_TOKEN         = &Token{MATH_SIGN, SUM_SIGN}
	SUBTRACTION_TOKEN = &Token{MATH_SIGN, SUBTRACTION_SIGN}

	// TOKENS WITH NOT PREDETERMINED VALUES
	STRING_LITERAL_TOKEN   = &Token{Group: STRING_LITERAL}
	INTEGER_LITERAL_TOKEN  = &Token{Group: INTEGER_LITERAL}
	FLOATING_LITERAL_TOKEN = &Token{Group: FLOATING_LITERAL}

	REFERENCER_TOKEN = &Token{Group: REFERENCER}

	reservedTokens = map[string]*Token{
		IF_KEYWORD:   IF_TOKEN,
		VAR_KEYWORD:  VAR_TOKEN,
		EXEC_KEYWORD: EXEC_TOKEN,

		BRACKET_DELIMITER:     BRACKET_TOKEN,
		RIGHT_SCOPE_DELIMITER: RIGHT_SCOPE_TOKEN,
		LEFT_SCOPE_DELIMITER:  LEFT_SCOPE_TOKEN,

		ASSIGNMENT_SIGN:  ASSIGNMENT_TOKEN,
		EQUAL_SIGN:       EQUAL_TOKEN,
		SUM_SIGN:         SUM_TOKEN,
		SUBTRACTION_SIGN: SUBTRACTION_TOKEN,
	}
)
