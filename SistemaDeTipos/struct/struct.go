package main

import "fmt"

// Struct -> estrutura
// Sem separar por virgula
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Reciver -> quem será o dono dessa funcao, pra poder usar algo em cima... tipo produto.faz alguma coisa

// Método com reciver
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	// Declarar a variavel
	var produto1 produto
	produto1 = produto{
		nome:     "Caderno",
		preco:    23.90,
		desconto: 0.05,
	}

	fmt.Println(produto1.preco, produto1.precoComDesconto())

	produto2 := produto{"Borracha", 2.99, 0.01}

	fmt.Println(produto2.preco, produto2.precoComDesconto())

}
