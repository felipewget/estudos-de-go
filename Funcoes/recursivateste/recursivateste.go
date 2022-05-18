package main

import "fmt"

func fatorial(n uint) uint {

	if n == 0 {
		return 1
	}
	fatorialAnterior := fatorial(n - 1)
	return n * fatorialAnterior

}

func main() {
	resultado := fatorial(3)
	fmt.Println(resultado)
}
