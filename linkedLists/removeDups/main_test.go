package main

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	before linkedList
	after linkedList
}

func Test(t *testing.T) {
	testCase := getTestCases()

	for _, tc := range testCase {
		t.Run(formatLL(tc.before), func(t *testing.T) {
			assert.Equal(t, tc.after, removeDups(tc.before), fmt.Sprintf("EXPECTED: %s", formatLL(tc.after)))
			assert.Equal(t, tc.after, removeDupsNoBuffer(tc.before), fmt.Sprintf("EXPECTED: %s", formatLL(tc.after)))
		})
	}
}

func getTestCases() []testCase {
	testCases := []testCase{}

	v := node{
		data: 0,
		next: &node{
			data: 1,
			next: nil,
		},
	}

	w := node{
		data: 0,
		next: nil,
	}

	x := node{
		data: 0,
		next: &node{
			data: 1,
			next: &node{
				data: 2,
				next: &node{
					data: 3,
					next: nil,
				},
			},
		},
	}

	y := node{
		data: 2,
		next: &x,
	}

	yxddpd := node{
		data: 2,
		next: &node{
			data: 0,
			next: &node{
				data: 1,
				next: &node{
					data: 3,
					next: nil,
				},
			},
		},
	}

	z := node{
		data: 3,
		next: &node{
			data: 1,
			next: &y,
		},
	}

	zyddpd :=  node{
		data: 3,
		next: &node{
			data: 1,
			next: & node{
				data: 2,
				next: &node{
					data: 0,
					next: nil,
				},
			},
		},
	}

	testCases = append(testCases, testCase{linkedList{head: &x}, linkedList{head: &x}})
	testCases = append(testCases, testCase{linkedList{head: &y}, linkedList{head: &yxddpd}})
	testCases = append(testCases, testCase{linkedList{head: &z}, linkedList{head: &zyddpd}})
	testCases = append(testCases, testCase{linkedList{}, linkedList{}})
	testCases = append(testCases, testCase{linkedList{head: &v}, linkedList{head: &v}})
	testCases = append(testCases, testCase{linkedList{head: &w}, linkedList{head: &w}})

	return testCases
}
