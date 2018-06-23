package bruteForce

// implements checker interface
type b struct{}

func New() *b {
	return new(b)
}

func (*b) GetName() string {
	return "BruteForce"
}

func (*b) Check(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}
