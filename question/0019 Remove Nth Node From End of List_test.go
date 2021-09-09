package question

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEndCore(t *testing.T) {
	var makeList func(arr []int) []*ListNode
	makeList = func(arr []int) []*ListNode {
		if len(arr) == 0 {
			return nil
		}
		head := &ListNode{
			Val:  arr[0],
			Next: nil,
		}
		next := makeList(arr[1:])
		if len(next)==0{
			return []*ListNode{head}
		}
		return append([]*ListNode{head},next...)

	}
	type args struct {
		arr  []*ListNode
		head *ListNode
		n    int
	}
	newArgs:=func(arr []int,n int)args{
		l:=makeList(arr)
		head:=l[0]
		p:=l[0]
		for i:=1;i<len(l);i++{
			p.Next=l[i]
			p=p.Next
		}
		return args{
			arr:  l,
			head: head,
			n:    n,
		}
	}
	arg1:=newArgs([]int{1,2,3,4,5},2)
	arg2:=newArgs([]int{1},1)
	arg3:=newArgs([]int{1,2},1)
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name:"Test1",
			args: arg1,
			want:arg1.head,
		},
		{
			name:"Test2",
			args: arg2,
			want:nil,
		},
		{
			name:"Test3",
			args: arg3,
			want:arg3.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEndCore() = %v, want %v", got, tt.want)
			}
		})
	}
}
