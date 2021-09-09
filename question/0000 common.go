package question

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func nextPermutationUnique(nums []int) []int {
	var i, j int
	j = len(nums) - 1
	for j > 0 && nums[j] <= nums[j-1] { // 找到第一个递减的
		j--
	}
	i = j
	for i < len(nums)-1 && nums[i+1] == nums[i] { //找到将被交换的位置
		i++
	}
	if j == 0 {
		for n := 0; n <= (len(nums)-1)/2; n++ {
			nums[n], nums[len(nums)-n-1] = nums[len(nums)-n-1], nums[n]
		}
		return nums
	}
	j--
	k := len(nums) - 1
	for k >= i && nums[k] <= nums[j] { // 找到第一个比nums[j]大的
		k--
	}
	nums[j], nums[k] = nums[k], nums[j] //交换
	for n := j + 1; n <= (len(nums)+j)/2; n++ {
		nums[n], nums[len(nums)+j-n] = nums[len(nums)+j-n], nums[n]
	}
	return nums
}

func checkListEqual(l1 *ListNode,l2 *ListNode)bool{
	if l1==nil && l2==nil{
		return true
	}
	if l1==nil || l2==nil{
		return false
	}
	p,q:=l1,l2
	for p!=nil && q!=nil{
		if p.Val != q.Val{
			log.Printf("p=%d,", p.Val)
			log.Printf("q=%d", q.Val)
			log.Printf("\n")
			return false
		}
		p=p.Next
		q=q.Next
	}
	if p!=nil ||q!=nil{
		return false
	}
	return true
}

func makeList(arr []int) *ListNode{
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{
		Val:  arr[0],
		Next: makeList(arr[1:]),
	}
	return head
}