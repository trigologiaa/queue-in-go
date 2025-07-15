# Queue - Generic Queue Implementation in Go

A fully featured generic **Queue** implementation in Go using slices.

This queue follows the **FIFO** (First In, First Out) principle and supports core operations along with utility methods for flexibility and convenience.

---

## Table of Contents

- [Queue - Generic Queue Implementation in Go](#queue---generic-queue-implementation-in-go)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Usage](#usage)
  - [Running Tests](#running-tests)
  - [Design Notes](#design-notes)
  - [Example](#example)
  - [Author](#author)
  - [License](#license)
  - [Contact](#contact)

---

## Features

- **Generic**: works with any comparable type (`Queue[T comparable]` in Go 1.18+)

- Core queue operations:

  - `Enqueue(T)` — add an element to the end of the queue
  - `Dequeue() (T, error)` — remove and return the front element (returns an error if empty)
  - `Front() (T, error)` — return (without removing) the front element (error if empty)
  - `PeekLast() (T, error)` — return (without removing) the last element (error if empty)
  - `IsEmpty() bool` — check if the queue is empty
  - `Size() int` — number of elements in the queue

- Utility methods:

  - `Clear()` — empties the queue
  - `Contains(T) bool` — checks if the queue contains a given element
  - `Copy() *Queue[T]` — creates a deep copy of the queue
  - `Reverse()` — reverses the order of elements in-place
  - `ToSlice() []T` — returns a slice copy of the queue elements
  - `Remove(T) bool` — removes the first occurrence of a value from the queue
  - `String() string` — returns a human-readable string representation

- Properly handles empty queue operations by returning Go idiomatic errors (no panics).

- Fully documented with GoDoc/GoComment style comments for `pkg.go.dev`.

---

## Usage

```go
package main

import (
	"fmt"
	"your/module/path/queue" // replace with your actual import path
)

func main() {
	q := queue.NewQueue[int]()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println(q) // Output: Queue: [10 20 30]

	front, err := q.Front()
	if err == nil {
		fmt.Println("Front:", front) // Output: Front: 10
	}

	value, err := q.Dequeue()
	if err == nil {
		fmt.Println("Dequeued:", value) // Output: Dequeued: 10
	}

	fmt.Println("Is empty?", q.IsEmpty()) // Output: false

	q.Clear()
	fmt.Println("Is empty after clear?", q.IsEmpty()) // Output: true
}
```

---

## Running Tests

The implementation comes with comprehensive unit tests using Go’s `testing` package.

To run all tests:

```bash
go test ./queue -v
```

This will execute tests for:

- Normal queue operations (enqueue, dequeue, front, peek last).
- Edge cases (empty queue operations).
- Utility methods (`Copy`, `Reverse`, `Contains`, `ToSlice`, `Remove`).

You can also check test coverage:

```bash
go test ./queue -cover
```

---

## Design Notes

- **Internals**: implemented with a Go slice (`[]T`) for dynamic resizing.
- **Generics**: uses Go 1.18+ type parameters (`T comparable`) for flexibility with any comparable type.
- **Error Handling**: all operations that can fail return idiomatic `error` values (no panics).
- **Copy()** creates a deep copy preserving the element order.
- **Reverse()** modifies the queue in-place reversing the element order.
- **Remove()** deletes the first occurrence of a value.
- **String()** implements `fmt.Stringer` for pretty printing.

---

## Example

```go
q := queue.NewQueue[string]()
q.Enqueue("Alice")
q.Enqueue("Bob")
q.Enqueue("Charlie")

q.Reverse()

for _, name := range q.ToSlice() {
	fmt.Println(name)
}
// Output:
// Charlie
// Bob
// Alice
```

---

## Author

trigologiaa

---

## License

This project is released under the MIT License. Feel free to use, modify, and distribute.

---

## Contact

For questions or contributions, open an issue or contact the author.
