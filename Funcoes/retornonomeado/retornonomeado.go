package main

import "fmt"

// Funcao de retorno nomeado
func trocar(p1, p2 int) (segundo int, primeiro int) {

	segundo = p2
	primeiro = p1
	// O x ta no retorno limpo
	return // retorno limpo pq eu nao preciso adicionar o que vai retornar

}

func main() {

	x, y := trocar(5, 6)

	fmt.Println(x, y)

	segundo, primeiro := trocar(5, 6)

	fmt.Println(segundo, primeiro)

}
