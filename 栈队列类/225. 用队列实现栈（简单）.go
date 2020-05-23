package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor_MyStack()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj.Top())
	fmt.Println(obj.Pop())
}

/** Initialize your data structure here. */
type MyStack struct {
	l list.List
	t list.List
}

/** Initialize your data structure here. */
func Constructor_MyStack() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.l.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for this.l.Front().Next() != nil {
		this.t.PushBack(this.l.Remove(this.l.Front()))
	}
	top := this.l.Remove(this.l.Front())
	for this.t.Front() != nil {
		this.l.PushBack(this.t.Remove(this.t.Front()))
	}
	return top.(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for this.l.Front().Next() != nil {
		this.t.PushBack(this.l.Remove(this.l.Front()))
	}
	top := this.l.Remove(this.l.Front())
	for this.t.Front() != nil {
		this.l.PushBack(this.t.Remove(this.t.Front()))
	}
	this.l.PushBack(top)
	return top.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.l.Len() == 0 {
		return true
	}
	return false
}

//切片实现

type MyStack_slice struct {
	slice []int
}

/** Initialize your data structure here. */
func Constructor_MyStack_slice() MyStack_slice {
	return MyStack_slice{}
}

/** Push element x onto stack. */
func (s *MyStack_slice) Push_slice(x int) {
	s.slice = append(s.slice, x)
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack_slice) Pop_slice() int {
	if len(s.slice) == 0 {
		return -1
	}
	r := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return r
}

/** Get the top element. */
func (s *MyStack_slice) Top_slice() int {
	if len(s.slice) == 0 {
		return -1
	}
	return s.slice[len(s.slice)-1]
}

/** Returns whether the stack is empty. */
func (s *MyStack_slice) Empty_slice() bool {
	return len(s.slice) == 0
}
