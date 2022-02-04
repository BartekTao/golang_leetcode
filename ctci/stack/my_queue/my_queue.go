package myQueue

type MyQueue struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{make([]int, 0)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	res := this.data[0]
	this.data = this.data[1:]
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.data[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
