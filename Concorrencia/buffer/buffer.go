package main

import (
	"fmt"
	"time"
)

/*
	Buffer É um espaço pra adicionar dados, ele para quando o buffer enxer
	-> se tenho um buffer de 10 posições ele nao vai travar o channel, só vai travar quando enxer o buffer, no caso o elemento11
*/

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou") // Note que esse log não envia pq tem 3 espaços porém se eu ler o canal, libera um espaço
	ch <- 5
	ch <- 6
}

func main() {

	ch := make(chan int, 3) // tem 3 slots
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)

	fmt.Println("Final")
}
