package main

import "fmt"

// Cursor|Cursores|ponteiros
func main() {

	// Inteira ocupa 64 bits ou 8 bytes
	// A partir dessa variavel é possivel pegar o endereço na memoria dela e compartilhar esse endereço
	// pra por exemplo colocar ele num ponteiro

	// Go não tem aritmetica de ponteiros
	// Pra criar um ponteiro, só colocar um * e vai ver que é um endereco de memoria, vc pode acessar o valor pelo endereco que a variavel tem
	// é tp um global.variavel no js, toda as outras referencias conseguem ver a alteracao
	i := 1

	// Criando um ponteiro
	// O nil é permitivo pra int apenas pq é um ponteiro
	var p *int = nil // nil == null

	// & <- endereco e coloca o endereco pra p
	// agora o P pega o msm endereco da variavel i
	p = &i // pegando o endereco da variavel

	// agora posso trabalhar com o ponteiro e trabalhar com I e alterar o ponteiro atraves do endreco

	// *p == valor do ponteiro
	// p  == endereco do ponteiro
	*p++
	i++

	// Endereco p e o valor do ponteiro que é o msm de i
	// 0xc0000ac058 3 3
	// &i <- endereco da variavel
	fmt.Println(p, *p, i, &i)

	// &variavel pega o endereco, o ponteiro

}
