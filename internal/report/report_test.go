package report

import "testing"

func TestIsNil(t *testing.T) {
	tests := []struct {
		in   any
		want bool
	}{
		{
			in:   nil,
			want: true,
		},
		{
			in:   1,
			want: false,
		},
	}

	for _, tc := range tests {
		got := isNil(tc.in)

		if got != tc.want {
			t.Errorf("isNil(%v) = got %v, want %v instead", tc.in, got, tc.want)
		}
	}

	t.Run("chan", func(t *testing.T) {
		var in chan struct{}

		got := isNil(in)

		if !got {
			t.Errorf("isNil(%v) = got %v, want %v instead", in, got, true)
		}

		in = make(chan struct{})

		got = isNil(in)

		if got {
			t.Errorf("isNil(%v) = got %v, want %v instead", in, got, true)
		}
	})

	t.Run("map", func(t *testing.T) {
		var in map[string]struct{}

		got := isNil(in)

		if !got {
			t.Errorf("isNil(%v) = got %v, want %v instead", in, got, true)
		}

		in = make(map[string]struct{})

		got = isNil(in)

		if got {
			t.Errorf("isNil(%v) = got %v, want %v instead", in, got, true)
		}
	})

	t.Run("slice", func(t *testing.T) {
		var in []int

		got := isNil(in)

		if !got {
			t.Errorf("isNil(%v) = got %v, want %v instead", in, got, true)
		}

		in = make([]int, 2)

		got = isNil(in)

		if got {
			t.Errorf("isNil(%v) = got %v, want %v instead", in, got, true)
		}
	})
}
