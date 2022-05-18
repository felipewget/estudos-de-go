// palavra reservada defer
// atrasa o momento antes do metodo chamar o return
// abre uma conexa com o banco de dados e quer lembrar de fechar a conexao antes de sair do metodo, na linha seguinte e adicionar defer e a linha que
// faz o que precisa, tp, fechar a conexao e no final, executa essa linha e retorna depois do defer
package main

import "fmt"

func obterValorAprovado(numero int) int {
	var x int = 50

	// essa linha é executada antes do return... retornou 50, nao pega as variaveis
	// então é bom pra executar funcoes sem parametros
	defer fmt.Println("Fim!", x)

	if numero > 5000 {
		x = 46
		fmt.Println("Valor alto")
		return 5000
	} else {
		x = 26
		fmt.Println("Valor baixo")
		return numero
	}
}

func main() {
	obterValorAprovado(7000)
}
