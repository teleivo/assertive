package report

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// report notifies a user of a failed assertion. Functions like t.Errorf, t.Fatalf.
type report func(string, ...any)

func NoError(t *testing.T, fn report, err error) {
	t.Helper()

	NoErrorf(t, fn, err, "")
}

func NoErrorf(t *testing.T, fn report, err error, format string, args ...any) {
	t.Helper()

	if err != nil {
		base := fmt.Sprintf("got error %q want none instead", err)
		reportFn(t, fn, base, format, args...)
	}
}

func False(t *testing.T, fn report, got bool) {
	t.Helper()

	Falsef(t, fn, got, "")
}

func Falsef(t *testing.T, fn report, got bool, format string, args ...any) {
	t.Helper()

	if got {
		base := fmt.Sprintf("got %t want %t instead", got, false)
		reportFn(t, fn, base, format, args...)
	}
}

func True(t *testing.T, fn report, got bool) {
	t.Helper()

	Truef(t, fn, got, "")
}

func Truef(t *testing.T, fn report, got bool, format string, args ...any) {
	t.Helper()

	if !got {
		base := fmt.Sprintf("got %t want %t instead", got, true)
		reportFn(t, fn, base, format, args...)
	}
}

func Nil(t *testing.T, fn report, got any) {
	t.Helper()

	Nilf(t, fn, got, "")
}

func Nilf(t *testing.T, fn report, got any, format string, args ...any) {
	t.Helper()

	if !isNil(got) {
		base := fmt.Sprintf("got %v want nil instead", got)
		reportFn(t, fn, base, format, args...)
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

func NotNil(t *testing.T, fn report, got any) {
	t.Helper()

	NotNilf(t, fn, got, "")
}

func NotNilf(t *testing.T, fn report, got any, format string, args ...any) {
	t.Helper()

	if isNil(got) {
		base := "got nil want != nil instead"
		reportFn(t, fn, base, format, args...)
	}
}

func Equals(t *testing.T, fn report, got, want any) {
	t.Helper()

	Equalsf(t, fn, got, want, "")
}

func Equalsf(t *testing.T, fn report, got, want any, format string, args ...any) {
	t.Helper()

	if got != want {
		base := fmt.Sprintf("got %v want %v instead", got, want)
		reportFn(t, fn, base, format, args...)
	}
}

func EqualValues(t *testing.T, fn report, got, want any) {
	t.Helper()

	EqualValuesf(t, fn, got, want, "")
}

func EqualValuesf(t *testing.T, fn report, got, want any, format string, args ...any) {
	t.Helper()

	if diff := cmp.Diff(want, got); diff != "" {
		base := fmt.Sprintf("mismatch (-want +got):\n%s", diff)
		reportFn(t, fn, base, format, args...)
	}
}

func reportFn(t *testing.T, fn report, base, format string, args ...any) {
	t.Helper()

	if format == "" {
		fn(base)
		return
	}

	format = base + "\nmessage=" + format
	fn(format, args...)
}
