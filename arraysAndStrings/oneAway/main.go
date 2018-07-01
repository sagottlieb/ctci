package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Run go test and go test -bench=.")
}

func checkOneAway(s, t string) bool {

	if len(s) == len(t) {
		// replacement case
		// could be equal or differ in one position due to a replacement
		return checkOneReplacement(s, t)
	}

	if math.Abs(float64(len(s)-len(t))) == 1 {
		// insertion into s / removal from t case, OR
		// insertion into t / removal from s case
		return checkOneInsertion(s, t)
	}

	return false
}

func checkOneReplacement(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	foundDiff := false

	for i, charS := range s {

		charT := t[i]

		if charS != int32(charT) {
			if foundDiff { // can't be off in 2 positions
				return false
			}

			// treat current position as the replaced one
			foundDiff = true
		}
	}

	return true
}

func checkOneInsertion(first, second string) bool {
	foundDiff := false

	shorter := first
	longer := second
	if len(second) < len(first) {
		shorter = second
		longer = first
	}

	idxLonger := 0

	for _, char := range shorter {
		if char == int32(longer[idxLonger]) {
			idxLonger++
			continue
		}

		if foundDiff { // can't be off in 2 positions
			return false
		}

		// treat longer[idxLonger] as the inserted character
		foundDiff = true
		// increment the counter and redo the check against the current char in shorter string
		idxLonger++
		if char != int32(longer[idxLonger]) {
			// 2 off in a row
			return false
		}

		idxLonger++
	}

	return true
}

func checkOneAwayRecursive(s, t string) bool {

	if math.Abs(float64(len(s)-len(t))) > 1 {
		return false
	}

	// check base case
	if len(s) <= 1 && len(t) <= 1 {
		// either they're both empty
		// or their 1 difference is manifesting at the end
		return true
	}

	// if reach here then both s and t are non-empty,
	// so we can examine their first chars

	if s[0] == t[0] {
		return checkOneAwayRecursive(s[1:], t[1:])
	}

	// if reach there then leading chars of s and t differ

	if len(s) == len(t)+1 {
		// insertion into s / removal from t case
		return s[1:] == t
	}

	if len(s)+1 == len(t) {
		// insertion into t / removal from s case
		return s == t[1:]
	}

	// if reach here then len(s) equals len(t)

	// replacement case
	// the rest of the strings must be an exact match
	return s[1:] == t[1:]

}
