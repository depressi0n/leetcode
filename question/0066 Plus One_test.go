package question

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name:"Test1",
			args: args{digits: []int{1,2,3}},
			want:[]int{1,2,4},
		},
		{
			name:"Test2",
			args: args{digits: []int{1,2,9}},
			want:[]int{1,3,0},
		},
		{
			name:"Test3",
			args: args{digits: []int{9,9,9}},
			want:[]int{1,0,0,0},
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
