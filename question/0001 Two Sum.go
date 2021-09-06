package question

// 题目描述：给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

func twoSum0001(nums []int, target int) []int {
	return twoSum2(nums,target)
}

// O(n^2) + O(1)
// 暴力枚举：在当前值后面寻找一个值，使得和为target
func twoSum1(nums []int, target int) []int {
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i]+nums[j]==target{
				return []int{i,j}
			}
		}
	}
	return nil
}
// O(n) + O(n)
// 先将所有的值存入Hash表中，然后通过对Hash表进行遍历，寻找和为target的一对数，要考虑如果target=2*k的情况
func twoSum2(nums []int, target int) []int {
	m:=make(map[int][]int)
	for i:=0;i<len(nums);i++{
		if _,ok:=m[nums[i]];!ok{
			m[nums[i]]=make([]int,0,2)
		}
		m[nums[i]]=append(m[nums[i]],i)
	}
	for k:=range m{
		if _,ok:=m[target-k];ok{
			if target >>1  == k {
				if len(m[k])>=2 {
					return m[k][:2]
				}else{
					continue
				}
			}
			return []int{m[k][0],m[target-k][0]}
		}
	}
	return nil
}
// O(n) + O(n)
// 将当前值之前的值存入Hash表中，寻找之前是否出现过一个值使得与当前值之和等于target
func twoSum3(nums []int, target int) []int {
	m:=make(map[int]int)
	for i:=0;i<len(nums);i++{
		if _,ok:=m[target-nums[i]];ok{
			return []int{i,m[target-nums[i]]}
		}
		m[nums[i]]=i
	}
	return nil
}
