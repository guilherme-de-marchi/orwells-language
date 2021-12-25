package syntax

import (
	interpreterError "github.com/Guilherme-De-Marchi/orwells-language/interpreter/error"
	"github.com/Guilherme-De-Marchi/orwells-language/interpreter/token"
)

type SyntaxModel [][]*token.Token

/*
Return true if the syntax model contains the target token
and the index of its first occurence.
*/
func (sm SyntaxModel) Contains(target *token.Token) (bool, int) {
	for i, tkArray := range sm {
		for _, tk := range tkArray {
			if target == tk {
				return true, i
			}
		}
	}

	return false, -1
}

// Validates the syntax of a TokenTable.
func ValidateSyntax(tkTable token.TokenTable) *interpreterError.Error {
	headerToken := token.IdMap[tkTable[0].Id]

	var completeSyntaxFlag bool
	var validSyntaxFlag bool

	for _, sModel := range SyntaxModels[headerToken] {

		validSyntaxFlag = false

		if len(tkTable) >= len(sModel) {
			completeSyntaxFlag = true
		} else {
			completeSyntaxFlag = false
			continue
		}

		if has, index := sModel.Contains(token.GENERAL_TOKEN); has {
			after := len(sModel) - (index + 1) // Number of tokens after <GENERAL_TOKEN>

			// Remove <GENERAL_TOKEN> from syntax model on the current scope
			sModel = append(sModel[:index], sModel[index+1:]...)

			// Remove any token from token table on the current scope,
			// leaving only tokens that have their respective tokens in syntax model
			tkTable = append(tkTable[:index], tkTable[len(tkTable)-after:]...)
		}

		for i, tkArray := range sModel {

			for _, tk := range tkArray {
				if tkTable[i].Id == tk.Id {
					validSyntaxFlag = true
					break
				} else {
					validSyntaxFlag = false
				}
			}
		}

		if validSyntaxFlag {
			break
		}
	}

	if !completeSyntaxFlag {
		return interpreterError.IncompleteSyntaxError
	}

	if !validSyntaxFlag {
		return interpreterError.InvalidSyntaxError
	}

	return nil
}
