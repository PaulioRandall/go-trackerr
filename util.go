package trackable

import (
	"errors"
	"fmt"
	"strings"
)

// IsTracked returns true if the error has a trackable ID greater than zero.
func IsTracked(e error) bool {
	te, ok := e.(*trackable)
	return ok && te.id > 0
}

// Is is an alias for errors.Is.
func Is(e, target error) bool {
	return errors.Is(e, target)
}

// All returns true only if errors.Is returns true for all targets.
func All(e error, targets ...error) bool {
	for _, t := range targets {
		if !errors.Is(e, t) {
			return false
		}
	}
	return true
}

// Any returns true if errors.Is returns true for any of the targets.
func Any(e error, targets ...error) bool {
	for _, t := range targets {
		if errors.Is(e, t) {
			return true
		}
	}
	return false
}

// Debug is convenience for fmt.Println("[Debug error]\n", ErrorStack(e)).
func Debug(e error) (int, error) {
	s := ErrorStack(e)

	if s == "" {
		return fmt.Print("[Debug error] nil error")
	}

	return fmt.Print("[Debug error]\n", s)
}

// ErrorStack is convenience for
// 		ErrorStackWith(e, "  ", "\n⤷ ", "\n——Interface: ", "\n").
//
// Output example:
//		  Failed to execuate packages
//		⤷ Could not do that thing
//		⤷ API returned an error
//		⤷ UnhappyAPI returned an error
//		——Interface——
//		⤷ This is the error wrapped at the API boundary
//		⤷ This is the root cause
func ErrorStack(e error) string {
	return ErrorStackWith(e, "  ", "\n⤷ ", "\n——Interface: ", "\n")
}

// ErrorStackWith returns a human readable representation of the error stack.
//
// Given:
// 		ErrorStackWith(e, "  ", "\n⤷ ", "\n——Interface: ", "\n").
//
// Output example:
//		  Failed to execuate packages
//		⤷ Could not do that thing
//		⤷ API returned an error
//		⤷ UnhappyAPI returned an error
//		——Interface: UnhappyAPI
//		⤷ This is the error wrapped at the API boundary
//		⤷ This is the root cause
func ErrorStackWith(e error, prefix, delim, ifaceDelim, suffix string) string {
	sb := strings.Builder{}
	sb.WriteString(prefix)

	for i, cause := range AsStack(e) {
		if i > 0 {
			sb.WriteString(delim)
		}

		s := ErrorWithoutCause(cause)
		if s != "" {
			sb.WriteString(s)
		}

		if iface := InterfaceName(cause); iface != "" {
			sb.WriteString(ifaceDelim)
			sb.WriteString(iface)
		}
	}

	sb.WriteString(suffix)
	return sb.String()
}

// AsStack recursively unwraps the error returning a slice of errors.
//
// The passed error will be first and root cause last.
func AsStack(e error) []error {
	var stack []error

	for ; e != nil; e = errors.Unwrap(e) {
		stack = append(stack, e)
	}

	return stack
}

// ErrorWithoutCause removes the cause from error messages that use the
// standard concaternation.
//
// The standard concaternation being in the format '%s: %w' where s is the
// error message and w is the cause's message.
func ErrorWithoutCause(e error) string {
	if stringer, ok := e.(fmt.Stringer); ok {
		return stringer.String()
	}

	cause := errors.Unwrap(e)
	s := e.Error()

	if cause == nil {
		return s
	}

	s = strings.TrimSuffix(s, cause.Error())
	s = strings.TrimSpace(s)
	return strings.TrimSuffix(s, ":")
}

// InterfaceName returns the name of the interface where the error occurred if
// one is available.
func InterfaceName(e error) string {
	type iface interface {
		InterfaceName() string
	}

	if v, ok := e.(iface); ok {
		return v.InterfaceName()
	}
	return ""
}
