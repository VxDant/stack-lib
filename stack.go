package stack

import (
	"fmt"
	"sync"
)

// Stack represents a LIFO data structure
type Stack struct {
	sync     sync.RWMutex
	elements []interface{}
}

// New creates and returns a new Stack
func New() *Stack {
	return &Stack{elements: make([]interface{}, 0)}
}

// Push adds an element to the top of the stack
func (st *Stack) Push(value interface{}) {
	st.sync.Lock()
	defer st.sync.Unlock()
	st.elements = append(st.elements, value)
}

// Pop removes and returns the top element
func (st *Stack) Pop() interface{} {
	st.sync.Lock()
	defer st.sync.Unlock()

	if len(st.elements) == 0 {
		return nil
	}
	l := len(st.elements)
	value := st.elements[l-1]
	st.elements = st.elements[:l-1]
	return value
}

// Peek return the last element from the stack
func (st *Stack) Peek() interface{} {
	st.sync.RLock()
	defer st.sync.RUnlock()
	if len(st.elements) == 0 {
		return nil
	}
	l := len(st.elements)
	value := st.elements[l-1]
	return value
}

// IsEmpty returns true if the stack has no elements
func (st *Stack) IsEmpty() bool {
	return len(st.elements) == 0
}

// String returns a string representation of the stack
func (st *Stack) String() string {
	st.sync.RLock()
	defer st.sync.RUnlock()
	return fmt.Sprintf("Stack%v", st.elements)
}
