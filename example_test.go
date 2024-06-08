package eoe_test

import (
	"errors"
	"log/slog"

	eoe "github.com/thombashi/eoe"
)

func ExampleExitOnError() {
	logger := slog.Default()

	// should not exit if the error is nil
	eoe.ExitOnError(nil, eoe.NewParams())

	// should exit if the error is not nil
	eoe.ExitOnError(errors.New("error"), eoe.NewParams().WithLogger(logger))
}
