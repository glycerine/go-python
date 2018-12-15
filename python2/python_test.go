package python2

import (
	"github.com/glycerine/go-python/pytest"
	"testing"
)

func TestPython2(t *testing.T) {
	pytest.TestRuntime(t, Runtime)
}
