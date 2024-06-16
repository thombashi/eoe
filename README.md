# eoe

Provide a simple function to exit the program with an error message on errors in Go.

[![Go Reference](https://pkg.go.dev/badge/github.com/thombashi/eoe.svg)](https://pkg.go.dev/github.com/thombashi/eoe)
[![Go Report Card](https://goreportcard.com/badge/github.com/thombashi/eoe)](https://goreportcard.com/report/github.com/thombashi/eoe)
[![CI](https://github.com/thombashi/eoe/actions/workflows/ci.yaml/badge.svg)](https://github.com/thombashi/eoe/actions/workflows/ci.yaml)
[![CodeQL](https://github.com/thombashi/eoe/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/thombashi/eoe/actions/workflows/github-code-scanning/codeql)


## Installation

```
go get -u github.com/thombashi/eoe
```


## Usage

```go
package main

import (
    "github.com/thombashi/eoe"
)

func successFunc() error {
	return nil
}

func errrorFunc() error {
	return errors.New("an error occurred")
}

func main() {
	var err error
	logger := slog.Default()
	params := eoe.NewParams().WithLogger(logger)

	// should not exit if the error is nil
	err = successFunc()
	eoe.ExitOnError(err, params.WithMessage("should not exit"))

	// should exit if the error is not nil
	err = errrorFunc()
	eoe.ExitOnError(err, params.WithMessage("should exit with an error message"))
}
```
