package main

import (
	"fmt"
	"testing"

	"github.com/sagottlieb/ctci/arraysAndStrings/isUnique/bitVector"
	"github.com/sagottlieb/ctci/arraysAndStrings/isUnique/booleanVector"
	"github.com/sagottlieb/ctci/arraysAndStrings/isUnique/bruteForce"
	"github.com/sagottlieb/ctci/arraysAndStrings/isUnique/hashmap"
	"github.com/sagottlieb/ctci/arraysAndStrings/isUnique/sortFirst"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	description   string
	str           string
	expectation   bool
	chckr         checker
	skipBenchmark bool // some cases only for correctness
}

func Test(t *testing.T) {
	testCase := getTestCases()

	bruteChecker := bruteForce.New()
	mapChecker := hashmap.New()
	sortedChecker := sortFirst.New()
	booleanChecker := booleanVector.New()
	bitVectorChecker := bitVector.New()

	for _, tc := range testCase {
		t.Run(tc.str, func(t *testing.T) {
			assert.Equal(t, tc.expectation, bruteChecker.Check(tc.str))
			assert.Equal(t, tc.expectation, mapChecker.Check(tc.str))
			assert.Equal(t, tc.expectation, sortedChecker.Check(tc.str))
			assert.Equal(t, tc.expectation, booleanChecker.Check(tc.str))
			assert.Equal(t, tc.expectation, bitVectorChecker.Check(tc.str))
		})
	}
}

func Benchmark(b *testing.B) {
	testCase := getTestCases()

	for _, tc := range testCase {

		if tc.skipBenchmark {
			continue
		}

		b.Run(fmt.Sprintf("%s_%s", tc.chckr.GetName(), tc.description), func(b *testing.B) {

			for i := 0; i < b.N; i++ {
				tc.chckr.Check(tc.str)
			}

		})
	}
}

type details struct {
	exp       bool
	desc      string
	skipBench bool
}

func getTestCases() []testCase {

	testStrings := map[string]details{
		"abcdefg":  {exp: true},
		"abcdefga": {exp: false},
		"999g":     {exp: false, skipBench: true},
		"abcdefgA": {exp: true, skipBench: true},
		"1233":     {exp: false, skipBench: true},
		"vcdfeET25HIJAKL0987]1!@#$":                                                                     {exp: true},
		"1234567890!@#$%^&*()-_=+`~qwertyuiop[]QWERTYUIOP{}|asdfghjkl;'ASDFGHJKL:zxcvbnm,./ZXCVBNM<>?":  {exp: true, desc: "long,unique"},
		"12234567890!@#$%^&*()-_=+`~qwertyuiop[]QWERTYUIOP{}|asdfghjkl;'ASDFGHJKL:zxcvbnm,./ZXCVBNM<>?": {exp: false, desc: "long,rep at beginning"},
		"1234567890!@#$%^&*()-_=+`~qwertyuiop[]QWERTYUIOP{}|asdfghjkl;'ASDFGHJKL:zxcvbnm,./ZXCVBNM<>?2": {exp: false, desc: "long,rep at beginnning once sorted"},
		"1234567890!@#$%^&*()-_=+`~qwertyuiop[]QWERTYUIOP{}|asdfghjkl;'ASDFGHJKL:zxcvbnm,./ZXCVBNM<>?}": {exp: false, desc: "long, rep at end once sorted"},
	}

	checkers := []checker{
		bruteForce.New(),
		hashmap.New(),
		sortFirst.New(),
		booleanVector.New(),
		bitVector.New(),
	}

	testCases := []testCase{}

	for s, d := range testStrings {
		for _, c := range checkers {
			tc := testCase{str: s, expectation: d.exp, chckr: c}
			if d.desc != "" {
				tc.description = d.desc
			} else {
				tc.description = s
			}
			if d.skipBench {
				tc.skipBenchmark = true
			}
			testCases = append(testCases, tc)
		}
	}

	return testCases
}
