package question

func Answer0238(nums []int) []int {
	return productExceptSelf(nums)
}
func productExceptSelf(nums []int) []int {
	// O(n) 没有除法
	// 前缀积，后缀积
	//prefixProduct:=make([]int,len(nums)+1)  // [0:i]
	//prefixProduct[0]=1
	//for i := 1; i <=len(nums) ; i++ {
	//	prefixProduct[i]=prefixProduct[i-1]*nums[i-1]
	//}
	//postProduct:=make([]int,len(nums)+1) //[i:len(nums)]
	//postProduct[len(nums)]=1
	//for i := len(nums)-1; i >=0 ; i-- {
	//	postProduct[i]=postProduct[i+1]*nums[i]
	//}
	//res:=make([]int,len(nums))
	//for i := 0; i < len(nums); i++ {
	//	res[i]=prefixProduct[i]*postProduct[i+1]
	//}
	//return res

	// 优化方向：节省空间，将prefixProduct和postProduct的计算融入到res中
	// 首先，res[i] 是prefixProduct[i]
	// 然后，用一个元素遍历跟踪postProduct[i]的值，然后更新res[i]
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * r
		r *= nums[i]
	}
	return res
}
