// Package eoe provides a function to exit the program with an error log when a passing error is not nil.
//
// The ExitOnError function is useful when you want to exit the program when an error argument is not nil.
// The function prints an error message with a slog.Logger and exits the program with the specified exit code.
// You can customize the behavior of the function by using the ExitOnErrorParams struct:
//
//   - error message
//   - exit code
//   - slog.Logger for logging
//   - log level for logging
//   - context.Context
//   - function to call when exiting
package eoe
