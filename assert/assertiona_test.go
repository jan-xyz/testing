package assert_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jan-xyz/testing/assert"
)

func TestEqualComparesSimpleTypes(t *testing.T) {
	mockT := new(testing.T)

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	cases := []testCase[int]{
		{
			expected: 1,
			actual:   1,
			result:   true,
		},
		{
			expected: 1,
			actual:   0,
			result:   false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.Equal(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestEqualComparesPointer(t *testing.T) {
	mockT := new(testing.T)

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	cases := []testCase[*int]{
		{
			expected: pointer(1),
			actual:   pointer(1),
			result:   true,
		},
		{
			expected: pointer(1),
			actual:   pointer(0),
			result:   false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.Equal(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestEqualComparesStructs(t *testing.T) {
	mockT := new(testing.T)

	type testInput struct {
		field int
	}

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	cases := []testCase[*testInput]{
		{
			expected: &testInput{field: 1},
			actual:   &testInput{field: 1},
			result:   true,
		},
		{
			expected: &testInput{field: 1},
			actual:   &testInput{field: 0},
			result:   false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.Equal(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestNotEqualComparesSimpleTypes(t *testing.T) {
	mockT := new(testing.T)

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	cases := []testCase[int]{
		{
			expected: 1,
			actual:   1,
			result:   false,
		},
		{
			expected: 1,
			actual:   0,
			result:   true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.NotEqual(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestNotEqualComparesPointer(t *testing.T) {
	mockT := new(testing.T)

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	cases := []testCase[*int]{
		{
			expected: pointer(1),
			actual:   pointer(1),
			result:   false,
		},
		{
			expected: pointer(1),
			actual:   pointer(0),
			result:   true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.NotEqual(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", *c.expected, *c.actual, c.result, c.remark)
			}
		})
	}
}

func TestNotEqualComparesStructs(t *testing.T) {
	mockT := new(testing.T)

	type testInput struct {
		field int
	}

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	cases := []testCase[*testInput]{
		{
			expected: &testInput{field: 1},
			actual:   &testInput{field: 1},
			result:   false,
		},
		{
			expected: &testInput{field: 1},
			actual:   &testInput{field: 0},
			result:   true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.NotEqual(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestSamComparesSimpleTypes(t *testing.T) {
	mockT := new(testing.T)

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	cases := []testCase[int]{
		{
			expected: 1,
			actual:   1,
			result:   true,
		},
		{
			expected: 1,
			actual:   0,
			result:   false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.Same(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestSameComparesPointer(t *testing.T) {
	mockT := new(testing.T)

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	samePtr := pointer(3)

	cases := []testCase[*int]{
		{
			expected: samePtr,
			actual:   samePtr,
			result:   true,
		},
		{
			expected: pointer(1),
			actual:   pointer(1),
			result:   false,
		},
		{
			expected: pointer(1),
			actual:   pointer(0),
			result:   false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.Same(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestSameComparesStructs(t *testing.T) {
	mockT := new(testing.T)

	type testInput struct {
		field int
	}

	samePtr := &testInput{field: 3}

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	cases := []testCase[*testInput]{
		{
			expected: samePtr,
			actual:   samePtr,
			result:   true,
		},
		{
			expected: &testInput{field: 1},
			actual:   &testInput{field: 1},
			result:   false,
		},
		{
			expected: &testInput{field: 1},
			actual:   &testInput{field: 0},
			result:   false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.Same(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestNotSameComparesSimpleTypes(t *testing.T) {
	mockT := new(testing.T)

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	cases := []testCase[int]{
		{
			expected: 1,
			actual:   1,
			result:   false,
		},
		{
			expected: 1,
			actual:   0,
			result:   true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.NotSame(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestNotSameComparesPointer(t *testing.T) {
	mockT := new(testing.T)

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	samePtr := pointer(3)

	cases := []testCase[*int]{
		{
			expected: samePtr,
			actual:   samePtr,
			result:   false,
		},
		{
			expected: pointer(1),
			actual:   pointer(1),
			result:   true,
		},
		{
			expected: pointer(1),
			actual:   pointer(0),
			result:   true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.NotSame(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", *c.expected, *c.actual, c.result, c.remark)
			}
		})
	}
}

func TestNotSameComparesStructs(t *testing.T) {
	mockT := new(testing.T)

	type testInput struct {
		field int
	}

	samePtr := &testInput{field: 1}

	type testCase[T comparable] struct {
		expected T
		actual   T
		result   bool
		remark   string
	}

	cases := []testCase[*testInput]{
		{
			expected: samePtr,
			actual:   samePtr,
			result:   false,
		},
		{
			expected: &testInput{field: 1},
			actual:   &testInput{field: 1},
			result:   true,
		},
		{
			expected: &testInput{field: 1},
			actual:   &testInput{field: 0},
			result:   true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := assert.NotSame(mockT, c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestNil(t *testing.T) {
	mockT := new(testing.T)

	type testCase[T comparable] struct {
		f      func() error
		result bool
		remark string
	}

	cases := []testCase[*struct{ field int }]{
		{
			f:      func() error { return errors.New("some error") },
			result: false,
		},
		{
			f:      func() error { return nil },
			result: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", nil, c.f()), func(t *testing.T) {
			err := c.f()
			res := assert.Nil(mockT, err)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", nil, err, c.result, c.remark)
			}
		})
	}
}

func TestNotNil(t *testing.T) {
	mockT := new(testing.T)

	type testCase struct {
		actual any
		result bool
		remark string
	}

	cases := []testCase{
		{
			actual: errors.New("some error"),
			result: true,
		},
		{
			actual: nil,
			result: false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", nil, c.actual), func(t *testing.T) {
			res := assert.NotNil(mockT, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", nil, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestLen(t *testing.T) {
	mockT := new(testing.T)

	type testCase struct {
		actual []int
		len    int
		result bool
		remark string
	}

	cases := []testCase{
		{
			actual: []int{1, 2, 3},
			len:    7,
			result: false,
		},
		{
			actual: []int{1, 2, 3},
			len:    3,
			result: true,
		},
		{
			actual: nil,
			result: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", nil, c.actual), func(t *testing.T) {
			res := assert.Len(mockT, c.actual, c.len)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", nil, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestLenMap(t *testing.T) {
	mockT := new(testing.T)

	type testCase struct {
		actual map[string]int
		len    int
		result bool
		remark string
	}

	cases := []testCase{
		{
			actual: map[string]int{"foo": 1, "bar": 2, "baz": 3},
			len:    7,
			result: false,
		},
		{
			actual: map[string]int{"foo": 1, "bar": 2, "baz": 3},
			len:    3,
			result: true,
		},
		{
			actual: nil,
			result: true,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", nil, c.actual), func(t *testing.T) {
			res := assert.LenMap(mockT, c.actual, c.len)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", nil, c.actual, c.result, c.remark)
			}
		})
	}
}

func pointer[T any](input T) *T {
	return &input
}
