package stackMin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	s   stack
	min int
}

func Test(t *testing.T) {
	testCase := getTestCases()

	for _, tc := range testCase {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tc.min, tc.s.min())
		})
	}
}

func getTestCases() []testCase {
	testCases := []testCase{}

	a := stack{
		top: &node{
			val: 1,
			min: 1,
		},
	}

	b := stack{
		top: &node{
			val: 1,
			min: 1,
			next: &node{
				val: 2,
				min: 2,
			},
		},
	}

	c := stack{
		top: &node{
			val: 2,
			min: 1,
			next: &node{
				val: 1,
				min: 1,
			},
		},
	}

	d := stack{
		top: &node{
			val: 2,
			min: 1,
			next: &node{
				val: 1,
				min: 1,
				next: &node{
					val: 3,
					min: 3,
				},
			},
		},
	}

	e := stack{
		top: &node{
			val:  -5,
			min:  -5,
			next: d.top,
		},
	}

	f := stack{
		top: &node{
			val: 2,
			min: 1,
			next: &node{
				val: 1,
				min: 1,
				next: &node{
					val: 5,
					min: 4,
					next: &node{
						val: 4,
						min: 4,
					},
				},
			},
		},
	}

	testCases = append(testCases, testCase{a, 1})
	testCases = append(testCases, testCase{b, 1})
	testCases = append(testCases, testCase{c, 1})
	testCases = append(testCases, testCase{d, 1})
	testCases = append(testCases, testCase{e, -5})
	testCases = append(testCases, testCase{f, 1})

	return testCases
}
