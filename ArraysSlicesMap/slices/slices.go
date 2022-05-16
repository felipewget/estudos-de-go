package main

import (
	"fmt"
	"reflect"
)

func main() {

	a1 := [3]int{1, 2, 3} // Array

	// Slices tem tamanho variavel, ou seja, posso adicionar e remover indices
	s1 := []int{1, 2, 3} // slice

	fmt.Println(s1, a1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	// a ideia do slice é ser um pedaco do arr... um arr grnadao e um slice leve pegando parte do array, slice não é array
	// slice é um trecho de um array
	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] // pegando o arrai no indice 1 e 2

	fmt.Println(s2)

	// Tamanho e um ponteiro que pega elementos do array
	// basciamente vc cria um array e um slice e com o slice de adiciona elementos, tira elemtnos e vai trabalhando com um array sem precsiar ficar criando array de estruturas diferentes

}
