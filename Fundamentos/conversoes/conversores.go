package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4 // convert to int int()
	y := 2   // Convert to float float64()

	fmt.Println(x / float64(y))

	// converter o int pra instring
	// string() converte pra letra do teclaro... a
	// fmt.Println("Teste " + string(97))

	fmt.Println("Teste " + strconv.Itoa(97))

	// String to int
	// usei _ pq so retorna mais de um valor... e uso o valor _ e nao crio err pq se nao preciso
	// usar a variavel err pq todas as variaveis tem q ser usadas
	num, _ := strconv.Atoi("1561")
	fmt.Println("agora é um int ", num)

	// Converto string to boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("true")
	}

}
