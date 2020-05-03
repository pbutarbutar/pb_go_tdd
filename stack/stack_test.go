package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Requirements:
// * A stack is empty on construction
// * A stack has size 0 on construction
// * After n pushes to an empty stack (n > 0), the stack is non-empty and its size equals n
// * If one pushes x then pops, the value popped is x. and the size is one less than old size
// * If one pushes x then peeks, the value returned is x, but the size stays the same
// * If the size is n, then after n pops, the stack is empty and has size 0
// * Popping from an empty stack return an error: ErrNoSuchElement
// * Peeking into an empty stack return an error: ErrNoSuchElement

func TestNewSet(t *testing.T) {
	t.Run("A stack is empty on construction", func(t *testing.T) {
		s := stack.New()
		assert.True(t, s.IsEmpty())
	})
}
