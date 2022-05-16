package main

import "fmt"

/*
GO nao tem ternario,
GO é uma linguagem minimalista e que preza pela facilidade da leitura, então não tem ternario
*/

func obterResultado(nota float64) string {

	// return nota >= 6 ? "Aprovado" : "Reprovado" (Não funciona em GO)
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"

}

func main() {
	fmt.Println(obterResultado(6.2))
}
