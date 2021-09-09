package question

// 	题目描述：给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。
//	在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
//	找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
//	说明：你不能倾斜容器。


func maxArea(height []int) int {
	return maxAreaCore(height)
}

// 暴力的想法：遍历每一组(left,right)，取最大的那一个
// => 双指针的思想
// 移动数字较小的指针，因为如果移动数字较大的指针，指针之间的距离（减小）和指针指向的较小值（不会变大）会减小
// 为什么正确？
// 对于S(i,j)，假定h[i]<h[j]，转向S(i+1,j)，不计算的状态是S(i,j-1)...S(i,i+1)，这是因为这些状态必然比S(i,j)小
// 因为：短板的高度相同或者更短，长度变小 => 不会大于S(i,j) 即i所在位置不适合作为容器边界了
func maxAreaCore(height []int)int{
	res:=0
	left,right:=0,len(height)-1
	for left<right{
		t:=min(height[left],height[right])*(right-left)
		res=max(res,t)
		if height[left]<=height[right]{
			left++
		}else{
			right--
		}
	}
	return res
}
