package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
	Generator ja cuida de abrir um goroutine e esperar o resultado
	nao precisa ficar abrindo chanels nem nada

	generator retorna um canal de apenas leitura

	<-chan -> canal apenas leitura chan é palavra reservada
*/

// <-chan canal somente leitura
// Essa funcao ja retorna o channel antes do for
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

	// Concorrencia, bem mais rapido acessar do que acessar uma por vez
	t1 := titulo("https://google.com", "https://youtube.com")
	t2 := titulo("https://youtube.com", "https://youtube.com")

	fmt.Println("Primeiros: ", <-t1, " | ", <-t2)
	fmt.Println("Segundos: ", <-t1, " | ", <-t2)

}
