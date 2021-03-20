// stack.go implements a stack which accepts generic data
// Go has no generics, however all types in Go implement the empty interface{}
// We can use this orthogonal design and specify a stack that accepts an empty interface type.
// Go will perform type assertion on the actual underlying types.

package stack

import (
	"fmt"
)

type Any interface {
}

type Stack struct {
	count   int
	storage []Any
}

func (s *Stack) length() int {
	return s.count
}

func (s *Stack) push(el Any) {
	s.storage[s.count] = el
	s.count++
}

func (s *Stack) peek() Any {
	if s.count == 0 {
		return nil
	}
	return s.storage[s.count-1]
}

func (s *Stack) pop() Any {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.storage[s.count]
}

func testStack() {
	kstack := Stack{0, make([]Any, 64)}
	fmt.Printf("Stack length: %v\n", kstack.length())
	kstack.push("world!")
	kstack.push("hello")
	fmt.Printf("Stack length: %v\n", kstack.length())
	fmt.Printf("Peek: %v\n", kstack.peek())
	fmt.Println(kstack.pop())
	fmt.Println(kstack.pop())
}
