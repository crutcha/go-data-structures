/*
Golang implementation from python datastructures book
http://interactivepython.org/runestone/static/pythonds/BasicDS/TheStackAbstractDataType.html
*/

package main

import (
	"fmt"
)

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

	test1 := "{(([]))}"
	test2 := "{({)}]"

	// hash of reversed characters
	reverse := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	test1IsValid := true
	test1Stack := Stack{length: 0, top: nil}
	for _, value := range test1 {
		if val, ok := reverse[value]; ok {
			if !(test1Stack.top.value == val) {
				test1IsValid = false
				break
			} else {
				test1Stack.Pop()
			}
		} else {
			test1Stack.Push(value)
		}
	}

	test2IsValid := true
	test2Stack := Stack{length: 0, top: nil}
	for _, value := range test2 {
		if val, ok := reverse[value]; ok {
			if !(test2Stack.top.value == val) {
				test2IsValid = false
				break
			} else {
				test2Stack.Pop()
			}
		} else {
			test2Stack.Push(value)
		}
	}

	fmt.Println("---test1---")
	fmt.Println(test1IsValid)
	fmt.Println("---test2---")
	fmt.Println(test2IsValid)
}
