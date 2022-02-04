package sortedstack

type SortedStack struct {
	sortedData []int
}

func Constructor() SortedStack {
	return SortedStack{make([]int, 0)}
}

func (this *SortedStack) Push(val int) {
	i := len(this.sortedData)
	if i == 0 || this.sortedData[i-1] <= val {
		this.sortedData = append(this.sortedData, val)
		return
	}
	for i != 0 && this.sortedData[i-1] > val {
		i--
	}
	this.sortedData = insert(this.sortedData, val, i)
}

func insert(a []int, c int, i int) []int {
	return append(a[:i], append([]int{c}, a[i:]...)...)
}

func (this *SortedStack) Pop() {
	if this.IsEmpty() {
		return
	}
	this.sortedData = this.sortedData[1:]
}

func (this *SortedStack) Peek() int {
	return this.sortedData[0]
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.sortedData) == 0
}
