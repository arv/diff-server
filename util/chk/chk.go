// Package chk implements a handful of runtime assertions.
package chk

import "fmt"

func Fail(msg string, params ...interface{}) {
	panic(fmt.Sprintf(msg, params...))
}

func True(cond bool, msg string, params ...interface{}) {
	if !cond {
		Fail(msg, params...)
	}
}

func False(cond bool, msg string, params ...interface{}) {
	True(!cond, msg, params...)
}

func Equal(expected interface{}, actual interface{}) {
	if expected != actual {
		Fail("Expected %#v, got: %#v", expected, actual)
	}
}

func NotNil(v interface{}, msgAndParams ...interface{}) {
	if v == nil {
		var msg string
		if len(msgAndParams) > 0 {
			msg = fmt.Sprintf(msgAndParams[0].(string), msgAndParams[1:]...)
		} else {
			msg = "Expected non-nil value"
		}
		Fail(msg)
	}
}

func NoError(err error) {
	if err != nil {
		Fail("Unexpected error: %#v", err)
	}
}
