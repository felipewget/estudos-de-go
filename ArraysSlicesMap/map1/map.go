package main

import "fmt"

func main() {

	// ao inves de ser um indice, passa uma chave, se passa a chave com valores diferentes, o valor é substituido
	// é similar a um dicionario

	// var aprovados = map[int]string
	// Mapas sempre devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[18454901922] = "Maria"
	aprovados[44314345004] = "Joao"

	// fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d) \n", nome, cpf)
	}

	// Posso buscar pela chave
	fmt.Println(aprovados[44314345004])
	delete(aprovados, 44314345004)

	fmt.Println(aprovados[44314345004])

}
