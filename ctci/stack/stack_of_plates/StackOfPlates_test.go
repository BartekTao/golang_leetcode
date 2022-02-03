package platesstack

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	s := Constructor(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.PopAt(0)
	s.PopAt(0)
	s.PopAt(0)

	type args struct {
		cap int
	}
	tests := []struct {
		name string
		args args
		want StackOfPlates
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.cap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
