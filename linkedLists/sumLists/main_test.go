package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	one linkedList
	two linkedList
	sum linkedList
}

func Test(t *testing.T) {
	testCase := getTestCases()

	for _, tc := range testCase {
		t.Run(fmt.Sprintf("%s+%s", formatLL(tc.one), formatLL(tc.two)), func(t *testing.T) {
			assert.Equal(t, tc.sum, sum(tc.one, tc.two), fmt.Sprintf("EXPECTED: %s", formatLL(tc.sum)))
		})
	}
}

func getTestCases() []testCase {
	testCases := []testCase{}

	a := node{
		data: 7,
		next: &node{
			data: 1,
			next: &node{
				data: 6,
				next: nil,
			},
		},
	}

	aa := node{
		data: 4,
		next: &node{
			data: 3,
			next: &node{
				data: 2,
				next: &node{
					data: 1,
				},
			},
		},
	}

	b := node{
		data: 5,
		next: &node{
			data: 9,
			next: &node{
				data: 2,
				next: nil,
			},
		},
	}

	bb := node{
		data: 0,
		next: &node{
			data: 9,
			next: &node{
				data: 5,
			},
		},
	}

	s := node{
		data: 2,
		next: &node{
			data: 1,
			next: &node{
				data: 9,
				next: nil,
			},
		},
	}

	zero := node{
		data: 0,
		next: nil,
	}

	testCases = append(testCases, testCase{linkedList{head: &a}, linkedList{head: &b}, linkedList{head: &s}})
	testCases = append(testCases, testCase{linkedList{head: &a}, linkedList{head: nil}, linkedList{head: &a}})
	testCases = append(testCases, testCase{linkedList{head: nil}, linkedList{head: &b}, linkedList{head: &b}})
	testCases = append(testCases, testCase{linkedList{head: nil}, linkedList{head: nil}, linkedList{head: nil}})
	testCases = append(testCases, testCase{linkedList{head: &a}, linkedList{head: &zero}, linkedList{head: &a}})
	testCases = append(testCases, testCase{linkedList{head: &zero}, linkedList{head: &a}, linkedList{head: &a}})
	testCases = append(testCases, testCase{linkedList{head: &zero}, linkedList{head: &b}, linkedList{head: &b}})
	testCases = append(testCases, testCase{linkedList{head: &s}, linkedList{head: &zero}, linkedList{head: &s}})
	testCases = append(testCases, testCase{linkedList{head: &node{data: 7}}, linkedList{head: &node{data: 9}}, linkedList{head: &node{data: 6, next: &node{data: 1}}}})
	testCases = append(testCases, testCase{linkedList{head: &a}, linkedList{head: &a}, linkedList{head: &aa}})
	testCases = append(testCases, testCase{linkedList{head: &b}, linkedList{head: &b}, linkedList{head: &bb}})

	return testCases
}
