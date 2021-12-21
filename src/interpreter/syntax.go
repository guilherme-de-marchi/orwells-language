package interpreter

func SyntaxAnalisys(instructions [][]*Token) *Error {
	for _, instruc := range instructions {
		mainToken := instruc[0].Reference
		models := syntaxModels[mainToken]

		var completeSyntaxFlag bool
		var validSyntaxFlag bool

		for _, model := range models {

			if len(instruc) != len(model) {
				continue
			} else {
				completeSyntaxFlag = true
			}

			validSyntaxFlag = false

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

var (
	VAR_MODEL = SyntaxModel{
		{VAR_TOKEN},
		{REFERENCER_TOKEN},
		{ASSIGNMENT_TOKEN},
		{REFERENCER_TOKEN, STRING_LITERAL_TOKEN, INTEGER_LITERAL_TOKEN, FLOATING_LITERAL_TOKEN},
		{INSTRUCTION_DELIMITER_TOKEN},
	}

	VAR_MODEL_2 = SyntaxModel{
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
	}

	syntaxModels = map[*Token][]SyntaxModel{
		VAR_TOKEN: {VAR_MODEL, VAR_MODEL_2},
		IF_TOKEN:  {IF_MODEL},
	}
)
