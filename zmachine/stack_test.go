package zmachine

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStackCapacity(t *testing.T) {
	s := NewStack(10)

	Convey("For a stack created with capacity of 10, stack.capacity should be 10", t, func() {
		So(s.capacity, ShouldEqual, 10)
	})
}

func TestStackSize(t *testing.T) {
	s := NewStack(10)

	Convey("For a newly-created stack, stack.size should be 0", t, func() {
		So(s.Size(), ShouldEqual, 0)
	})
}

func TestStackPush(t *testing.T) {
	s := NewStack(10)

	Convey("After pushing value onto a newly-created stack, stack.size should be 1", t, func() {
		s.Push(42)
		So(s.Size(), ShouldEqual, 1)
	})
}

func TestStackPop(t *testing.T) {
	s := NewStack(10)

	Convey("Popping a value off a stack with a size of 1, should return the value and leave stack.size equal to 0", t, func() {
		s.Push(42)
		v, _ := s.Pop()
		So(v, ShouldEqual, 42)
		So(s.Size(), ShouldEqual, 0)
	})
}

func TestStackPeek(t *testing.T) {
	s := NewStack(10)

	Convey("Peeking at a value on stack with a size of 1, should return the value and leave stack.size equal to 1", t, func() {
		s.Push(42)
		v, _ := s.Peek()
		So(v, ShouldEqual, 42)
		So(s.Size(), ShouldEqual, 1)
	})
}

func TestStackPopUnderflow(t *testing.T) {
	s := NewStack(10)

	Convey("Popping a value off an empty stack returns an error", t, func() {
		_, e := s.Pop()
		So(e, ShouldResemble, errors.New("Stack underflow"))
	})
}

func TestStackPeekUnderflow(t *testing.T) {
	s := NewStack(10)

	Convey("Peeking at a value on an empty stack returns an error", t, func() {
		_, e := s.Peek()
		So(e, ShouldResemble, errors.New("Stack underflow"))
	})
}

func TestStackPushOverflow(t *testing.T) {
	s := NewStack(0)

	Convey("Pushing a value onto an empty stack with a capacity of 0 returns an error", t, func() {
		e := s.Push(42)
		So(e, ShouldResemble, errors.New("Stack overflow"))
	})
}
