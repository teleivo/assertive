package assert

import (
	"errors"
	"testing"
)

func TestAssert(t *testing.T) {
	t.Skip()

	NoError(t, errors.New("error"))
	NoError(t, errors.New("error"), "failed to %q", "wow")

	Nil(t, 0)
	NotNil(t, []int{})
	var s []int
	NotNil(t, s)

	True(t, false, "methodCall(%d)", 10)
	True(t, false, "methodCall()")

	NoDiff(t,
		"func main() {\n\tfmt.Println(\"hello\")\n}\n",
		"func main() {\n\tfmt.Println(\"world\")\n}\n",
	)
}
