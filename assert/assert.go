// Package assert provides test assertions that mark the test as failed but allow execution to
// continue by calling [testing.T.Errorf].
package assert

import (
	"testing"

	"github.com/teleivo/assertive/internal/report"
)

// NoError expects err to be nil. An optional message can be provided as a format string followed by
// args.
func NoError(t *testing.T, err error, msgAndArgs ...any) {
	t.Helper()

	report.NoError(t, t.Errorf, err, msgAndArgs...)
}

// False expects got to be false. An optional message can be provided as a format string followed by
// args.
func False(t *testing.T, got bool, msgAndArgs ...any) {
	t.Helper()

	report.False(t, t.Errorf, got, msgAndArgs...)
}

// True expects got to be true. An optional message can be provided as a format string followed by
// args.
func True(t *testing.T, got bool, msgAndArgs ...any) {
	t.Helper()

	report.True(t, t.Errorf, got, msgAndArgs...)
}

// Nil expects got to be nil. An optional message can be provided as a format string followed by
// args.
func Nil(t *testing.T, got any, msgAndArgs ...any) {
	t.Helper()

	report.Nil(t, t.Errorf, got, msgAndArgs...)
}

// NotNil expects got to not be nil. An optional message can be provided as a format string followed
// by args.
func NotNil(t *testing.T, got any, msgAndArgs ...any) {
	t.Helper()

	report.NotNil(t, t.Errorf, got, msgAndArgs...)
}

// Equals expects got to equal want using the == operator. An optional message can be provided as a
// format string followed by args.
func Equals[T comparable](t *testing.T, got, want T, msgAndArgs ...any) {
	t.Helper()

	report.Equals(t, t.Errorf, got, want, msgAndArgs...)
}

// EqualValues expects got to equal want. Equality is determined using
// [github.com/google/go-cmp/cmp.Diff]. An optional message can be provided as a format string
// followed by args.
func EqualValues[T any](t *testing.T, got, want T, msgAndArgs ...any) {
	t.Helper()

	report.EqualValues(t, t.Errorf, got, want, msgAndArgs...)
}

// NoDiff expects got to equal want. When they differ, the failure message shows a line-level diff
// in gutter format with whitespace made visible on changed lines. Deletions are colored red and
// insertions green unless the NO_COLOR environment variable is set.
// An optional message can be provided as a format string followed by args.
func NoDiff(t *testing.T, got, want string, msgAndArgs ...any) {
	t.Helper()

	report.NoDiff(t, t.Errorf, got, want, msgAndArgs...)
}
