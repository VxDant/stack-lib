package stack

import "sync"

// Stack represents a LIFO data structure
type Stack struct {
	sync     sync.RWMutex
	Elements []interface{}
}

// New creates and returns a new Stack
func New() *Stack {
	return &Stack{Elements: make([]interface{}, 0)}
}

// Push adds an element to the top of the stack
func (st *Stack) Push(value interface{}) {
	st.sync.Lock()
	defer st.sync.Unlock()
	st.Elements = append(st.Elements, value)
}

// Pop removes and returns the top element
func (st *Stack) Pop() interface{} {
	st.sync.Lock()
	defer st.sync.Unlock()

	if len(st.Elements) == 0 {
		return nil
	}
	l := len(st.Elements)
	value := st.Elements[l-1]
	st.Elements = st.Elements[:l-1]
	return value
}

// Peek return the last element from the stack
func (st *Stack) Peek() interface{} {
	st.sync.RLock()
	defer st.sync.RUnlock()
	if len(st.Elements) == 0 {
		return nil
	}
	l := len(st.Elements)
	value := st.Elements[l-1]
	return value
}

// IsEmpty returns true if the stack has no Elements
func (st *Stack) IsEmpty() bool {
	return len(st.Elements) == 0
}
