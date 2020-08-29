package math

import "testing"

const defaultError = "Valor esperado %v, mas o resultado encontrado foi %v."

// Os nomes dos testes devem começar com a palavra "Test"
func TestMedia(t *testing.T) {
	expected := 7.28
	media := Media(7.2, 9.9, 6.1, 5.9)

	if media != expected {
		t.Errorf(defaultError, expected, media)
	}
}
