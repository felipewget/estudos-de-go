package main

import (
	"fmt"
	"time"
)

// Comparando apenas valores, nao ta comparando referencia de memorias
// @internal: vai explicar os ponteiros no go mais pra frente

func main() {

	fmt.Println("String:", "Banana" == "Banana")
	fmt.Println("String:", 3 != 2)
	fmt.Println("< :", 3 < 2)
	fmt.Println("> :", 3 > 2)
	fmt.Println("<= :", 3 < 2)
	fmt.Println(">= :", 3 > 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas: ", d1 == d2)
	fmt.Println("Datas: ", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João das Bota"}
	p2 := Pessoa{"Zé Pequeno"}
	fmt.Println("Pessoas: ", p1 == p2)

}
