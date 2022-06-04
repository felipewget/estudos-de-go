package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
	Misturar 2 canais dentro de um unico canal
*/

func encaminhar(origem <-chan string, destino chan string) {

	for {
		destino <- <-origem
	}

}

// Multiplexar - misturar num unico canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {

	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c

}

// <-chan canal somente leitura
// Essa funcao ja retorna o channel antes do for
// Copiando do exercicio anterior
func titulo(urls ...string) <-chan string {

	c := make(chan string)

	for _, url := range urls {

		go func(url string) {

			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<")
			c <- r.FindStringSubmatch(string(html))[1]

		}(url)

	}

	return c

}

func main() {

	c := juntar(
		titulo("https://google.com", "https://youtube.com"),
		titulo("https://youtube.com", "https://youtube.com"),
	)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
