package compiler

func SyntaxAnalisys(instructions [][]*Token) string {
	for _, instruc := range instructions {
		mainToken := instruc[0].Reference
		model := syntaxModels[mainToken]

		if model == nil {
			return "syntax invalid"
		}

		if len(instruc) != len(model) {
			return "incomplete syntax"
		}

		for i, tokenArray := range model {

			var flag bool
			for _, allowedToken := range tokenArray {

				if instruc[i].Reference == allowedToken {
					flag = true
					break
				}
			}

			if !flag {
				return "syntax error"
			}
		}
	}

	return "syntax ok"
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

var (
	VAR_MODEL = [][]*Token{
		{VAR_TOKEN},
		{REFERENCER_TOKEN},
		{ASSIGNMENT_TOKEN},
		{STRING_LITERAL_TOKEN, INTEGER_LITERAL_TOKEN, FLOATING_LITERAL_TOKEN},
		{INSTRUCTION_DELIMITER_TOKEN},
	}

	syntaxModels = map[*Token][][]*Token{
		VAR_TOKEN: VAR_MODEL,
	}
)
