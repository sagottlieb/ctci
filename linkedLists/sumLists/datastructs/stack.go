package datastructs

type Stack struct {
	Top *Node
}

func (s *Stack) Push(v int) {
	newTop := Node{
		Data: v,
		Next: s.Top,
	}
	s.Top = &newTop
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("'pop' called on empty stack")
	}
	formerTop := s.Top
	s.Top = s.Top.Next
	return formerTop.Data
}
