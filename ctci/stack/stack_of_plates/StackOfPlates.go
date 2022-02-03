package platesstack

type StackOfPlates struct {
	data [][]int
	cap  int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{data: make([][]int, 0), cap: cap}
}

func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 {
		return
	}
	// 目前有幾個row
	n := len(this.data)

	// 若為第一個 => 新增一個row
	if n == 0 || len(this.data[n-1]) == this.cap {
		this.data = append(this.data, []int{val})
	} else {
		this.data[n-1] = append(this.data[n-1], val)
	}
}

func (this *StackOfPlates) Pop() int {
	return this.PopAt(len(this.data) - 1)
}

func (this *StackOfPlates) PopAt(index int) int {
	if this.cap == 0 {
		return -1
	}
	// 目前有幾個 row
	n := len(this.data)

	if index < 0 || index > n-1 {
		return -1
	} else {
		// 當前要 pop 的 array 長度
		ln := len(this.data[index])
		res := this.data[index][ln-1]
		// 若當前要 pop 的 array 長度為1，代表 pop 後需要 delete 該 row
		if len(this.data[index]) == 1 {
			// 若為最後一個row
			if n-1 == index {
				this.data = this.data[:index]
			} else {
				this.data = append(this.data[:index], this.data[index+1:]...)
			}
		} else {
			this.data[index] = this.data[index][:ln-1]
		}
		return res
	}
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
