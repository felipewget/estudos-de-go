package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operacao bloqueante, só vai prosseguir quando alguem ler esse dado
	fmt.Println("Só depois que for lido que vai chegar aki")
}

func main() {

	// sem buffer... quando tem canal com buffer, pode adicionar de forma assincrona até enxer...
	// quando enxer, precisa ler pra liberar espaço... quando não tem buffer ele ja fica cheio apos o primeiro preenchimento com dados
	channel := make(chan int)

	go rotina(channel)

	// Gorotine + channel é similar a await do NodeJS

	// So vai vir essa linha apo
	fmt.Println(<-channel) // operacao bloqueante

	fmt.Println(<-channel) // enviei apenas um dado, entao vai dar deadlock
	fmt.Println("Fim")

}
