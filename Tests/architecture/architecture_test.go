package architecture

import (
	"runtime"
	"testing"
)

func TestDependency(t *testing.T) {
	// executa os testes em paralelo
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquiteturas amd64")
	}

	t.Fail()
}
