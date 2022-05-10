package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha.")

	fmt.Println(" Nova")

	x := 3.141516

	// Convert float pra string
	xs := fmt.Sprint(x)
	fmt.Println("O valor", xs)

	a := 1
	b := 2

	fmt.Printf("oi %v <> %v", a, b)
}
