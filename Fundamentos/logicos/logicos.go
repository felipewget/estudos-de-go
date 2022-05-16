package main

import "fmt"

/*
	2 trabalhos, um na terça e quinta
	se conseguir os 2, compra a tv e sorvete
	se um der, pega uma tv menor e sorvete
	se nao der trabalho, nem tv e nem sorvete
*/

// trab1, trab2 bool pega a inferencia pros 2
func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusiveo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {

	tv50, tv32, sorvete := comprar(true, true)

	fmt.Printf("Tv 50: %t, Tv 32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)

}
