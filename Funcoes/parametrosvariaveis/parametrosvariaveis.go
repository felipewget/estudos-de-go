package main

import "fmt"

// Funcoes com operador de parametros variados... pode passar 1,2,3,4,5
// chamase funcoes variaticas

// ...TIPO, significa que é uma funcao variatica
func mediana(numeros ...float64) (media float64) {
	total := 0.0

	if len(numeros) == 0 {
		media = 0
		return
	}

	for _, num := range numeros {
		total += num
	}

	media = total / float64(len(numeros))
	return
}

func main() {

	fmt.Printf("Mediana de: %.2f\n", mediana(7.7, 5.5, 0.8, 4, 6.3))
	fmt.Printf("Mediana de: %.2f\n", mediana())

}
