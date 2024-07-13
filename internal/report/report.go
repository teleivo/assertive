package report

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// report notifies a user of a failed assertion. Functions like t.Errorf, t.Fatalf.
type report func(string, ...any)

func NoError(t *testing.T, fn report, err error) {
	NoErrorf(t, fn, err, "")
}

func NoErrorf(t *testing.T, fn report, err error, format string, args ...any) {
	t.Helper()

	if err != nil {
		form := fmt.Sprintf("got error %q want none instead", err)
		if len(format) > 0 {
			form += "\nmessage=" + format
		}
		fn(form, args...)
	}
}

func False(t *testing.T, fn report, got bool) {
	t.Helper()

	if got {
		fn(fmt.Sprintf("got %t want %t instead", got, false))
	}
}

func True(t *testing.T, fn report, got bool) {
	t.Helper()

	if !got {
		fn(fmt.Sprintf("got %t want %t instead", got, true))
	}
}

func Nil(t *testing.T, fn report, got any) {
	t.Helper()

	if got != nil {
		fn(fmt.Sprintf("got %v want nil instead", got))
	}
}

func NotNil(t *testing.T, fn report, got any) {
	t.Helper()

	if got == nil {
		fn("got nil want != nil instead")
	}
}

func Equals(t *testing.T, fn report, got, want any) {
	t.Helper()

	if got != want {
		fn(fmt.Sprintf("got %v want %v instead", got, want))
	}
}

func EqualValues(t *testing.T, fn report, got, want any) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		fn(fmt.Sprintf("mismatch (-want +got):\n%s", diff))
	}
}
