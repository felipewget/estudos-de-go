package main

import "fmt"

type item struct {
	produtoId  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item // Esse diz que é varios itens do tipo item no pedido a cima
}

func (p pedido) valorTotal() (total float64) {

	total = 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return

}

func main() {

	// Lição do dia... escrever mais pra trazer mais clareza pro programador
	// vale mais do que código curto

	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoId: 1, quantidade: 2, preco: 12.10},
			item{2, 2, 12.10},
			item{3, 2, 12.10},
		},
	}

	fmt.Printf("Total é %.2f", pedido.valorTotal())

}
