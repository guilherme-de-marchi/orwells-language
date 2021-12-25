package token

type TokenTable []*Token

func GetTokenTables(tkArray []*Token, statementDelimiter *Token) []TokenTable {
	var tkTables []TokenTable
	var flag int

	for i, tk := range tkArray {
		if IdMap[tk.Id] == statementDelimiter || i == len(tkArray)-1 {

			tkTable := TokenTable(tkArray[flag : i+1])
			tkTables = append(tkTables, tkTable)

			flag = i + 1
		}
	}

	return tkTables
}
