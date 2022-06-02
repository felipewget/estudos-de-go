package main

import (
	"fmt"
	"strconv"
	"time"
)

func fale(pessoa, texto string, quantidade int) {

	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Println(pessoa + ": " + texto + " | Interacao: " + strconv.Itoa(i))
	}

}

func main() {

	// // Primeiro termina essa funcao
	// fale("Maria", "pq você não fala?", 3)

	// // Apos vai executar essa funcao
	// fale("João", "Só posso falar depois de você?", 3)

	/*
		Resp:
		Não tem processamentos infependentes

		Maria: pq você não fala? | Interacao: 0
		Maria: pq você não fala? | Interacao: 1
		Maria: pq você não fala? | Interacao: 2
		João: Só posso falar depois de você? | Interacao: 0
		João: Só posso falar depois de você? | Interacao: 1
		João: Só posso falar depois de você? | Interacao: 2
	*/

	// // Coloquei o nome go, ou seja, criou um goroutine, criou de forma independente
	// // como se fosse uma thread independente que consome poucos recursos
	// // Thread -> linha de execução independente

	// // Não teve o print, pq só via continuar se a tread principal estiver funcionando
	// go fale("Maria", "pq você não fala?", 5)
	// go fale("João", "Só posso falar depois de você?", 5)

	// // Há uma thread principal após o paralelismo, isso deve fazer printar
	// // Se a goroutine tiver fazer algo e a thred principal tiver terminado antes da goroutine, não vai printar

	// time.Sleep(time.Second * 3) // após terminar os 3 segundos, ele fecha a thread principal e as goroutines

	// fmt.Println("Fim")

	// ------

	// Chama as 2 funções de forma assincrona, e chama as 2 ao mesmo tempo, no que o joao terminar de falar,
	// fecha a thread principal e o programa se encerra
	go fale("Maria", "pq você não fala?", 5)
	fale("João", "Só posso falar depois de você?", 5)

	/*
		João: Só posso falar depois de você? | Interacao: 0
		Maria: pq você não fala? | Interacao: 0
		Maria: pq você não fala? | Interacao: 1
		João: Só posso falar depois de você? | Interacao: 1
		João: Só posso falar depois de você? | Interacao: 2
		Maria: pq você não fala? | Interacao: 2
		Maria: pq você não fala? | Interacao: 3
		João: Só posso falar depois de você? | Interacao: 3
		João: Só posso falar depois de você? | Interacao: 4
	*/

	// --------------

	// Como esperar a concorrencia acabar?

	// GOROUTINES trabalham em concorrencia

}
