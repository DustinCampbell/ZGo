package zmachine

import "errors"

// Stack represents the program stack of a Z-Machine
type Stack struct {
	words    []uint16
	pointer  uint
	capacity uint
}

// Push a value onto a stack.
func (s *Stack) Push(value uint16) (error error) {
	if s.pointer == s.capacity {
		return errors.New("Stack overflow")
	}

	s.words[s.pointer] = value
	s.pointer++

	return nil
}

// Pop a value off of a stack.
func (s *Stack) Pop() (result uint16, error error) {
	if s.pointer == 0 {
		return 0, errors.New("Stack underflow")
	}

	s.pointer--
	return s.words[s.pointer], nil
}

// Peek returns the top value on a stack without popping it.
func (s *Stack) Peek() (result uint16, error error) {
	if s.pointer == 0 {
		return 0, errors.New("Stack underflow")
	}

	return s.words[s.pointer-1], nil
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
