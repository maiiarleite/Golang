package strings
//Dataset

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()
	testes := []struct {
		texto string
		parte string
		esperado int
	}{
		{"Sorvete é bom", "Sorvete", 0},
		{"","", 0},
		{"Ola", "Oi", -1},
		{"Leonardo","o", 2},
	}

	for _, teste := range testes {
		t.Log("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}