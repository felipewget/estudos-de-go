/*
	GO tem ferramenta pra mostrar quanto % do seu codigo foi testado e permite acompanhamento


	go test ./... coverprofile=arquivo
	// posso ler em:
	go tool cover -func=resultado
	// Mosta em html
	 go tool cover -html=resultado
*/

package strings

import (
	"strings"
	"testing"
)

const msg_index = "%s (parte: %s) - Indices: esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {

	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r é legal", "Cod3r", 0},
		{"", "", 0},
		{"felipe", "e", 1},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msg_index, teste.texto, teste.parte, teste.esperado, atual)
		}

	}

}
