package zmachine

// Stack represents the program stack of a Z-Machine
type Stack struct {
	words    []uint16
	pointer  uint
	capacity uint
}

// Push a value onto a stack.
func (s *Stack) Push(value uint16) {
	s.words[s.pointer] = value
	s.pointer++
}

// Pop a value off of a stack.
func (s *Stack) Pop() uint16 {
	s.pointer--
	return s.words[s.pointer]
}

// Size returns the current size of a stack.
func (s *Stack) Size() uint {
	return s.pointer
}

// NewStack creates a new instance of a Stack.
func NewStack(capacity uint) Stack {
	return Stack{
		words:    make([]uint16, capacity),
		pointer:  0,
		capacity: capacity,
	}
}
