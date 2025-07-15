package queue

import "testing"

func TestEnqueueAndDequeue(t *testing.T) {
	q := NewQueue[int]()
	if !q.IsEmpty() {
		t.Error("expected queue to be empty initially")
	}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	if q.Size() != 3 {
		t.Errorf("expected size 3, got %d", q.Size())
	}
	if q.IsEmpty() {
		t.Error("expected queue to be non-empty after enqueues")
	}
	value, err := q.Dequeue()
	if err != nil {
		t.Error("unexpected error on Dequeue:", err)
	}
	if value != 10 {
		t.Errorf("expected dequeued value 10, got %d", value)
	}
	value, err = q.Dequeue()
	if err != nil {
		t.Error("unexpected error on Dequeue:", err)
	}
	if value != 20 {
		t.Errorf("expected dequeued value 20, got %d", value)
	}
	value, err = q.Dequeue()
	if err != nil {
		t.Error("unexpected error on Dequeue:", err)
	}
	if value != 30 {
		t.Errorf("expected dequeued value 30, got %d", value)
	}
	_, err = q.Dequeue()
	if err == nil {
		t.Error("expected error when dequeuing from empty queue")
	}
}

func TestFront(t *testing.T) {
	q := NewQueue[string]()
	_, err := q.Front()
	if err == nil {
		t.Error("expected error on Front from empty queue")
	}
	q.Enqueue("foo")
	front, err := q.Front()
	if err != nil {
		t.Error("unexpected error on Front:", err)
	}
	if front != "foo" {
		t.Errorf("expected Front value 'foo', got '%s'", front)
	}
	q.Enqueue("bar")
	front, _ = q.Front()
	if front != "foo" {
		t.Errorf("expected Front value 'foo' after enqueue, got '%s'", front)
	}
}

func TestPeekLast(t *testing.T) {
	q := NewQueue[int]()
	_, err := q.PeekLast()
	if err == nil {
		t.Error("expected error on PeekLast from empty queue")
	}
	q.Enqueue(1)
	q.Enqueue(2)
	last, err := q.PeekLast()
	if err != nil {
		t.Error("unexpected error on PeekLast:", err)
	}
	if last != 2 {
		t.Errorf("expected PeekLast value 2, got %d", last)
	}
}

func TestClear(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Clear()
	if !q.IsEmpty() {
		t.Error("expected queue to be empty after Clear")
	}
	if q.Size() != 0 {
		t.Errorf("expected size 0 after Clear, got %d", q.Size())
	}
	_, err := q.Dequeue()
	if err == nil {
		t.Error("expected error when dequeuing from cleared queue")
	}
}

func TestCopy(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	clone := q.Copy()
	if clone.Size() != q.Size() {
		t.Errorf("expected clone size %d, got %d", q.Size(), clone.Size())
	}
	clone.Enqueue(3)
	if q.Size() == clone.Size() {
		t.Error("expected original and clone to diverge after modifying clone")
	}
}

func TestReverse(t *testing.T) {
	q := NewQueue[int]()
	for i := 1; i <= 3; i++ {
		q.Enqueue(i)
	}
	q.Reverse()
	expected := []int{3, 2, 1}
	actual := q.ToSlice()
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("expected %v at position %d, got %v", v, i, actual[i])
		}
	}
}

func TestToSlice(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(42)
	q.Enqueue(7)
	slice := q.ToSlice()
	if len(slice) != 2 || slice[0] != 42 || slice[1] != 7 {
		t.Errorf("unexpected slice content: %v", slice)
	}
}

func TestContains(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("apple")
	q.Enqueue("banana")
	if !q.Contains("apple") {
		t.Error("expected queue to contain 'apple'")
	}
	if q.Contains("orange") {
		t.Error("did not expect queue to contain 'orange'")
	}
}

func TestRemove(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	removed := q.Remove(2)
	if !removed {
		t.Error("expected Remove to return true for existing element")
	}
	if q.Contains(2) {
		t.Error("expected element 2 to be removed from queue")
	}
	removed = q.Remove(42)
	if removed {
		t.Error("expected Remove to return false for non-existing element")
	}
}

func TestString(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	got := q.String()
	want := "Queue: [1 2]"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
