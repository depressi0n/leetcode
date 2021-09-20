package question

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"Test1",
			args: args{s: "2"},
			want:true,
		},
		{
			name:"Test2",
			args: args{s: "0089"},
			want:true,
		},
		{
			name:"Test3",
			args: args{s: "-0.1"},
			want:true,
		},
		{
			name:"Test4",
			args: args{s: "+3.14"},
			want:true,
		},
		{
			name:"Test5",
			args: args{s: "-.9"},
			want:true,
		},
		{
			name:"Test6",
			args: args{s: "2e10"},
			want:true,
		},
		{
			name:"Test7",
			args: args{s: "-90E3"},
			want:true,
		},
		{
			name:"Test8",
			args: args{s: "+6e-1"},
			want:true,
		},
		{
			name:"Test9",
			args: args{s: "-123.456e789"},
			want:true,
		},
		{
			name:"Test10",
			args: args{s: "abc"},
			want:false,
		},
		{
			name:"Test11",
			args: args{s: "1a"},
			want:false,
		},
		{
			name:"Test12",
			args: args{s: "e3"},
			want:false,
		},
		{
			name:"Test13",
			args: args{s: "99e2.5"},
			want:false,
		},
		{
			name:"Test14",
			args: args{s: "--6"},
			want:false,
		},
		{
			name:"Test15",
			args: args{s: "-+3"},
			want:false,
		},
		{
			name:"Test16",
			args: args{s: "95a54e53"},
			want:false,
		},
		{
			name:"Test17",
			args: args{s: "."},
			want:false,
		},
		{
			name:"Test18",
			args: args{s: "0e"},
			want:false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
