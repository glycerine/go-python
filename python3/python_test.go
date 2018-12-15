package python3

import (
	"github.com/glycerine/go-python/pytest"
	"testing"
)

func TestPython3(t *testing.T) {
	pytest.TestRuntime(t, Runtime)
}
