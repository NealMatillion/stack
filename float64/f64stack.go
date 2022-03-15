package f64stack

import (
	"errors"
	"sync"
)

// Stack is a stack of integers
type Stack struct {
	mu   sync.Mutex
	data []float64
}

// Len returns the number of items in the stack.
func (s *Stack) Len() int {
	return len(s.data)
}

// Push the given value on to the stack
func (s *Stack) Push(f float64) {
	s.mu.Lock()
	s.data = append(s.data, f)
	s.mu.Unlock()
}

// Pop the top value off the stack if there is one.  Returns an error if
// the stack is empty.
func (s *Stack) Pop() (float64, error) {
	if len(s.data) == 0 {
		return 0, errors.New("cannot pop an empty stack")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	last := len(s.data) - 1
	v := s.data[last]
	s.data = s.data[:last]
	return v, nil
}
