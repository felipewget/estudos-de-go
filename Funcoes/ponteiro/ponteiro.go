package main

// Vamos passar um ponteiro pra uma funcao(que vai acessar a memoria e pegar o valor da variavel que nao tem a ver com seu escopo)

func incrementa1(n int) {
	n++
}

// Revisão: o ponteiro é representado por um * e é um cursor/ponteiro pra um endereço de memoria
func incremento2(n *int) {
	// revisão: * é usado pra desrefenciar, ou seja,
	// ter acesso ao valor onde o ponteiro aponta
	*n++
}

func main() {
	n := 1

	incrementa1(n)
	println("Sem ponteiro", n)
	incrementa1(n)
	println("Sem ponteiro", n)
	incrementa1(n)
	println("Sem ponteiro", n)

	// &n eu passo o cursor, o endereco de memoria da variavel
	// Essas funcoes nao sao mais puras
	incremento2(&n)
	println("Com ponteiro", n)
	incremento2(&n)
	println("Com ponteiro", n)
	incremento2(&n)
	println("Com ponteiro", n)
	incremento2(&n)
	println("Com ponteiro", n)

}
