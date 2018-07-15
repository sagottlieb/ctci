package stackMin

type node struct {
	val  int
	min  int // the min of all node values for all nodes in the stack below (and including) this node
	next *node
}

type stack struct {
	top *node
}

func (s *stack) push(v int) {
	newTop := node{
		val:  v,
		min:  v, // possibly overwritten below if this is not the first node in the stack
		next: s.top,
	}

	if s.top != nil && s.top.min < newTop.min {
		newTop.min = s.top.min
	}

	s.top = &newTop

	return
}

func (s *stack) pop() int {
	if s.top == nil {
		panic("Pop called on empty stack")
	}

	formerTop := s.top
	s.top = formerTop.next

	return formerTop.val
}

func (s *stack) peek() int {
	if s.top == nil {
		panic("Peek called on empty stack")
	}

	return s.top.val
}

func (s *stack) isEmpty() bool {
	if s.top == nil {
		return true
	}

	return false
}

func (s *stack) min() int {
	if s.top == nil {
		panic("Min called on empty stack")
	}

	return s.top.min
}
