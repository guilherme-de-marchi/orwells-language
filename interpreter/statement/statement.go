package statement

import "github.com/Guilherme-De-Marchi/orwells-language/interpreter/token"

type Statement struct {
	TabAddition int
}

type Variable struct {
	Identifier, Value *token.Token

	Statement
}

func NewVariable(identifier, value *token.Token) *Variable {
	return &Variable{
		identifier,
		value,
		Statement{TabAddition: 0},
	}
}
