package main

import "fmt"

// Estourar uma pilha, um stackoverflow
// Importante adicionar uma funcao muito clara de entrada

// Muito comum achar isso no GO, uma resposta e um parametros erro
// Tipo error
func fatorial(n int) (int, error) {

	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}

}

func main() {
	resultado, err := fatorial(-3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}
}
