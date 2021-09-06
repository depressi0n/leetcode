package question

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2.0000,
		},
		{
			name: "Test2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3,4},
			},
			want: 2.5000,
		},
		{
			name: "Test3",
			args: args{
				nums1: []int{0,0},
				nums2: []int{0,0},
			},
			want: 0.0000,
		},
		{
			name: "Test4",
			args: args{
				nums1: []int{},
				nums2: []int{1},
			},
			want: 1.0000,
		},
		{
			name: "Test2",
			args: args{
				nums1: []int{2},
				nums2: []int{},
			},
			want: 2.0000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
