package main

import "fmt"

// Essa funcao é uma funcao de inicializacao, eu nem preciso chamar esse metodo pq ele vai ser executado antes
// do metodo main... é similar ao __contruct() do PHP ou contructor() do JS
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main")
}
