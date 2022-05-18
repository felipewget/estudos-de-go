package main

import "fmt"

// quando eu crio uma funcao dentro de outra funcao, ela agrupa o que ta dentro, ou seja, ele
// pega os contextos pai, do escopo
// funcao com funcao interna, tem memoria do contexto que foi declarada

func closure() func() {
	// Vai ler esse x?
	// Acho que vai ler aki pq essa funcao foi declarada dentro da funcao closure
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	// Vai ler esse x? do escopo pai de todos
	x := 20
	fmt.Println(x)
	imprimeX := closure()
	imprimeX()
}
