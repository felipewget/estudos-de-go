package matematica

import "testing"

const erro_padrao = "Valor esperado %v, mas o resultado foi %v"

// Esse testing tem tudo que o teste pode precisar, fail, tudo mais
/*
	go test .\matematica.go .\matematica_test.go -run=TestMedia
	go test .\matematica.go .\matematica_test.go


	- Funcao de teste comeca com TestAGUMACOISA
	- Arquivo de execucao do teste tem o _test.go
	- paramtros do teste é t *testing.T
*/

// É necessario comecar com Test
func TestMedia(t *testing.T) {

	valor_esperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valor_esperado {
		t.Errorf(erro_padrao, valor_esperado, valor)
	}

}
