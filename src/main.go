package main

import (
	"fmt"

	"github.com/Guilherme-De-Marchi/orwells-language/src/compiler"
)

func main() {
	text := "var a = $b ;"

	fmt.Println("\n#-#-#-#-# Lexical Analisys #-#-#-#-#\n")

	tokens := compiler.LexicalAnalisys(text)
	for _, v := range tokens {
		fmt.Println(*v)
	}

	fmt.Println("\n#-#-#-#-# Getting the instructions #-#-#-#-#\n")

	instructions := compiler.SplitOn(tokens, compiler.INSTRUCTION_DELIMITER_TOKEN)
	for _, v := range instructions {
		fmt.Println("----- Block -----")
		for _, w := range v {
			fmt.Println(*w)
		}
		fmt.Println("----- END Block -----\n")
	}

	fmt.Println("\n#-#-#-#-# Syntax Analisys #-#-#-#-#\n")

	err := compiler.SyntaxAnalisys(instructions)
	if err != nil {
		fmt.Println(*err)
		return
	} else {
		fmt.Println("OK")
	}
}
