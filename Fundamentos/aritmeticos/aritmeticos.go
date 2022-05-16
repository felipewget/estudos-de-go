package main

import (
	"fmt"
	"math"
)

func main() {

	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Subtracao =>", a-b)
	fmt.Println("Divisao =>", a/b)
	fmt.Println("Multiplicacao =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// Opreações bit a bit
	fmt.Println("E => ", a&b)
	fmt.Println("Ou =>", a|b)
	fmt.Println("Xor =>", a^b)

	c := 3.0
	d := 4.0

	fmt.Println("Maior valor =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor valor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))

}
