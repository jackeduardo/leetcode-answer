package main

type MyQueue struct {
	a []int
	b []int
}

//双栈实现
/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.a = append(this.a, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.Peek()
	e := this.b[len(this.b)-1]
	this.b = this.b[:len(this.b)-1]
	return e
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.b) == 0 {
		for len(this.a) > 0 {
			this.b = append(this.b, this.a[len(this.a)-1])
			this.a = this.a[:len(this.a)-1]
		}
	}
	return this.b[len(this.b)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.a) == 0 && len(this.b) == 0 {
		return true
	}
	return false
}
