package main

import (
	"fmt"
	"math"
	"reflect"
)

// Reflet fala o tipo de uma variavel

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Tipo de 3200 é ", reflect.TypeOf(3200)) // Tipo da variavel

	// sem sinal(só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255 // byte é um apelido pra uint8
	fmt.Println("O byte é ", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor maximo é", i1)
	fmt.Println("o tipo é", reflect.TypeOf(i1))

	var i2 rune = 'a' // Unicode pra int32
	fmt.Println(i2)

	var x float32 = 49.99
	// se colocar 49.99, ele define float64
	fmt.Println("Tipo de x é", reflect.TypeOf(x))

	// Boolean
	bo := true
	// !1 .... uns nao funciona... só true or false
	fmt.Println("Tipo de x é", reflect.TypeOf(bo))

	// Nao é permitido criar string com aspas simples, apenas aspas duplas
	s1 := "Ola, oi"
	fmt.Println("Numero de caracteres de (", s1, ") é ", len(s1))

	s2 := `
		agroa eu posso
		quebrar a linha
		mas na msm escript	
	`

	fmt.Println("(", s2, ") é ", len(s1))

}
