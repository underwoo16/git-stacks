package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/underwoo16/git-stacks/queue"
)

func TestIsEmpty(t *testing.T) {
	q := queue.New()

	assert.True(t, q.IsEmpty())

	q.Push(1)

	assert.False(t, q.IsEmpty())

	q.Pop()

	assert.True(t, q.IsEmpty())
}

func TestPush(t *testing.T) {
	q := queue.New()
	q.Push(1)

	assert.Equal(t, 1, q.Size())

	q.Push(2)
	assert.Equal(t, 2, q.Size())

	assert.Equal(t, 1, q.Peek())
}

func New() {
	panic("unimplemented")
}

func TestPop(t *testing.T) {
	q := queue.New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 2, q.Pop())
	assert.Equal(t, 3, q.Pop())
}

func TestSize(t *testing.T) {
	q := queue.New()
	q.Push(1)
	assert.Equal(t, 1, q.Size(), "Queue should have size 1")

	q.Push(2)
	assert.Equal(t, 2, q.Size(), "Queue should have size 2")
	
	q.Pop()
	assert.Equal(t, 1, q.Size(), "Queue should have size 1")

	q.Pop()
	assert.Equal(t, 0, q.Size(), "Queue should have size 0")
}

func TestClear(t *testing.T) {
	q := queue.New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	q.Clear()
	assert.Equal(t, 0, q.Size(), "Queue should be empty after Clear")
}

func TestPeek(t *testing.T) {
	q := queue.New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(t, 1, q.Peek(), "Queue peek should be 1")

	q.Pop()
	assert.Equal(t, 2, q.Peek(), "Queue peek should be 2")

	q.Pop()
	assert.Equal(t, 3, q.Peek(), "Queue peek should be 3")
}
