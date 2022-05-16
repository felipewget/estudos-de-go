package main

import "fmt"

func notaPraConceito(nota float64) string {

	switch {
	case nota >= 9 && nota < 10:
		return "A"
	case nota >= 7 && nota < 8:
		return "B"
	case nota >= 5 && nota < 7:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	default:
		return "E"
	}

}

func main() {

	fmt.Println(notaPraConceito(3.2))
	fmt.Println(notaPraConceito(7.5))
	fmt.Println(notaPraConceito(6))
	fmt.Println(notaPraConceito(9))
	fmt.Println(notaPraConceito(9.9))
	fmt.Println(notaPraConceito(7.8))

}
