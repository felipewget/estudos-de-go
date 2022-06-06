package matematica

import (
	"fmt"
	"strconv"
)

/*
	Dentro da linguagem GoLang, existe uma estrutura pensada pra testes,
	a propria linguagem oferece suporte a testes de cobertura e testes unitarios

	Conteito de teste unitario -> étestar uma unidade(metodo, funcao) ... quanto mais isolar uma
	funcao, vc testa a funcao de forma unica

	e quando vc vai criar o sistema, vc ja tem a funcao, os testes, a forma de entrada
	saidas e tudo mais, isso é um tipo de desenvolvimento orientado a testes

	desenvolver com mocs, você pode fazer um moc e simular resposta do banco de dados... ou seja, quando chamar
	a query x, retorna um json pronto

	teste de integracao é uma funcao que sai acessando outras funções, seria um teste de integração com
	outros metodos
*/

// Media é responsavel por calcular a media de numeros
func Media(numeros ...float64) (media float64) {

	total := 0.0
	for _, num := range numeros {
		total += num
	}

	media = total / float64(len(numeros))

	// Apenas 2 casas decimais
	media_com_2_casas, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64) // 64 pq é float64
	media = media_com_2_casas
	return

}
