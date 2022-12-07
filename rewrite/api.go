// Package track aims to facilitate creation of referenceable errors and
// elegant stack traces.
package track

// TODO 1: Write up a realistic example for this interface using test data
// TODO 2: Implement package interface

// TODO: Think about how to integrate file names and line numbers.
// TODO: - How, where, and when to collect them?
// TODO: - How to optimise print outs with them?
// TODO: - May have to redesign the Debug function?

var (
	globalRealm IntRealm

	// ErrTodo is a convenience tracked error for specifying a TODO.
	//
	// This can be useful if you're taking a stepwise refinement or test driven
	// approach to writing code.
	ErrTodo = Error("TODO: Implementation needed")

	// ErrBug is a convenience tracked error for use at the site of known bugs.
	ErrBug = Error("BUG: Fix needed")

	// ErrInsane is a convenience tracked error for sanity checks.
	ErrInsane = Error("Sanity check failed!!")
)

// ErrorWrap represents an error that may or may not have a cause.
//
// This interface is primarily for documentation.
type ErrorWrap interface {
	error

	// Unwrap returns the error's underlying cause or nil if none exists.
	//
	// It is designed to work with the Is function exposed by the standard
	// errors package.
	Unwrap() error

	// Wrap returns a copy of the receiving error with the passed error as the
	// underlying cause.
	Wrap(error) error

	// Copy returns a shallow copy of the error.
	Copy() error
}

// Untracked returns a new error without a tracking ID.
//
// It is no different than using errors.New except it has a handy fmt.Sprintf
// signiture and a few extra receiving functions for any niche use cases one
// may encounter.
func Untracked(msg string, args ...any) *untrackedError {
	return globalRealm.Untracked(msg, args...)
}

// Error returns a new tracked error from this package's singleton Realm.
//
// This is recommended way to use to create all trackable errors outside of
// testing.
func Error(msg string, args ...any) *trackedError {
	return globalRealm.Error(msg, args...)
}

// Checkpoint returns a new trackable checkpoint error from this package's
// singleton Realm.
//
// This is recommended way to use to create all checkpoint errors outside of
// testing.
func Checkpoint(msg string, args ...any) *checkpointError {
	return globalRealm.Checkpoint(msg, args...)
}

// Debug pretty prints the error stack trace to terminal for debugging
// purposes.
//
// If e is nil then a message will be printed indicating so. While this
// function can be used for logging it's not design for such a use case.
func Debug(e error) (int, error) {
	panic("TODO api.Debug")
}

// HasTracked returns true if the error or one of the underlying causes are
// tracked.
//
// This includes only errors created via Error and Checkpoint functions.
func HasTracked(e error) bool {
	panic("TODO api.HasTracked")
}

// IsTracked returns true if the error is being tracked.
//
// This includes only errors created via Error and Checkpoint functions.
func IsTracked(e error) bool {
	panic("TODO api.IsTracked")
}

// IsCheckpoint returns true if the error is a trackable checkpoint.
func IsCheckpoint(e error) bool {
	panic("TODO api.IsCheckpoint")
}

// Is is an alias for errors.Is.
func Is(e, target error) bool {
	panic("TODO api.Is")
}

// All returns true only if errors.Is returns true for all targets.
func All(e error, targets ...error) bool {
	panic("TODO api.All")
}

// Any returns true if errors.Is returns true for any target.
func Any(e error, targets ...error) bool {
	panic("TODO api.Any")
}

// ErrorStack returns a human readable stack trace for the error.
func ErrorStack(e error) string {
	panic("TODO api.ErrorStack")
}

// AsStack recursively unwraps the error returning a slice of errors.
//
// The passed error will be first and root cause last.
func AsStack(e error) []error {
	panic("TODO api.AsStack")
}
