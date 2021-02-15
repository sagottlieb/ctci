package datastructs

type Stack struct {
	Top *Node
}

func (s *Stack) push(v int) {
	newTop := Node{
		Data: v,
		Next: s.Top,
	}
	s.Top = &newTop
}

func (s *Stack) isEmpty() bool {
	return s.Top == nil
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		panic("'pop' called on empty stack")
	}
	formerTop := s.Top
	s.Top = s.Top.Next
	return formerTop.Data
}
