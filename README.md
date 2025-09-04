# goshka (tiny go in Russian)

[![Go Reference](https://pkg.go.dev/badge/github.com/plutov/goshka.svg)](https://pkg.go.dev/github.com/plutov/goshka)

Go utils for minimalists.

Principles:
- No external dependencies: uses only Go standard library
- Small and simple, easy to read and understand
- Fast and efficient

## Testing helpers

```go
import (
    "testing"

    "github.com/plutov/goshka/is"
)

func TestSomething(t *testing.T) {
    is.Equal(t, 1, 1)

    var err error
    is.NilErr(t, nil)
}
```
