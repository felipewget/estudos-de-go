package main

import "fmt"

func main() {

	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	// Slice sem um ponteiro referenciando pra um array
	// tem 10 numeros intanicados mas ele tem capacidade pra 20 elementos
	s = make([]int, 10, 20)
	fmt.Println(s)

	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	fmt.Println(s, len(s), cap(s))

	// E se adicionar mais 1?
	// o Slice aumenta a estrutura, é maleavel, ou seja, vai pra 40 a capacidade e com 21 indexes setados
	s = append(s, 1)

	fmt.Println(s, len(s), cap(s))

}
