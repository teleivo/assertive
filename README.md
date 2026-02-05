# assertive

[![Go Reference](https://pkg.go.dev/badge/github.com/teleivo/assertive.svg)](https://pkg.go.dev/github.com/teleivo/assertive)

A minimal assertion library for writing tests in Go.

## Install

```sh
go get github.com/teleivo/assertive
```

## Library

```go
import (
	"github.com/teleivo/assertive/assert"
	"github.com/teleivo/assertive/require"
)

func TestExample(t *testing.T) {
	// soft assertions - test continues on failure
	assert.Equals(t, got, want)
	assert.EqualValuesf(t, got, want, "comparing %q", name)

	// hard assertions - test stops on failure
	require.NoError(t, err)
	require.NotNil(t, result)
}
```

## Disclaimer

I wrote this library for my personal projects and it is provided as-is without warranty. It is
tailored to my needs and my intention is not to adjust it to someone else's liking. Feel free to use
it!

See [LICENSE](LICENSE) for full license terms.
