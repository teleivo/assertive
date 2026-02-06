// Package require provides test assertions that mark the test as failed and stop execution by
// calling [testing.T.Fatalf].
package require

import (
	"testing"

	"github.com/teleivo/assertive/internal/report"
)

// NoError expects err to be nil. An optional message can be provided as a format string followed by
// args.
func NoError(t *testing.T, err error, msgAndArgs ...any) {
	t.Helper()

	report.NoError(t, t.Fatalf, err, msgAndArgs...)
}

// False expects got to be false. An optional message can be provided as a format string followed by
// args.
func False(t *testing.T, got bool, msgAndArgs ...any) {
	t.Helper()

	report.False(t, t.Fatalf, got, msgAndArgs...)
}

// True expects got to be true. An optional message can be provided as a format string followed by
// args.
func True(t *testing.T, got bool, msgAndArgs ...any) {
	t.Helper()

	report.True(t, t.Fatalf, got, msgAndArgs...)
}

// Nil expects got to be nil. An optional message can be provided as a format string followed by
// args.
func Nil(t *testing.T, got any, msgAndArgs ...any) {
	t.Helper()

	report.Nil(t, t.Fatalf, got, msgAndArgs...)
}

// NotNil expects got to not be nil. An optional message can be provided as a format string followed
// by args.
func NotNil(t *testing.T, got any, msgAndArgs ...any) {
	t.Helper()

	report.NotNil(t, t.Fatalf, got, msgAndArgs...)
}

// Equals expects got to equal want using the == operator. An optional message can be provided as a
// format string followed by args.
func Equals[T comparable](t *testing.T, got, want T, msgAndArgs ...any) {
	t.Helper()

	report.Equals(t, t.Fatalf, got, want, msgAndArgs...)
}

// EqualValues expects got to equal want. Equality is determined using
// [github.com/google/go-cmp/cmp.Diff]. An optional message can be provided as a format string
// followed by args.
func EqualValues[T any](t *testing.T, got, want T, msgAndArgs ...any) {
	t.Helper()

	report.EqualValues(t, t.Fatalf, got, want, msgAndArgs...)
}

// NoDiff expects got to equal want. When they differ, the failure message shows a line-level diff
// in gutter format with whitespace made visible on changed lines. This is useful for comparing
// multi-line strings like formatted code output. An optional message can be provided as a format
// string followed by args.
func NoDiff(t *testing.T, got, want string, msgAndArgs ...any) {
	t.Helper()

	report.NoDiff(t, t.Fatalf, got, want, msgAndArgs...)
}
