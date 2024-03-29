package question

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test1",
			args: args{nums: []int{1, 2, 3}},
			want: []int{1, 3, 2},
		},
		{
			name: "Test2",
			args: args{nums: []int{3, 2, 1}},
			want: []int{1, 2, 3},
		},
		{
			name: "Test3",
			args: args{nums: []int{1,1,5}},
			want: []int{1,5,1},
		},
		{
			name: "Test4",
			args: args{nums: []int{1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPermutation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
