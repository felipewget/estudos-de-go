package main

import "fmt"

func main() {

	funcsESalarios := map[string]float64{
		"José": 125.00,
		"João": 99.9,
	}

	funcsESalarios["Ola"] = 1280.00

	fmt.Println(funcsESalarios)

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
