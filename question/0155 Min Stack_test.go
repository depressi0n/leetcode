package question

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		minStack := MinStackConstructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		if got := minStack.GetMin(); got != -3 {
			t.Errorf("MinStack() = %v, want %v", got, -3)
		}
		minStack.Pop()
		if got := minStack.Top(); got != 0 {
			t.Errorf("MinStack() = %v, want %v", got, 0)
		}
		if got := minStack.GetMin(); got != -2 {
			t.Errorf("MinStack() = %v, want %v", got, -2)
		}
	})

}
