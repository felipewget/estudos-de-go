package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // dessa forma, vai criar campos anonimos do tipo carro, ou seja, ja tem nome e velocidade
	turboLigado bool
}

func main() {

	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Println(f)

	// Print vao mostrar: Estrutura carro e dps a turbo ligado
	//{{F40 0} true}

}
