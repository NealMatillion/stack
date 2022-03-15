package f64stack

import "testing"

func TestLen(t *testing.T) {
	s := Stack{}

	if s.Len() != 0 {
		t.Log("An empty stack should have length 0")
		t.Fail()
	}
}

func TestPushPop(t *testing.T) {
	s := Stack{}

	s.Push(42)

	got, err := s.Pop()
	if err != nil {
		t.Log("unexpected error", err)
		t.Fail()
	}

	if got != 42 {
		t.Logf("s.Push(42) ; s.Pop() = %f ; want %d", got, 42)
		t.Fail()
	}
}

func TestPopEmpty(t *testing.T) {
	s := Stack{}

	_, err := s.Pop()
	if err == nil {
		t.Log("popping an empty stack should return an error")
		t.Fail()
	}
}
