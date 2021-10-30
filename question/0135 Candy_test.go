package question

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{ratings: []int{1, 0, 2}},
			want: 5,
		},
		{
			name: "Test2",
			args: args{ratings: []int{1,2,2}},
			want: 4,
		},
		{
			name: "Test3",
			args: args{ratings: []int{5,7,8,3,4,2,1}},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
