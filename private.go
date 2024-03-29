package trackerr

import (
	"fmt"
)

func fmtMsg(msg string, args ...any) string {
	return fmt.Sprintf(msg, args...)
}

func because(msg string, args ...any) *UntrackedError {
	return &UntrackedError{
		msg: fmtMsg(msg, args...),
	}
}

func causedBy(cause error, msg string, args ...any) *UntrackedError {
	return &UntrackedError{
		msg:   fmtMsg(msg, args...),
		cause: cause,
	}
}
