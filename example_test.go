package eoe_test

import (
	"errors"
	"log/slog"

	eoe "github.com/thombashi/eoe"
)

func successFunc() error {
	return nil
}

func errrorFunc() error {
	return errors.New("an error occurred")
}

func ExampleExitOnError() {
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
