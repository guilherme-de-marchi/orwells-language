package interpreter

func SyntaxAnalisys(instructions [][]*Token) *Error {
	for _, instruc := range instructions {
		mainToken := instruc[0].Reference
		models := syntaxModels[mainToken]

		var completeSyntaxFlag bool
		var validSyntaxFlag bool

		for _, model := range models {

			validSyntaxFlag = false

			if len(instruc) >= len(model) {
				completeSyntaxFlag = true
			} else {
				completeSyntaxFlag = false
				continue
			}

			if has, anyIndex := Contains(model, ANY_TOKEN); has {

				// Quantity of tokens after <anyIndex>
				posQtty := len(model) - (anyIndex + 1)

				model = append(model[:anyIndex], model[anyIndex+1:]...)
				instruc = append(instruc[:anyIndex], instruc[len(instruc)-posQtty:]...)
			}

			for i, tokenArray := range model {

				for _, allowedToken := range tokenArray {
					if instruc[i].Reference == allowedToken {
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
			return IncompleteSyntaxError
		}

		if !validSyntaxFlag {
			return InvalidSyntaxError
		}
	}

	return nil
}

func SplitOn(tokens []*Token, target *Token) [][]*Token {
	var slices [][]*Token
	var flag int
	for i, v := range tokens {
		if v.Reference == target || i == len(tokens)-1 {
			slices = append(slices, tokens[flag:i+1])
			flag = i + 1
		}
	}

	return slices
}

type SyntaxModel [][]*Token

func Contains(tkArray [][]*Token, target *Token) (bool, int) {
	for i, allowedToken := range tkArray {
		for _, token := range allowedToken {
			if target == token {
				return true, i
			}
		}
	}

	return false, -1
}

var (
	VAR_MODEL_UNITARY = SyntaxModel{
		{VAR_TOKEN},
		{REFERENCER_TOKEN},
		{ASSIGNMENT_TOKEN},
		{REFERENCER_TOKEN, STRING_LITERAL_TOKEN, INTEGER_LITERAL_TOKEN, FLOATING_LITERAL_TOKEN},
		{INSTRUCTION_DELIMITER_TOKEN},
	}

	VAR_MODEL_BINARY = SyntaxModel{
		{VAR_TOKEN},
		{REFERENCER_TOKEN},
		{ASSIGNMENT_TOKEN},
		{REFERENCER_TOKEN, STRING_LITERAL_TOKEN, INTEGER_LITERAL_TOKEN, FLOATING_LITERAL_TOKEN},
		{SUM_TOKEN, SUBTRACTION_TOKEN},
		{REFERENCER_TOKEN, STRING_LITERAL_TOKEN, INTEGER_LITERAL_TOKEN, FLOATING_LITERAL_TOKEN},
		{INSTRUCTION_DELIMITER_TOKEN},
	}

	IF_MODEL = SyntaxModel{
		{IF_TOKEN},
		{REFERENCER_TOKEN, STRING_LITERAL_TOKEN, INTEGER_LITERAL_TOKEN, FLOATING_LITERAL_TOKEN},
		{EQUAL_TOKEN},
		{REFERENCER_TOKEN, STRING_LITERAL_TOKEN, INTEGER_LITERAL_TOKEN, FLOATING_LITERAL_TOKEN},
		{INSTRUCTION_DELIMITER_TOKEN},
	}

	ENDIF_MODEL = SyntaxModel{
		{ENDIF_TOKEN},
		{INSTRUCTION_DELIMITER_TOKEN},
	}

	EXEC_MODEL = SyntaxModel{
		{EXEC_TOKEN},
		{REFERENCER_TOKEN},
		{LEFT_PARENTESIS_TOKEN},
		{ANY_TOKEN},
		{RIGHT_PARENTESIS_TOKEN},
		{INSTRUCTION_DELIMITER_TOKEN},
	}

	EXEC_MODEL_NO_ARGUMENTS = SyntaxModel{
		{EXEC_TOKEN},
		{REFERENCER_TOKEN},
		{LEFT_PARENTESIS_TOKEN},
		{RIGHT_PARENTESIS_TOKEN},
		{INSTRUCTION_DELIMITER_TOKEN},
	}

	syntaxModels = map[*Token][]SyntaxModel{
		VAR_TOKEN:   {VAR_MODEL_UNITARY, VAR_MODEL_BINARY},
		IF_TOKEN:    {IF_MODEL},
		ENDIF_TOKEN: {ENDIF_MODEL},
		EXEC_TOKEN:  {EXEC_MODEL, EXEC_MODEL_NO_ARGUMENTS},
	}
)
