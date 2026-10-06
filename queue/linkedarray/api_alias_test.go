package arrayqueue

import "testing"

func TestLenAlias(t *testing.T) {
	queue := New[int]()
	if queue.Len() != 0 {
		t.Fatalf("empty queue Len should be 0, got %d", queue.Len())
	}
	queue.PushBack(1)
	queue.PushFront(2)
	if queue.Len() != 2 {
		t.Fatalf("Len should report 2, got %d", queue.Len())
	}
}
