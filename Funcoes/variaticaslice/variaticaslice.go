package main

import "fmt"

func media(numeros ...float64) float64 {

	total := 0.0

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))

}

// parametros variados
// é uma funcao variativa... ...tipo
func imprimirAprovados(aprovados ...string) {

	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d), %s\n", i+1, aprovado)
	}

}

// ARRAY NAO FUNCIONA, apenas slices

func main() {

	slice := make([]float64, 5, 5)
	slice[2] = 6.8
	slice[3] = 7

	media_com_slice := media(slice...)
	fmt.Printf("Média %.2f\n", media_com_slice)

	slice2 := []string{"Maria", "Pedro", "Joao"}

	// Adiciono o ... no final eu consigo enviar
	imprimirAprovados(slice2...)

}
