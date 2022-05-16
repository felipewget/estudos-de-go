package main

import "fmt"

func notasParaConceito(n float64) string {

	// Multiplas selecoes, o switch é melhor
	if n >= 9 && n < 10 {
		return "A"
	} else if n >= 7 && n < 9 {
		return "B"
	} else if n >= 5 && n < 7 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}

}

func main() {

	fmt.Println(notasParaConceito(9.8))
	fmt.Println(notasParaConceito(6.9))
	fmt.Println(notasParaConceito(2.1))
}
