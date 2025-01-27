// Package require contains functions marking tests as failed and stops their execution when
// expected conditions are not met by calling [testing.T.Fatalf]. The mismatch in expectations is
// logged via [testing.T.Logf].
package require

import (
	"testing"

	"github.com/teleivo/assertive/internal/report"
)

// NoError expects err to be nil. The executing test will be marked as failed if the expectation is
// not met.
func NoError(t *testing.T, err error) {
	t.Helper()

	report.NoError(t, t.Fatalf, err)
}

// NoErrorf expects err to be nil. The executing test will be marked as failed if the expectation is
// not met. A message will be logged via [testing.T.Logf] with given format and args.
func NoErrorf(t *testing.T, err error, format string, args ...any) {
	t.Helper()

	report.NoErrorf(t, t.Fatalf, err, format, args...)
}

// False expects got to be false. The executing test will be marked as failed if the expectation is
// not met.
func False(t *testing.T, got bool) {
	t.Helper()

	report.False(t, t.Fatalf, got)
}

// Falsef expects got to be false. The executing test will be marked as failed if the expectation is
// not met. A message will be logged via [testing.T.Logf] with given format and args.
func Falsef(t *testing.T, got bool, format string, args ...any) {
	t.Helper()

	report.Falsef(t, t.Fatalf, got, format, args...)
}

// True expects got to be true. The executing test will be marked as failed if the expectation is
// not met.
func True(t *testing.T, got bool) {
	t.Helper()

	report.True(t, t.Fatalf, got)
}

// Truef expects got to be true. The executing test will be marked as failed if the expectation is
// not met. A message will be logged via [testing.T.Logf] with given format and args.
func Truef(t *testing.T, got bool, format string, args ...any) {
	t.Helper()

	report.Truef(t, t.Fatalf, got, format, args...)
}

// Nil expects got to be nil. The executing test will be marked as failed if the expectation is not
// met.
func Nil(t *testing.T, got any) {
	t.Helper()

	report.Nil(t, t.Fatalf, got)
}

// Nilf expects got to be nil. The executing test will be marked as failed if the expectation is not
// met. A message will be logged via [testing.T.Logf] with given format and args.
func Nilf(t *testing.T, got any, format string, args ...any) {
	t.Helper()

	report.Nilf(t, t.Fatalf, got, format, args...)
}

// NotNil expects got to be nil. The executing test will be marked as failed if the expectation is
// not met.
func NotNil(t *testing.T, got any) {
	t.Helper()

	report.NotNil(t, t.Fatalf, got)
}

// NotNilf expects got to be nil. The executing test will be marked as failed if the expectation is
// not met. A message will be logged via [testing.T.Logf] with given format and args.
func NotNilf(t *testing.T, got any, format string, args ...any) {
	t.Helper()

	report.NotNilf(t, t.Fatalf, got, format, args...)
}

// Equals expects got to equal want. The executing test will be marked as failed if the expectation
// is not met.
func Equals(t *testing.T, got, want any) {
	t.Helper()

	report.Equals(t, t.Fatalf, got, want)
}

// Equalsf expects got to equal want. The executing test will be marked as failed if the expectation
// is not met. A message will be logged via [testing.T.Logf] with given format and args.
func Equalsf(t *testing.T, got, want any, format string, args ...any) {
	t.Helper()

	report.Equalsf(t, t.Fatalf, got, want, format, args...)
}

// EqualValues expects got to be equal want. Equality is determined by calling
// [github.com/google/go-cmp/cmp.Diff]. The executing test will be marked as failed if the
// expectation is not met.
func EqualValues(t *testing.T, got, want any) {
	t.Helper()

	report.EqualValues(t, t.Fatalf, got, want)
}

// EqualValues expects got to be equal want. Equality is determined by calling
// [github.com/google/go-cmp/cmp.Diff]. The executing test will be marked as failed if the
// expectation is not met. A message will be logged via [testing.T.Logf] with given format and args.
func EqualValuesf(t *testing.T, got, want any, format string, args ...any) {
	t.Helper()

	report.EqualValuesf(t, t.Fatalf, got, want, format, args...)
}
