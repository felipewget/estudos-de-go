package main

import "fmt"

func main() {

	// Estrutura homogênea(mesmo tipo) e estatica(fixo... 10 posições, sempre via ter 10 posicoes)
	// vc pode pegar outra variavel e adicionar mas nao pode alterar o msm array

	// Se nao tiver valor, eles ja vem com o valor default
	var notas [3]float64

	fmt.Println(notas)
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	// Médias
	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("Média %.2f", media)

}
