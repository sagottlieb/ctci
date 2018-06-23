package bitVector

// implements checker interface
type b struct{}

func New() *b {
	return new(b)
}

func (*b) GetName() string {
	return "BitVector"
}

func (*b) Check(s string) bool {
	var leftBits uint64
	var rightBits uint64

	for i := 0; i < len(s); i++ {
		val := uint64(s[i])
		if val < 64 {
			if rightBits&(uint64(1)<<val) > 0 {
				return false
			}
			rightBits |= uint64(1) << val
		} else {
			v2 := uint64(val - 64)
			if leftBits&(uint64(1)<<v2) > 0 {
				return false
			}
			leftBits |= uint64(1) << v2
		}
	}

	return true
}
