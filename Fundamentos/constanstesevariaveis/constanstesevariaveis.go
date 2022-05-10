package main

import (
	"fmt"
	m "math" // Posso colocar uma alias... e ai nao preciso usar math.Pow(), so m.pow()
)

// Compilador consegue inferir a tipagem da variavel
func main() {
	const PI float64 = 3.1415 // Constante, nome, tipo e valor
	var raio = 3.2

	// area = significa que a variavel ja existe e esta setando valor
	// area := significa que a variavel nao foi usada e esta tenddo um valor setado
	area := PI * m.Pow(raio, 2) // Variaveis declaradas tem q ser usadas

	fmt.Println("A area da circunferencia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Print("Ola ", a, b, c, d)

	// posso declarar variaveis assim
	var e, f bool = true, false

	g, h, i := 2, false, "epa!"

	// Variaveis sao fortemente tipadas, entao essa variavel nao muda
	fmt.Println(e, f, g, h, i)
}
