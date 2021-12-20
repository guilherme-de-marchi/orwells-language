package main

import (
	"fmt"

	"github.com/Guilherme-De-Marchi/orwells-language/src/compiler"
)

func main() {
	text := "var a = 17.54 ; var b = as ; var comida = frango ;"

	tokens := compiler.LexicalAnalisys(text)
	for _, v := range tokens {
		fmt.Println(*v)
	}

	fmt.Println("\n###########\n")

	instructions := compiler.SplitOn(tokens, compiler.INSTRUCTION_DELIMITER_TOKEN)
	for _, v := range instructions {
		fmt.Println("-----------")
		for _, w := range v {
			fmt.Println(*w)
		}
		fmt.Println("-----------")
	}

	fmt.Println("\n###########\n")

	syntax := compiler.SyntaxAnalisys(instructions)
	fmt.Println(syntax)
}
