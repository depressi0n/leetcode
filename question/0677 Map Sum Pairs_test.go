package question

import (
	"testing"
)

func Test0677(t *testing.T) {
	t.Run("Test0677", func(t *testing.T) {
		m := Constructor()
		m.Insert("apple", 3)
		if got := m.Sum("ap"); got != 3 {
			t.Errorf(" m.Sum() = %v, want %v", got, 3)
		}
		m.Insert("app", 2)
		if got := m.Sum("ap"); got != 5 {
			t.Errorf(" m.Sum() = %v, want %v", got, 5)
		}
		m.Insert("apple", 2)
		if got := m.Sum("ap"); got != 4 {
			t.Errorf(" m.Sum() = %v, want %v", got, 4)
		}
	})
}
