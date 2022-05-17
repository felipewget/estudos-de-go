package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 12351.00,
			"Guga":     99.9,
		},
		"H": {
			"Higor": 13.7,
		},
		"Z": {
			"Zelia": 9999.99,
		},
	}

	fmt.Println(funcPorLetra)

	delete(funcPorLetra, "Z")

	for letra, funcionarios := range funcPorLetra {

		fmt.Println("Letra", letra)

		for nome, salario := range funcionarios {

			fmt.Println("Nome:", nome, "( ", salario, ")")

		}

	}

}
