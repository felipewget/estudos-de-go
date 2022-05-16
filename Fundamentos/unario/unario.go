package main

import "fmt"

func main() {

	x := 1
	y := 2

	// Apenas postfix
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // y -= 1
	fmt.Println(y)

	// Em go eu não posso usar um unario dentro de uma comparacao, isso facilita a leitura e go
	// é um metodo minimalista e limpo
	// fmt.Println(x == --y)

}
