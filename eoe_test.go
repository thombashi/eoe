package eoe

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	envExitTriggerKey   = "EXIT_ON_ERROR_TEST"
	envExitTriggerValue = "1"
)

func TestExitOnError(t *testing.T) {
	// should not exit if the error is nil
	ExitOnError(nil, NewParams())
}

func TestExitOnErrorExitWithDefaultParams(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	const (
		errorMsg = "TestExitOnErrorExitWithDefaultParams"
	)

	if os.Getenv(envExitTriggerKey) == envExitTriggerValue {
		params := NewParams()
		ExitOnError(errors.New(errorMsg), params)
		t.Fatalf("ExitOnError did not exit")
	}

	var stderr bytes.Buffer
	cmd := exec.Command(os.Args[0], "-test.run=TestExitOnErrorExit")
	cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", envExitTriggerKey, envExitTriggerValue))
	cmd.Stderr = &stderr
	err := cmd.Run()

	e, ok := err.(*exec.ExitError)
	r.True(ok)
	a.False(e.Success())
	a.Equal(1, e.ExitCode())
	a.Contains(stderr.String(), errorMsg)
}

func TestExitOnErrorExitWithCustomParams(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	const (
		exitMsg  = "ExitOnError with an error"
		exitCode = 2
		errorMsg = "TestExitOnErrorExitWithCustomParams"
	)

	if os.Getenv(envExitTriggerKey) == envExitTriggerValue {
		params := NewParams().WithMessage(exitMsg).WithExitCode(exitCode).WithLogLevel(slog.LevelWarn)
		ExitOnError(errors.New(errorMsg), params)
		t.Fatalf("ExitOnError did not exit")
	}

	var stderr bytes.Buffer
	cmd := exec.Command(os.Args[0], "-test.run=TestExitOnErrorExitWithCustomParams")
	cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", envExitTriggerKey, envExitTriggerValue))
	cmd.Stderr = &stderr
	err := cmd.Run()

	e, ok := err.(*exec.ExitError)
	r.True(ok)
	a.False(e.Success())
	a.Equal(exitCode, e.ExitCode())
	a.Contains(stderr.String(), exitMsg)
	a.Contains(stderr.String(), errorMsg)
	a.Contains(stderr.String(), "WARN")
}

func TestExitOnErrorExitWithExitFunc(t *testing.T) {
	a := assert.New(t)

	var exitFuncCalled bool
	ExitOnError(errors.New("error"), NewParams().WithExitFunc(func(params *ExitOnErrorParams) {
		t.Log("called exit func")
		exitFuncCalled = true
	}))
	a.True(exitFuncCalled)
}
