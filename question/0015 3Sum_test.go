package question

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name:"Test1",
			args: args{nums: []int{-1,0,1,2,-1,-4}},
			want:[][]int{{-1,-1,2},{-1,0,1}},
		},
		{
			name:"Test2",
			args: args{nums: []int{}},
			want:[][]int{},
		},
		{
			name:"Test3",
			args: args{nums: []int{0}},
			want:[][]int{},
		},
		{
			name:"Test4",
			args: args{nums: []int{0,0,0,0}},
			want:[][]int{{0,0,0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
