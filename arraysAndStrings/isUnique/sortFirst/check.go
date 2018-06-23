package sortFirst

import (
	"sort"
)

// implements checker interface
type s struct{}

func New() *s {
	return new(s)
}

func (*s) GetName() string {
	return "Sorted"
}

func stringToByteSlice(s string) []byte {
	var b []byte
	for i := 0; i < len(s); i++ {
		b = append(b, s[i])
	}

	return b
}

func sortByteSlice(in []byte) {
	sort.Slice(in, func(i, j int) bool { return in[i] < in[j] })
}

func (*s) Check(s string) bool {
	if len(s) == 0 {
		return true
	}

	asSlice := stringToByteSlice(s)

	sortByteSlice(asSlice)

	previous := asSlice[0]

	for i := 1; i < len(asSlice); i++ {
		if previous == asSlice[i] {
			return false
		}

		previous = asSlice[i]
	}

	return true
}
