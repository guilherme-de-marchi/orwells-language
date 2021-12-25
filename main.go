package main

import (
	"fmt"
	"strings"

	"github.com/Guilherme-De-Marchi/orwells-language/interpreter/lexeme"
	"github.com/Guilherme-De-Marchi/orwells-language/interpreter/statement"
	"github.com/Guilherme-De-Marchi/orwells-language/interpreter/syntax"
	"github.com/Guilherme-De-Marchi/orwells-language/interpreter/token"
	"github.com/Guilherme-De-Marchi/orwells-language/interpreter/utils"
)

func main() {
	rawText := utils.FileToString("code.orwell")
	refinedText := strings.Replace(rawText, "\r\n", "", -1)

	fmt.Println(refinedText)

	fmt.Println("\n#-#-#-#-# Lexical analisys #-#-#-#-#\n")

	lexemes := lexeme.GetLexemes(refinedText)
	tkArray := lexeme.TokenizeLexemes(lexemes)
	for _, v := range tkArray {
		fmt.Println(*v)
	}

	fmt.Println("\n#-#-#-#-# Getting token tables #-#-#-#-#\n")

	tkTables := token.GetTokenTables(tkArray, token.INSTRUCTION_DELIMITER_TOKEN)
	for _, v := range tkTables {
		fmt.Println("----- Token Table -----")
		for _, tk := range v {
			fmt.Println(*tk)
		}
		fmt.Println("--------- END ---------\n")
	}

	fmt.Println("\n#-#-#-#-# Syntax Analisys #-#-#-#-#\n")

	for _, tkTable := range tkTables {
		if err := syntax.ValidateSyntax(tkTable); err != nil {
			fmt.Println(err.String())
			return
		}
	}
	fmt.Println("OK")

	fmt.Println("\n#-#-#-#-# Getting statements #-#-#-#-#\n")

	// testing
	test := tkTables[0]
	st := statement.NewVariable(test[1], test[3])
	fmt.Println(*st.Identifier, *st.Value)
}
