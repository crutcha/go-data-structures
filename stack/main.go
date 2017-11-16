/*
Golang implementation from python datastructures book
http://interactivepython.org/runestone/static/pythonds/BasicDS/TheStackAbstractDataType.html
*/

package main

import "fmt"

type Stack struct {
	length int
	top    *StackNode
}

type StackNode struct {
	value    interface{}
	previous *StackNode
}

func (stack *Stack) IsEmpty() bool {
	if stack.length == 0 {
		return true
	}
	return false
}

func (stack *Stack) Push(data interface{}) {
	// Set previous to current top
	top := &StackNode{
		value:    data,
		previous: stack.top,
	}

	// incremenet stack length
	stack.length++

	// Stacks new top is this stack node
	stack.top = top
}

func (stack *Stack) Pop() interface{} {
	item := stack.top

	// set top to previous
	stack.top = item.previous

	// decrement stack length
	stack.length--

	return item.value
}

func (stack *Stack) Peek() interface{} {
	return stack.top.value
}

func (stack *Stack) Size() int {
	return stack.length
}

func main() {
	mystack := Stack{length: 0, top: nil}
	fmt.Println(mystack)
	fmt.Println(mystack.IsEmpty())

	mystack.Push(3)
	fmt.Println(mystack)
	fmt.Println(mystack.IsEmpty())

	mystack.Push("a string")
	fmt.Println(mystack)
	fmt.Println(mystack.IsEmpty())

	fmt.Println(mystack.Peek())
	fmt.Println(mystack.Size())

	this := mystack.Pop()
	fmt.Println(mystack)
	fmt.Println(this)

	fmt.Println(mystack.Peek())

	fmt.Println(mystack.Size())
}
