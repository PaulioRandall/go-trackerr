package trackerr

var (
	globalRealm       IntRealm
	globalInitialised bool
)

// Initialised causes all future calls to New or Track to panic.
//
// When called from an init function in the main package, it prevents creation
// of trackable errors after program initialisation.
//
//		package main
//
//		import "github.com/PaulioRandall/go-trackerr"
//
//		func init() {
//			trackerr.Initialised()
//		}
func Initialised() {
	globalInitialised = true
}

func checkInitState() {
	if globalInitialised {
		panic(Untracked("No tracked errors may be created after initialisation."))
	}
}

// Realm represents a space where each trackable error (stack trace node)
// has its own unique ID.
//
// There is a private package scooped Realm that will suffice for most
// purposes. It is used via the package scooped Track functions.
//
// Receiving functions are designed to be called during package initialisation.
// This means it should only be used to initialise package global variables and
// within init functions. The exception is where Realms are in use.
//
// Furthermore, all functions return a shallow copy of any passed or
// receiving errors creating a somewhat immutability based ecosystem.
//
// This interface is primarily for documentation.
type Realm interface {

	// New is an alias for Track.
	New(msg string, args ...any) *TrackedError

	// Track returns a new tracked error, that is, one with a tracking ID.
	Track(msg string, args ...any) *TrackedError
}

// IntRealm is a Realm that uses a simple incrementing integer field as the
// pool of unique IDs.
//
//		realm := IntRealm{}
//
// The recommended way to use this package is to ignore this struct and use the
// New or Track package functions. If this package's API is used as intended
// then it would be impossible to cause an integer overflow scenario in any
// real world use case. However, Realms were conceived for such an event
// and for those who really hate the idea of relying on a singleton they have
// no control over.
type IntRealm struct {
	idPool *int
}

// New is an alias for Track.
func (r *IntRealm) New(msg string, args ...any) *TrackedError {
	return Track(msg, args...)
}

// Track returns a new tracked error.
//
// Calls to HasTracked, IsTracked, and IsTrackerr will all return true when
// the error is passed to them.
func (r *IntRealm) Track(msg string, args ...any) *TrackedError {
	return &TrackedError{
		id:  r.newID(),
		msg: fmtMsg(msg, args...),
	}
}

func (r *IntRealm) newID() int {
	if r.idPool == nil {
		n := 0
		r.idPool = &n
	}

	(*r.idPool)++
	return *r.idPool
}
