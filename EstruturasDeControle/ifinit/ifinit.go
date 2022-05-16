package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Em Go, você pode criar um bloco de execução antes de executar uma estrutura de if
// e essa variavel do bloco pre if vai ficar disponivel

func numeroAleatorio() int {

	s := rand.NewSource(time.Now().UnixNano())
	// pega o numero da execucao atual em nanosegundo
	r := rand.New(s)
	// Cria um numero aleatorio ate 10
	return r.Intn(10)
}

func main() {
	// criando uma variavel que estará presente apenas no if/else
	// Variavel e inicializacao ; condicoes do if
	if i := numeroAleatorio(); i > 5 { // esse tipo de bloco apenas no if tbm funciona no switch
		fmt.Println("Ganhou")
	} else {
		fmt.Println("Perdeu")
	}

	// Fora do if else essa variavel nao existe mais
	// fmt.Println(i) // se descomentar essa linha ira dar erro pq o i nao existe, é a msm ideia do for

}
