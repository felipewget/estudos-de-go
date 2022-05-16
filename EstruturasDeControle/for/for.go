package main

import (
	"fmt"
	"time"
)

// a unica estrutura de controle de loop é o for, ainda assim existe o map, o slice e td mais
// as variaveis morrem dentro do escopo

func main() {

	// Como se faz o while em GO pois so existe o for
	i := 1
	for i <= 10 { // nao existe while em GO
		fmt.Println(i)
		i++
	}

	// Note que o I dentro desse IF é uma outra variavel
	for i := 0; i < 20; i += 2 {
		fmt.Println(i) // Esse é o I dentro do looping
	}

	fmt.Println(i) // Esse é o I do comeco e nao do i

	fmt.Println("\nMisturando")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	// Loop infinito
	for {

		fmt.Println("Pra sempre...")
		time.Sleep(time.Second) // Sleep

	}

}

// Código raso
