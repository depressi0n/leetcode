package question

import (
	"reflect"
	"testing"
)

func Test_twoSum1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want:[][]int{{0,1},{1,0}},
		},
		{
			name: "Test2",
			args: args{
				nums:   []int{3,2,4},
				target: 6,
			},
			want:[][]int{{1,2},{2,1}},
		},
		{
			name: "Test3",
			args: args{
				nums:   []int{3,3},
				target: 6,
			},
			want:[][]int{{0,1},{1,0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum0001(tt.args.nums, tt.args.target)
			// 因为map遍历时具有随机的特性，所以需要特别判定
			i:=0
			for ;i<len(tt.want);i++{
				if reflect.DeepEqual(got, tt.want[i]) {
					break
				}
			}
			if i == len(tt.want) {
				t.Errorf("twoSum1() = %v, want %v", got, tt.want)
			}

		})
	}
}
