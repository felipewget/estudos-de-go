package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	// Pode adicionar outros meotodos sem problema algum
}

type bmw7 struct{}

func (b bmw7) fazerBaliza() {
	fmt.Println("aa")
}

func (b bmw7) ligarTurbo() {
	fmt.Println("aa")
}

// Atende as 3 interfaces então adiciona as 3 interfaces

func main() {

	var b esportivoLuxuoso = bmw7{}

	b.ligarTurbo()
	b.fazerBaliza()

}
