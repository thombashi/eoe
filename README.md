# eoe

Provide a simple function to exit the program with an error message on errors in Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/thombashi/eoe.svg)](https://pkg.go.dev/github.com/thombashi/eoe)
[![Go Report Card](https://goreportcard.com/badge/github.com/thombashi/eoe)](https://goreportcard.com/report/github.com/thombashi/eoe)


## Usage

```go
package main

import (
    "github.com/thombashi/eoe"
)

func main() {
    logger := slog.Default()
    params := eoe.NewParams().WithLogger(logger)

    // exit the program with an error message with the logger when an error is not nil
    err := someFunction()
    eoe.ExitOnError(err, params.WithMessage("someFunction failed"))

    err = anotherFunction()
    eoe.ExitOnError(err, params.WithMessage("anotherFunction failed"))
}
```
