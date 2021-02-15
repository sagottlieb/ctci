package main

import (
	"fmt"
	"testing"

	"github.com/sagottlieb/ctci/linkedLists/sumLists/datastructs"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	one datastructs.LinkedList
	two datastructs.LinkedList
	sum datastructs.LinkedList
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

	a := datastructs.Node{
		Data: 7,
		Next: &datastructs.Node{
			Data: 1,
			Next: &datastructs.Node{
				Data: 6,
				Next: nil,
			},
		},
	}

	aa := datastructs.Node{
		Data: 4,
		Next: &datastructs.Node{
			Data: 3,
			Next: &datastructs.Node{
				Data: 2,
				Next: &datastructs.Node{
					Data: 1,
				},
			},
		},
	}

	b := datastructs.Node{
		Data: 5,
		Next: &datastructs.Node{
			Data: 9,
			Next: &datastructs.Node{
				Data: 2,
				Next: nil,
			},
		},
	}

	bb := datastructs.Node{
		Data: 0,
		Next: &datastructs.Node{
			Data: 9,
			Next: &datastructs.Node{
				Data: 5,
			},
		},
	}

	s := datastructs.Node{
		Data: 2,
		Next: &datastructs.Node{
			Data: 1,
			Next: &datastructs.Node{
				Data: 9,
				Next: nil,
			},
		},
	}

	zero := datastructs.Node{
		Data: 0,
		Next: nil,
	}

	testCases = append(testCases, testCase{datastructs.LinkedList{Head: &a}, datastructs.LinkedList{Head: &b}, datastructs.LinkedList{Head: &s}})
	testCases = append(testCases, testCase{datastructs.LinkedList{Head: &a}, datastructs.LinkedList{Head: nil}, datastructs.LinkedList{Head: &a}})
	testCases = append(testCases, testCase{datastructs.LinkedList{Head: nil}, datastructs.LinkedList{Head: &b}, datastructs.LinkedList{Head: &b}})
	testCases = append(testCases, testCase{datastructs.LinkedList{Head: nil}, datastructs.LinkedList{Head: nil}, datastructs.LinkedList{Head: nil}})
	testCases = append(testCases, testCase{datastructs.LinkedList{Head: &a}, datastructs.LinkedList{Head: &zero}, datastructs.LinkedList{Head: &a}})
	testCases = append(testCases, testCase{datastructs.LinkedList{Head: &zero}, datastructs.LinkedList{Head: &a}, datastructs.LinkedList{Head: &a}})
	testCases = append(testCases, testCase{datastructs.LinkedList{Head: &zero}, datastructs.LinkedList{Head: &b}, datastructs.LinkedList{Head: &b}})
	testCases = append(testCases, testCase{datastructs.LinkedList{Head: &s}, datastructs.LinkedList{Head: &zero}, datastructs.LinkedList{Head: &s}})
	testCases = append(testCases, testCase{datastructs.LinkedList{Head: &datastructs.Node{Data: 7}}, datastructs.LinkedList{Head: &datastructs.Node{Data: 9}}, datastructs.LinkedList{Head: &datastructs.Node{Data: 6, Next: &datastructs.Node{Data: 1}}}})
	testCases = append(testCases, testCase{datastructs.LinkedList{Head: &a}, datastructs.LinkedList{Head: &a}, datastructs.LinkedList{Head: &aa}})
	testCases = append(testCases, testCase{datastructs.LinkedList{Head: &b}, datastructs.LinkedList{Head: &b}, datastructs.LinkedList{Head: &bb}})

	return testCases
}
