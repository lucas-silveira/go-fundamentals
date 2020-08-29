package strings_test

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - Índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Lucas", "a", 3},
	}

	for _, test := range tests {
		t.Logf("Massa: %v", test)
		current := strings.Index(test.text, test.part)

		if current != test.expected {
			t.Errorf(msgIndex, test.text, test.part, test.expected, current)
		}
	}
}
