package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

// Passa um ponteiro pra alterar o valor... se nao passa apenas uma copia do valor
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {

	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// Quando crio o struct a partir da interface, preciso passar o ponteiro
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)

}
