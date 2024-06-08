package eoe

import (
	"context"
	"log/slog"
	"os"
)

// ExitOnErrorParams is a struct for ExitOnError function.
type ExitOnErrorParams struct {
	// Message is an error message to print.
	Message string

	// ExitCode is an exit code to use when exiting the program.
	// Default is 1.
	ExitCode int

	// Logger is a logger to use for logging.
	// Default is slog.Default().
	Logger *slog.Logger

	// LogLevel is a log level to use for logging.
	// Default is slog.LevelError.
	LogLevel slog.Level

	// Context is a context to use for logging.
	// Default is context.Background().
	Context context.Context

	// ExitFunc is a function to use for exiting the program.
	ExitFunc func(params *ExitOnErrorParams)
}

func exitFunc(params *ExitOnErrorParams) {
	os.Exit(params.ExitCode)
}

// NewParams creates a new ExitOnErrorParams instance.
func NewParams() *ExitOnErrorParams {
	return &ExitOnErrorParams{
		ExitCode: 1,
		LogLevel: slog.LevelError,
		Context:  context.Background(),
		ExitFunc: exitFunc,
	}
}

// WithMessage sets the error message.
func (params *ExitOnErrorParams) WithMessage(msg string) *ExitOnErrorParams {
	params.Message = msg
	return params
}

// WithExitCode sets the exit code.
func (params *ExitOnErrorParams) WithExitCode(code int) *ExitOnErrorParams {
	params.ExitCode = code
	return params
}

// WithLogger sets the logger.
func (params *ExitOnErrorParams) WithLogger(logger *slog.Logger) *ExitOnErrorParams {
	params.Logger = logger
	return params
}

// WithContext sets the context.
func (params *ExitOnErrorParams) WithContext(ctx context.Context) *ExitOnErrorParams {
	params.Context = ctx
	return params
}

// WithLogLevel sets the log level.
func (params *ExitOnErrorParams) WithLogLevel(level slog.Level) *ExitOnErrorParams {
	params.LogLevel = level
	return params
}

func (params *ExitOnErrorParams) WithExitFunc(exitFunc func(params *ExitOnErrorParams)) *ExitOnErrorParams {
	params.ExitFunc = exitFunc
	return params
}

// ExitOnError prints an error message and exits the program with the specified exit code.
func ExitOnError(err error, params *ExitOnErrorParams) {
	if err == nil {
		return
	}

	var logger *slog.Logger
	if params.Logger != nil {
		logger = params.Logger
	} else {
		logger = slog.Default()
	}

	if params.Message != "" {
		logger.Log(params.Context, params.LogLevel, params.Message, slog.Any("error", err))
	} else {
		logger.Log(params.Context, params.LogLevel, err.Error())
	}

	params.ExitFunc(params)
}
