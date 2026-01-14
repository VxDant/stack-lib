package stack

// Stack represents a LIFO data structure
type Stack struct {
	elements []interface{}
}

// New creates and returns a new Stack
func New() *Stack {
	return &Stack{elements: make([]interface{}, 0)}
}

// Push adds an element to the top of the stack
func (st *Stack) Push(value interface{}) {
	st.elements = append(st.elements, value)
}

// Pop removes and returns the top element
func (st *Stack) Pop() interface{} {
	if len(st.elements) == 0 {
		return nil
	}
	l := len(st.elements)
	value := st.elements[l-1]
	st.elements = st.elements[:l-1]
	return value
}

// IsEmpty returns true if the stack has no elements
func (st *Stack) IsEmpty() bool {
	return len(st.elements) == 0
}
