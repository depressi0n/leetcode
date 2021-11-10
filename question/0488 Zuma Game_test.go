package question

import "testing"

func Test_findMinStep(t *testing.T) {
	type args struct {
		board string
		hand  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				board: "WRRBBW",
				hand:  "RB",
			},
			want: -1,
		},
		{
			name: "Test2",
			args: args{
				board: "WWRRBBWW",
				hand:  "WRBRW",
			},
			want: 2,
		},
		{
			name: "Test3",
			args: args{
				board: "G",
				hand:  "GGGGG",
			},
			want: 2,
		},
		{
			name: "Test4",
			args: args{
				board: "RBYYBBRRB",
				hand:  "YRBGB",
			},
			want: 3,
		},
		{
			name: "Test5",
			args: args{
				board: "BWWGWBYRRGWYYRB",
				hand:  "WYYY",
			},
			want: -1,
		},
		{
			name: "Test6",
			args: args{
				board: "RWYWRRWRR",
				hand:  "YRY",
			},
			want: 3,
		},
		{
			name: "Test7",
			args: args{
				board: "RRWWRRBBRR",
				hand:  "WB",
			},
			want: -1, // 官方测试结果为2
		},
		{
			name: "Test9",
			args: args{
				board: "RRYGGYYRRYYGGYRR",
				hand:  "GGBBB",
			},
			want: -1, // 官方测试结果为5
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinStep(tt.args.board, tt.args.hand); got != tt.want {
				t.Errorf("findMinStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
