# stack-lib (Open-Source Go library for Stack Datastructure)

A simple, thread-safe stack (LIFO) implementation in Go.

## Installation

```bash
go get github.com/VxDant/stack-lib
```


## Usage

```go
package main

import (
    "fmt"
    "github.com/VxDant/stack-lib"
)

func main() {
    s := stack.New()
    
    // Push elements
    s.Push("first")
    s.Push("second")
    s.Push("third")
    
    // Pop element
    value := s.Pop()
    fmt.Println(value) // Output: third
    
    // Print stack
    fmt.Println(s) // Output: Stack[first second]
}

```

## Code Explanation

```
New() *Stack
```
Creates and returns a new empty stack.

```
Push(value interface{})
```
Adds an element to the top of the stack.

````
Pop() interface{}
````

Removes and returns the top element. Returns nil if stack is empty.

```go
String() string
```

Returns a string representation of the stack.

##  Features

Thread-safe operations using sync.RWMutex

Simple and intuitive API

Works with any data type using interface{}

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.
