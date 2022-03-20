package assert_test

import (
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
	expected := 1
	actual := 1

	assert.Equal(mockT, expected, actual)
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
	expected := 1
	actual := 1

	assert.Equal(mockT, expected, actual)
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
	expected := 1
	actual := 1

	assert.Equal(mockT, expected, actual)
}

func pointer[T any](input T) *T {
	return &input
}
