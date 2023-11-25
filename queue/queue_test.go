package queue

import "testing"

func TestIsEmpty(t *testing.T) {
	q := New()
	if !q.IsEmpty() {
		t.Error("Queue should be empty")
	}

	q.Push(1)
	if q.IsEmpty() {
		t.Error("Queue should not be empty")
	}

	q.Pop()
	if !q.IsEmpty() {
		t.Error("Queue should be empty")
	}
}

func TestPush(t *testing.T) {
	q := New()
	q.Push(1)
	if q.Size() != 1 {
		t.Error("Queue should have size 1")
	}

	q.Push(2)
	if q.Size() != 2 {
		t.Error("Queue should have size 2")
	}

	if q.Peek() != 1 {
		t.Error("Queue peek should be 1")
	}
}

func TestPop(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Pop() != 1 {
		t.Error("Queue pop should be 1")
	}

	if q.Pop() != 2 {
		t.Error("Queue pop should be 2")
	}

	if q.Pop() != 3 {
		t.Error("Queue pop should be 3")
	}
}

func TestSize(t *testing.T) {
	q := New()
	q.Push(1)
	if q.Size() != 1 {
		t.Error("Queue should have size 1")
	}

	q.Push(2)
	if q.Size() != 2 {
		t.Error("Queue should have size 2")
	}

	q.Pop()
	if q.Size() != 1 {
		t.Error("Queue should have size 1")
	}

	q.Pop()
	if q.Size() != 0 {
		t.Error("Queue should have size 0")
	}
}

func TestClear(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	q.Clear()
	if !q.IsEmpty() {
		t.Error("Queue should be empty")
	}
}

func TestPeek(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.Peek() != 1 {
		t.Error("Queue peek should be 1")
	}

	q.Pop()
	if q.Peek() != 2 {
		t.Error("Queue peek should be 2")
	}

	q.Pop()
	if q.Peek() != 2 {
		t.Error("Queue peek should be 3")
	}
}
