package main

import (
	"fmt"
	"time"
)

func notaPraConceito(n float64) string {

	// Converto pra int
	var nota = int(n)

	switch nota {

	case 10:
		// Em Goland nao precisa adicionar o break, em Go ele entra num case e nao executa outros cases
		// e entao vc executa o case abaixo
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "INota nvalida"

	}

}

func modoSwitch2() {
	t := time.Now()
	switch { // quando vc cria um switch sem nenhum valor associado, o switch vai procurar um case que seja true
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}
}

func tipo(i interface{}) string { // Com tipo interface eu posso trabalhar com qualquer tipo de dados... string, inf, float o que for

	switch i.(type) {
	case int:
		return "Integer"
	case float32, float64:
		return "Float"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "I don't know"
	}

}

func main() {

	fmt.Println(notaPraConceito(9.5))
	fmt.Println(notaPraConceito(7))
	fmt.Println(notaPraConceito(5))
	fmt.Println(notaPraConceito(6.2))
	fmt.Println(notaPraConceito(4.6))
	fmt.Println(notaPraConceito(10))
	fmt.Println("------")
	modoSwitch2()
	fmt.Println("------")
	fmt.Println(tipo(5))
	fmt.Println(tipo("!aasdas"))
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now().Hour()))
	fmt.Println(tipo(time.Now()))

}
