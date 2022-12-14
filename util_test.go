package trackerr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	untrackedAlpha   = mockUntracked(nil, "untracked alpha")
	untrackedBeta    = mockUntracked(nil, "untracked beta")
	untrackedCharlie = mockUntracked(nil, "untracked charlie")

	trackedAlpha   = mockTracked(1, nil, "tracked alpha")
	trackedBeta    = mockTracked(2, nil, "tracked beta")
	trackedCharlie = mockTracked(3, nil, "tracked charlie")
)

func mockUntracked(cause error, msg string, args ...any) *UntrackedError {
	return &UntrackedError{
		cause: cause,
		msg:   fmtMsg(msg, args...),
	}
}

func mockTracked(id int, cause error, msg string, args ...any) *TrackedError {
	return &TrackedError{
		id:    id,
		msg:   fmtMsg(msg, args...),
		cause: cause,
	}
}

func Test_IsTracked_1(t *testing.T) {
	require.True(t, IsTracked(trackedAlpha))
}

func Test_IsTracked_2(t *testing.T) {
	require.False(t, IsTracked(untrackedAlpha))
}

func Test_HasTracked_1(t *testing.T) {
	c := trackedCharlie
	b := untrackedBeta.Wrap(c)
	a := untrackedAlpha.Wrap(b)

	require.True(t, HasTracked(a))
}

func Test_HasTracked_2(t *testing.T) {
	c := untrackedCharlie
	b := untrackedBeta.Wrap(c)
	a := untrackedAlpha.Wrap(b)

	require.False(t, HasTracked(a))
}

func Test_All_1(t *testing.T) {
	c := untrackedCharlie
	b := untrackedBeta.Wrap(c)
	a := untrackedAlpha.Wrap(b)

	e := a

	require.True(t, All(e))
	require.True(t, All(e, a, b, c))

	require.False(t, All(e, a, b, trackedCharlie))
}

func Test_Allordered_1(t *testing.T) {
	cu := untrackedCharlie
	ct := trackedCharlie.Wrap(cu)
	bu := untrackedBeta.Wrap(ct)
	bt := trackedBeta.Wrap(bu)
	au := untrackedAlpha.Wrap(bt)
	at := trackedAlpha.Wrap(au)

	e := at

	require.True(t, AllOrdered(e))

	require.True(t, AllOrdered(e, at, au))
	require.False(t, AllOrdered(e, au, at)) // Reversed

	require.True(t, AllOrdered(e, at, au, bt, bu, ct, cu))
	require.False(t, AllOrdered(e, at, bt, bu, au, ct, cu)) // Badly placed au
	require.False(t, AllOrdered(e, at, au, bt, bu, cu, ct)) // Swapped last two
}

func Test_Any_1(t *testing.T) {
	c := untrackedCharlie
	b := untrackedBeta.Wrap(c)
	a := untrackedAlpha.Wrap(b)

	e := a

	require.True(t, Any(e, a))
	require.True(t, Any(e, a, b))
	require.True(t, Any(e, a, b, c))
	require.True(t, Any(e, b, c))
	require.True(t, Any(e, c))

	require.False(t, Any(e))
	require.False(t, Any(e, trackedAlpha, trackedBeta, trackedCharlie))
}
