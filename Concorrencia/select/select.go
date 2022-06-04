package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

// copiei dos exemplos anteriores
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

func oMaisRapido(url1, url2, url3 string) string {

	c1 := titulo(url1)
	c2 := titulo(url2)
	c3 := titulo(url3)

	// Select retorna a primeira cpmfocap de uma goroutine, a primeira resposta
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(time.Second): // Timeout
		return "Todos perderam"
		// default:
		// 	return "Sem resposta"
	}

}

func main() {

	campeao := oMaisRapido(
		"https://google.com",
		"https://amazon.com",
		"https://youtube.com",
	)

	fmt.Print(campeao)

}
