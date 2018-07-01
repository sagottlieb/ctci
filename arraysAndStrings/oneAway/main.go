package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Run go test and go test -bench=.")
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

	if len(s) == len(t) + 1 {
		// insertion into s / removal from t case
		return s[1:] == t
	}

	if len(s) + 1 == len(t) {
		// insertion into t / removal from s case
		return s == t[1:]
	}

	// if reach here then len(s) equals len(t)

	// replacement case
	// the rest of the strings must be an exact match
	return s[1:] == t[1:]

}
