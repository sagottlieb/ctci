package hashmap

// implements checker interface
type m struct{}

func New() *m {
	return new(m)
}

func (*m) GetName() string {
	return "Map"
}

func (*m) Check(s string) bool {
	seen := map[byte]bool{}

	for i := 0; i < len(s); i++ {
		if _, exists := seen[s[i]]; exists {
			return false
		}

		seen[s[i]] = true
	}

	return true
}
