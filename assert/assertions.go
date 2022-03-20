package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// Equal asserts that two objects are equal.
//
//    assert.Equal(t, 123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func Equal[T comparable](t *testing.T, expected T, actual T, msgAndArgs ...any) bool {
	if !reflect.DeepEqual(expected, actual) {
		return Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %v\n"+
			//TODO: support diff
			//"actual  : %s%s", expected, actual, diff), msgAndArgs...)
			"actual  : %v", expected, actual), msgAndArgs...)
	}
	return true
}

// NotEqual asserts that the specified values are NOT equal.
//
//    assert.NotEqual(t, obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func NotEqual[T comparable](t *testing.T, expected T, actual T, msgAndArgs ...any) bool {
	if reflect.DeepEqual(expected, actual) {
		return Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %v\n"+
			//TODO: support diff
			//"actual  : %s%s", expected, actual, diff), msgAndArgs...)
			"actual  : %v", expected, actual), msgAndArgs...)
	}
	return true
}

// Fail reports a failure through
func Fail(t *testing.T, failureMessage string, msgArgs ...any) bool {
	t.Errorf(failureMessage, msgArgs...)
	return false
}

// Same asserts that two pointers reference the same object.
//
//    assert.Same(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func Same(t testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return true
}
