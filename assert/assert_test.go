package assert

import (
	"errors"
	"testing"
)

func TestAssert(t *testing.T) {
	NoError(t, errors.New("error"))
	NoErrorf(t, errors.New("error"), "failed to %q", "wow")
}
