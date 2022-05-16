package main

import "fmt"

func imprimirResultado(nota float64) {

	// Sem parenteses
	// sempre tem q ter o bloco msm que seja uma linha pq não tem ternário
	if nota >= 7 && (1 == 1) {
		fmt.Println("Aprovado com a nota", nota)
	} else {
		fmt.Println("Reprovado com a nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
