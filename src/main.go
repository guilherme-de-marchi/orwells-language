package main

import (
	"fmt"

	"github.com/Guilherme-De-Marchi/orwells-language/src/compiler"
)

func main() {
	text := "if exec var a = \" oi \" 10 10.2 { } + - == "

	tokens := compiler.LexicalAnalisys(text)
	for _, v := range tokens {
		fmt.Println(*v)
	}
}
