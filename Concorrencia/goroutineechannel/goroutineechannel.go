package main

import (
	"fmt"
	"time"
)

// Channel(canal) é a forma de comunicação entre as goroutines, quando um dado não chega, ele fica esperando
// Channel é um tipo de variavel como qualquer outro... int, string, float64
// Goroutines serve pra processar de forma concorrente

func doisTresQuatroVezes(base int, c chan int) {
	c <- 2 * base // enviando dados pro canal

	time.Sleep(time.Second)
	c <- 3 * base // enviando dados pro canal

	time.Sleep(time.Second * 3)
	c <- 4 * base // enviando dados pro canal
}

func main() {
	c := make(chan int)
	d := make(chan int)

	go doisTresQuatroVezes(2, c)
	go doisTresQuatroVezes(50, d)

	fmt.Println("Antes do channel")

	a, b := <-c, <-c // recebendo os 2 primeiros dados do canal
	g, h := <-d, <-d // recebendo os 2 primeiros dados do canal

	fmt.Println("Só executa depois de receber os 2 primeiros dados do channel")

	fmt.Println(a, b)
	fmt.Println(g, h)

	// Note que segura a thread até retornar os dados necesarios

	fmt.Println(<-c)
	fmt.Println(<-d)

	// fmt.Println(<-c) se sentar receber mais um valor mas nao tem mais valor pra receber, da um erro
	// Deadlock, all goroutines are asleep
}
