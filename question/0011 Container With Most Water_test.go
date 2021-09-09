package question

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"Test1",
			args: args{height: []int{1,8,6,2,5,4,8,3,7}},
		want: 49,
		},
		{
			name:"Test2",
			args: args{height: []int{1,1}},
			want: 1,
		},
		{
			name:"Test3",
			args: args{height: []int{4,3,2,1,4}},
			want: 16,
		},
		{
			name:"Test4",
			args: args{height: []int{1,8,100,2,100,4,8,3,7}},
			want: 200,
		},
		{
			name:"Test5",
			args: args{height: []int{4,4,2,11,0,11,5,11,13,8}},
			want: 55,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
