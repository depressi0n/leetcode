package question

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"Test1",
			args: args{
				s:    "ab",
				goal: "ba",
			},
			want:true,
		},
		{
			name:"Test2",
			args: args{
				s:    "ab",
				goal: "ab",
			},
			want:false,
		},
		{
			name:"Test3",
			args: args{
				s:    "aa",
				goal: "aa",
			},
			want:true,
		},
		{
			name:"Test4",
			args: args{
				s:    "aaaaaaabc",
				goal: "aaaaaaacb",
			},
			want:true,
		},
		{
			name:"Test5",
			args: args{
				s:    "abac",
				goal: "abad",
			},
			want:false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
