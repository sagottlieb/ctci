package main

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	str1   string
	str2           string
	expectation   bool
}

func Test(t *testing.T) {
	testCase := getTestCases()

	for _, tc := range testCase {
		t.Run(fmt.Sprintf("%s-%s", tc.str1, tc.str2), func(t *testing.T) {
			assert.Equal(t, tc.expectation, checkOneAwayRecursive(tc.str1, tc.str2))
		})
	}
}


func getTestCases() []testCase {
	testCases := []testCase{}

	testCases = append(testCases, testCase{"pale", "ple", true})
	testCases = append(testCases, testCase{"pales", "pale", true})
	testCases = append(testCases, testCase{"pale", "bale", true})
	testCases = append(testCases, testCase{"pale", "bake", false})
	testCases = append(testCases, testCase{"pale", "pal", true})
	testCases = append(testCases, testCase{"ale", "pale", true})
	testCases = append(testCases, testCase{"pal", "pale", true})
	testCases = append(testCases, testCase{"pale", "ppale", true})
	testCases = append(testCases, testCase{"pale", "palee", true})
	testCases = append(testCases, testCase{"pale", "paleee", false})
	testCases = append(testCases, testCase{"a", "abc", false})
	testCases = append(testCases, testCase{"", "ab", false})
	testCases = append(testCases, testCase{"", "x", true})

	switchedCases := []testCase{}

	for _, t := range testCases {
		switchedCases = append(switchedCases, testCase{t.str2, t.str1, t.expectation})
	}

	testCases = append(testCases, switchedCases...)

	return testCases
}
