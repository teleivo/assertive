package assert

import (
	"errors"
	"testing"
)

func TestAssert(t *testing.T) {
	t.Skip()

	NoError(t, errors.New("error"))
	NoErrorf(t, errors.New("error"), "failed to %q", "wow")

	Nil(t, 0)
	NotNil(t, []int{})
	var s []int
	NotNil(t, s)

	Truef(t, false, "methodCall(%d)", 10)
}
