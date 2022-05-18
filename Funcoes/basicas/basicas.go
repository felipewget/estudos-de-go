package main

import "fmt"

// @NOTA QUEIME NEURONIOS ESCOLHENDO NOMES PRAS FUNCOES

// Essa funcao nao tem parametros e nem retorno
func f1() {
	fmt.Println("F1")
}

// Funcao com parametros, o parametro e o tipo de parametro,
// Essa funcao tbm nao retorna nada
func f2(p1 string, p2 float64) {
	fmt.Printf("F2, paramtros: %s %f\n", p1, p2)
}

// Essa funcao retorna uma string
// Obrigatorio retornar um valor
func f3() string {
	return "F3"
}

// Essa funcao recebe 2 parametros string e uma float e atribui o tipo string as 2
// ela tbm retornara um valor string
func f4(p1, p2 string, p3 float64) string {
	return "F4"
}

// FUNCAO PURA é uma funcao que nao mexe em nada que esta fora da funcao
// nao altera nada externo

// Essa funcao nao recebe parametros mas tem 3 retornos espeficicados
func f5() (string, string, [3]int) {
	return "retorno 1", "retorno 2", [3]int{1, 2, 3}
}

// NEM TODA FUNCAO EU PRECISO ATRIBUIR A UMA VARIAVEL MAS SE EU RECEBER UM,
// PRECISO TRATAR TODOS OS RETORNOS(IGNORANDO OU NAO)

func main() {
	f1()
	f2("Hello world", 12.90)

	r3, r4 := f3(), f4("ola", "ola 2", 1.99)
	fmt.Println(r3)
	fmt.Println(r4)

	r5_1, _, array3 := f5()
	fmt.Println(r5_1, array3)

	// Ignorando todos os retornos
	f5()
}
