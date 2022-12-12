package util_test

import (
	"fmt"
	"testing"

	"github.com/ngicks/type-param-common/util"
	"github.com/stretchr/testify/require"
)

func TestMust(t *testing.T) {
	require := require.New(t)

	testRecover := func(shouldPanic bool, must func() any) any {
		defer func() {
			recovered := recover()
			if shouldPanic {
				require.NotNil(recovered)
			} else {
				require.Nil(recovered)
			}
		}()

		return must()
	}

	require.Equal("foo", testRecover(false, func() any { return util.Must("foo", nil) }))
	require.Equal(nil, testRecover(true, func() any { return util.Must("foo", fmt.Errorf("foo")) }))

	require.Equal(
		[2]any{"foo", 12},
		testRecover(
			false,
			func() any {
				val1, val2 := util.Must3("foo", 12, nil)
				return [2]any{val1, val2}
			},
		),
	)
	require.Equal(
		nil,
		testRecover(
			true,
			func() any {
				val1, val2 := util.Must3("foo", 12, fmt.Errorf("bar"))
				return [2]any{val1, val2}
			},
		),
	)
}

func TestEscape(t *testing.T) {
	require := require.New(t)

	require.Equal("foo", *util.Escape("foo"))
	var empty *string
	require.Equal(empty, *util.Escape[*string](nil))
}

func TestAssert(t *testing.T) {
	require := require.New(t)

	var val any = "aaa"

	var asserted string
	var ok bool

	asserted, ok = util.Assert[string](val)
	require.True(ok)
	require.Equal("aaa", asserted)

	val = util.Escape("aaa")
	asserted, ok = util.Assert[string](val)
	require.True(ok)
	require.Equal("aaa", asserted)

	var strP *string
	asserted, ok = util.Assert[string](strP)
	require.True(ok)
	require.Equal("", asserted)

	val = 123
	asserted, ok = util.Assert[string](val)
	require.False(ok)
	require.Equal("", asserted)
}
