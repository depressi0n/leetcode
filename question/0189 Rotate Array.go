package question

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	//解法一，三次逆序
	//convert:= func(nums []int,k int)[]int {
	//	for i:=0;i<len(nums)/2;i++{
	//		nums[i],nums[len(nums)-i-1]=nums[len(nums)-i-1],nums[i]
	//	}
	//	return nums
	//}
	//nums = convert(nums, len(nums))
	//tmp1:=convert(nums[:k],k)
	//tmp2:=convert(nums[k:],len(nums)-k)
	//nums=append(tmp1[:],tmp2...)
	//for i := 0; i < len(nums); i++ {
	//	fmt.Printf("%d,",nums[i])
	//}
	//解法二，直接移动
	//done:=make([]bool,len(nums))
	//cur:=0
	//tmp:=nums[cur]
	//for {
	//	if done[cur] {
	//		cur=0
	//		for ;cur<len(done)&&done[cur];cur++{}
	//		if cur==len(done){
	//			break
	//		}
	//		tmp=nums[cur]
	//	}
	//	nums[(cur+k)%len(nums)],tmp=tmp,nums[(cur+k)%len(nums)]
	//	done[cur]=true
	//	cur=(cur+k)%len(nums)
	//}
	//fmt.Println(nums)
	//return
	//解法三：TODO——对解法二的优化
	//环状替换的例外情况并不仅仅在n%k==0时发生,
	//而是当k和n的最大公约数不为1时出现, 比如n=24, k=14时, 其最大公约数为2, 每次环状替换会替换12个数字, 所以需要遍历两次.
}
