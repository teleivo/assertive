// Package require contains functions marking tests as failed and stops their execution when
// expected conditions are not met by calling [testing.T.Fatalf]. The mismatch in expectations is
// logged via [testing.T].
package require

import (
	"testing"

	"github.com/teleivo/assertive/internal/report"
)

func NoError(t *testing.T, err error) {
	t.Helper()

	report.NoError(t, t.Fatalf, err)
}

func False(t *testing.T, got bool) {
	t.Helper()

	report.False(t, t.Fatalf, got)
}

func True(t *testing.T, got bool) {
	t.Helper()

	report.True(t, t.Fatalf, got)
}

func Nil(t *testing.T, got any) {
	t.Helper()

	report.Nil(t, t.Fatalf, got)
}

func NotNil(t *testing.T, got any) {
	t.Helper()

	report.NotNil(t, t.Fatalf, got)
}

func Equals(t *testing.T, got, want any) {
	t.Helper()

	report.Equals(t, t.Fatalf, got, want)
}

func EqualValues(t *testing.T, in, got, want any) {
	t.Helper()

	report.EqualValues(t, t.Fatalf, got, want)
}
