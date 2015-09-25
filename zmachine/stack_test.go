package zmachine

import (
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
		v := s.Pop()
		So(v, ShouldEqual, 42)
		So(s.Size(), ShouldEqual, 0)
	})
}
