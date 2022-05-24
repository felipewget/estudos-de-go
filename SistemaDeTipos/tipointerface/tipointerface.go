package main

import "fmt"

type curso struct {
	nome string
}

func main() {

	var coisa interface{}
	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	// Como é uma interface vazia, ele pode receber qualquer valor
	// algo mais generico

	coisa2 = curso{"curso de GO"}
	fmt.Println(coisa2)

	// É possivel usar interface como um tipo generico

}
