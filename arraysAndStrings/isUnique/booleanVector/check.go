package booleanVector

// implements checker interface
type b struct{}

func New() *b {
	return new(b)
}

func (*b) GetName() string {
	return "BooleanVector"
}

func (*b) Check(s string) bool {
	var bools [128]bool

	for i := 0; i < len(s); i++ {
		offset := int(s[i]) - 1
		if bools[offset] {
			return false
		}
		bools[offset] = true
	}

	return true
}
