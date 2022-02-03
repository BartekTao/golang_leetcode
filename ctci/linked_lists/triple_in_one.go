package ctci

type TripleInOne struct {
	stack      []int
	stackIndex [3]int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{make([]int, stackSize*3), [3]int{-1, -1, -1}}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	// 檢查stack是否滿了
	if this.stackIndex[stackNum] >= len(this.stack)/3-1 {
		return
	} else {
		this.stackIndex[stackNum]++
		this.stack[stackNum*len(this.stack)/3+this.stackIndex[stackNum]] = value
	}
}

func (this *TripleInOne) Pop(stackNum int) int {
	if this.stackIndex[stackNum] != -1 {
		res := this.stack[this.stackIndex[stackNum]+stackNum*len(this.stack)/3]
		this.stackIndex[stackNum]--
		return res
	} else {
		return -1
	}
}

func (this *TripleInOne) Peek(stackNum int) int {
	if this.stackIndex[stackNum] == -1 {
		return -1
	} else {
		return this.stack[this.stackIndex[stackNum]+stackNum*len(this.stack)/3]
	}
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.stackIndex[stackNum] == -1
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
