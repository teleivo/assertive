package report

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// report notifies a user of a failed assertion. Functions like t.Errorf, t.Fatalf.
type report func(string, ...any)

// NoError fails if err is not nil.
func NoError(t *testing.T, fn report, err error, msgAndArgs ...any) {
	t.Helper()

	if err != nil {
		base := fmt.Sprintf("got error %q want none instead", err)
		reportMsg(t, fn, base, msgAndArgs...)
	}
}

// False fails if got is not false.
func False(t *testing.T, fn report, got bool, msgAndArgs ...any) {
	t.Helper()

	if got {
		base := fmt.Sprintf("got %t want %t instead", got, false)
		reportMsg(t, fn, base, msgAndArgs...)
	}
}

// True fails if got is not true.
func True(t *testing.T, fn report, got bool, msgAndArgs ...any) {
	t.Helper()

	if !got {
		base := fmt.Sprintf("got %t want %t instead", got, true)
		reportMsg(t, fn, base, msgAndArgs...)
	}
}

// Nil fails if got is not nil.
func Nil(t *testing.T, fn report, got any, msgAndArgs ...any) {
	t.Helper()

	if !isNil(got) {
		base := fmt.Sprintf("got %v want nil instead", got)
		reportMsg(t, fn, base, msgAndArgs...)
	}
}

func isNil(got any) bool {
	if got == nil {
		return true
	}

	v := reflect.ValueOf(got)
	switch v.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Interface,
		reflect.Slice:
		return v.IsNil()
	}

	return false
}

// NotNil fails if got is nil.
func NotNil(t *testing.T, fn report, got any, msgAndArgs ...any) {
	t.Helper()

	if isNil(got) {
		base := "got nil want != nil instead"
		reportMsg(t, fn, base, msgAndArgs...)
	}
}

// Equals fails if got != want.
func Equals[T comparable](t *testing.T, fn report, got, want T, msgAndArgs ...any) {
	t.Helper()

	if got != want {
		base := fmt.Sprintf("got %v want %v instead", got, want)
		reportMsg(t, fn, base, msgAndArgs...)
	}
}

// EqualValues fails if got and want differ according to [cmp.Diff].
func EqualValues[T any](t *testing.T, fn report, got, want T, msgAndArgs ...any) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		base := fmt.Sprintf("mismatch (-want +got):\n%s", diff)
		reportMsg(t, fn, base, msgAndArgs...)
	}
}

// reportMsg reports a failed assertion with an optional user message. If msgAndArgs is non-empty,
// the first element is used as a format string and the rest as format arguments.
func reportMsg(t *testing.T, fn report, base string, msgAndArgs ...any) {
	t.Helper()

	if len(msgAndArgs) == 0 {
		fn(base)
		return
	}

	format, ok := msgAndArgs[0].(string)
	if !ok {
		fn(base)
		return
	}

	msg := fmt.Sprintf(format, msgAndArgs[1:]...)
	fn(base + "\nmessage=" + msg)
}
