package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	// Posso usar pq eu estou dentro do pacote main...
	// e privado no nivel do pacote e nao nos arquivos
	fmt.Println(catetos(p1, p2))

	fmt.Println(Distancia(p1, p2))
}
