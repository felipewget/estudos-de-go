package main

import "fmt"

func main() {

	/*
		O canal é algo intrícico na linguágem e channel é um tipo da linguagem(tipo int, tipo float64)
	*/

	// Segundo parametro é um buffer... depois via ter explicação
	ch := make(chan int, 1)
	ch <- 1
	<-ch

	// Como enviar e como reber a informação de um canal pra itens que estão sendo tratados de formas independentes

	ch <- 2 // estou enviando para o canal o valor 2
	fmt.Println(<-ch)

	// Canal é um tipo de sincronismo, a partir do canal eu posso esperar os dados chegarem de ma goroutine

}
