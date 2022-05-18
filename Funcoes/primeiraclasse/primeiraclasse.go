package main

import "fmt"

// forma diferente de chamar uma funcao... funcao anonima
var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
