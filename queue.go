// Package queue provides a generic FIFO (First In, First Out)
// implementation for Go.
//
// The Queue type supports standard operations such as Enqueue, Dequeue, and Front,
// along with utility methods like PeekLast, Remove, Contains, Copy, Reverse,
// ToSlice, and Clear.
//
// It is implemented as a wrapper around a Go slice, supporting any comparable type
// T, and offers dynamic resizing as elements are added or removed.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(10)
//	q.Enqueue(20)
//	q.Enqueue(30)
//	fmt.Println(q) // Queue: [10 20 30]
//	front, _ := q.Front()
//	fmt.Println("Front:", front) // Front: 10
//	value, _ := q.Dequeue()
//	fmt.Println("Dequeued:", value) // Dequeued: 10
//	fmt.Println(q.IsEmpty()) // false
//	q.Clear()
//	fmt.Println(q.IsEmpty()) // true
package queue

import (
	"fmt"
	"slices"
)

// A generic FIFO (First In, First Out) data structure.
//
// Queue[T] holds elements of any comparable type T.
//
// Internally, it uses a dynamically growing slice to store elements.
type Queue[T comparable] struct {
	data []T
}

// Creates and returns a new empty Queue for type T.
//
// Returns:
//   - *Queue[T]: A new empty queue for type T.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(42)
//	fmt.Println(q) // Queue: [42]
func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

// Adds a new element to the end of the queue.
//
// Parameters:
//   - data: The element to be added to the queue.
//
// Example:
//
//	q := queue.NewQueue[string]()
//	q.Enqueue("hello")
//	q.Enqueue("world")
func (q *Queue[T]) Enqueue(data T) {
	q.data = append(q.data, data)
}

// Removes and returns the front element of the queue.
//
// Returns:
//   - value: The front element of the queue.
//   - error: An error if the queue is empty.
//
// If the queue is empty, Dequeue returns the zero value of T and an error.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(1)
//	value, err := q.Dequeue()
//	if err == nil {
//	    fmt.Println(value) // 1
//	}
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty queue")
	}
	head := q.data[0]
	q.data = q.data[1:]
	return head, nil
}

// Returns the front element of the queue without removing it.
//
// Returns:
//   - value: The front element of the queue.
//   - error: An error if the queue is empty.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(5)
//	front, err := q.Front()
//	if err == nil {
//	    fmt.Println(front) // 5
//	}
func (q *Queue[T]) Front() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty queue")
	}
	head := q.data[0]
	return head, nil
}

// Reports whether the queue contains no elements.
//
// Returns:
//   - bool: true if the queue is empty; false otherwise.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	fmt.Println(q.IsEmpty()) // true
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// Returns the number of elements currently in the queue.
//
// Returns:
//   - int: The count of elements in the queue.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(1)
//	q.Enqueue(2)
//	fmt.Println(q.Size()) // 2
func (q *Queue[T]) Size() int {
	return len(q.data)
}

// Removes all elements from the queue, resetting it to empty.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(1)
//	q.Clear()
//	fmt.Println(q.IsEmpty()) // true
func (q *Queue[T]) Clear() {
	q.data = make([]T, 0)
}

// Returns a string representation of the queue.
//
// Returns:
//   - string: A string representation of the queue.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(1)
//	q.Enqueue(2)
//	fmt.Println(q.String()) // Queue: [1 2]
func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue: %v", q.data)
}

// Returns the last element of the queue without removing it.
//
// Returns:
//   - value: The last element of the queue.
//   - error: An error if the queue is empty.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(1)
//	q.Enqueue(2)
//	last, err := q.PeekLast()
//	if err == nil {
//	    fmt.Println(last) // 2
//	}
func (q *Queue[T]) PeekLast() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty queue")
	}
	return q.data[q.Size()-1], nil
}

// Reports whether the queue contains the given value.
//
// Parameters:
//   - data: The value to search for.
//
// Returns:
//   - bool: true if the value exists in the queue; false otherwise.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(10)
//	fmt.Println(q.Contains(10)) // true
//	fmt.Println(q.Contains(5))  // false
func (q *Queue[T]) Contains(data T) bool {
	return slices.Contains(q.data, data)
}

// Returns a copy of the queue's elements as a slice.
//
// Returns:
//   - []T: A copy of the queue's internal slice.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(1)
//	q.Enqueue(2)
//	slice := q.ToSlice()
//	fmt.Println(slice) // [1 2]
func (q *Queue[T]) ToSlice() []T {
	result := make([]T, q.Size())
	copy(result, q.data)
	return result
}

// Deletes the first occurrence of a given value from the queue.
//
// Parameters:
//   - data: The value to remove.
//
// Returns:
//   - bool: true if the element was found and removed; false otherwise.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(1)
//	q.Enqueue(2)
//	q.Remove(1)
//	fmt.Println(q) // Queue: [2]
func (q *Queue[T]) Remove(data T) bool {
	for i, v := range q.data {
		if v == data {
			q.data = slices.Delete(q.data, i, i+1)
			return true
		}
	}
	return false
}

// Creates and returns a deep copy of the queue.
//
// Returns:
//   - *Queue[T]: A new queue with the same elements.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(1)
//	clone := q.Copy()
//	fmt.Println(clone) // Queue: [1]
func (q *Queue[T]) Copy() *Queue[T] {
	newData := make([]T, q.Size())
	copy(newData, q.data)
	return &Queue[T]{data: newData}
}

// Reverses the order of elements in the queue.
//
// Example:
//
//	q := queue.NewQueue[int]()
//	q.Enqueue(1)
//	q.Enqueue(2)
//	q.Enqueue(3)
//	q.Reverse()
//	fmt.Println(q) // Queue: [3 2 1]
func (q *Queue[T]) Reverse() {
	for i, j := 0, q.Size()-1; i < j; i, j = i+1, j-1 {
		q.data[i], q.data[j] = q.data[j], q.data[i]
	}
}
