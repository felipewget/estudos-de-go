package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {

	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprint(pessoa, " esta falando: ", i)
		}
	}()

	return c

}

func juntar(entrada1, entrada2 <-chan string) <-chan string {

	c := make(chan string)
	go func() {
		for {
			select {
			case e1 := <-entrada1:
				c <- e1
			case e2 := <-entrada2:
				c <- e2
			}
		}
	}()

	return c

}

func main() {

	c := juntar(
		falar("João"),
		falar("Maria"),
	)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c, <-c)

}
