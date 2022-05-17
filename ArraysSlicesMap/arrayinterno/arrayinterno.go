package main

import "fmt"

// O mesmo array interno é referenciado por varios slices
func main() {

	s := make([]int, 5, 10)
	s2 := append(s, 1, 2, 3)

	fmt.Println(s, s2)

	s[0] = 7

	// Mudou pros 2 slices ou seja, é o msm endereco de index
	fmt.Println(s, s2)

}
