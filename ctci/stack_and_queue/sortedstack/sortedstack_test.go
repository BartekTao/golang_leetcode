package sortedstack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	data := Constructor()
	data.Push(1)
	data.Push(2)
	fmt.Println(data.Peek())
	data.Pop()
	fmt.Println(data.Peek())

	tests := []struct {
		name string
		want SortedStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
