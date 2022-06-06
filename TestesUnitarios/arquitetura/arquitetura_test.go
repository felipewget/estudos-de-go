package arquitetura

import (
	"runtime"
	"testing"
)

/*
pra mostrar a msg tem q ter o go test -v
t.Skip("Não funciona em arquitetura amd64")

go test ./... <- não funciona no windows
*/

func TestDependente(t *testing.T) {
	t.Parallel() // Esse teste pode ser executado de forma paralela
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}

	// ... continua

	t.Fail()

}
