package main

import "fmt"

func main() {

	// Zeros é o valor padrao de um tipo de variavel...
	// Quando cravo um valor, nao posso alterar, fortemente tipada
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %v %v ", a, b, c, d, e)
}
