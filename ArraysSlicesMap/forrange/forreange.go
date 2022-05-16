package main

import "fmt"

func main() {

	// Os 3 pontinhos... ele é um array, tem um tamanho fixo mas eu quero que vc conta ele quando eu colocar os elementos dentro dele
	// na inicializacao
	numeros := [...]int{1, 2, 3, 4, 5} // Compilador que conta

	for i, numero := range numeros {
		fmt.Println(i, ") é o valor ", numero)
	}

	for _, num := range numeros {
		fmt.Println("ignoro o indice com _ e o valor é ", num)
	}

}
