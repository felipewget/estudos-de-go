package main

import "fmt"

// Não existe uma marcação com a interface, criou um metodo to string,
// o proprio compilador do GO ja adiciona e respeita a interface
// Bata ter todos os metodos compativel
type imprimivel interface {
	toString() string // Pra saber se é imprimivel ou nao, precisa ter esse metodo
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// Interfaces sao implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2F", p.nome, p.preco)
}

func imprimr(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {

	// Variavel polimorfica
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimr(coisa)

	coisa = produto{"Ze do Picadinho", 1.99}
	fmt.Println(coisa.toString())
	imprimr(coisa)
}

// Polimorfismo... a variavel ter multiplas formas
