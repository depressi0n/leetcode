package question

// 给你两个 没有重复元素 的数组nums1 和nums2，其中nums1是nums2的子集。
//
//请你找出 nums1中每个元素在nums2中的下一个比其大的值。
//
//nums1中数字x的下一个更大元素是指x在nums2中对应位置的右边的第一个比x大的元素。如果不存在，对应位置输出 -1 。
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	return nextGreaterElementCore2(nums1, nums2)
}

func nextGreaterElementCore1(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	cache := make(map[int]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		cache[nums1[i]] = i
	}
	for i := 0; i < len(nums2); i++ {
		if index, ok := cache[nums2[i]]; ok {
			j := i + 1
			for ; j < len(nums2) && nums2[i] >= nums2[j]; j++ {
				// do nothing
			}
			if j == len(nums2) {
				res[index] = -1
			} else {
				res[index] = nums2[j]
			}
		}
	}
	return res
}

// 找到nums1中每个元素在nums2的位置，然后找到右边第一个比它大的值
// 前者可以通过map来完成映射，时间复杂度为O(n)，后者可以通过逐跳的思想，复杂度为O(n)，
// 逐跳的思想：next(a[i])记为在a[i]之后比a[i]大的第一个值的下标
// 我们来更新 a[i-1]
// (1)a[i-1]<a[i]  => next(a[i-1])=i
// (2)a[i-1]==a[i] => next(a[i-1])=next(a[i-1])
// (3)a[i-1]>a[i] 往后跳
//		a[i-1] == a[next(a[i])] next(a[i-1])=next(a[next(a[i])])
//		a[i-1] <  a[next(a[i])] next(a[i-1])=next(a[i])
//		a[i-1] > a[next(a[i])] 继续往后跳
func nextGreaterElementCore2(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	nexts := make([]int, len(nums2))
	nexts[len(nums2)-1] = -1

	for i := len(nums2) - 2; i >= 0; i-- {
		if nums2[i] < nums2[i+1] {
			nexts[i] = i + 1
		} else if nums2[i] == nums2[i+1] {
			nexts[i] = nexts[i+1]
		} else {
			k := nexts[i+1]
			for k != -1 && nums2[i] > nums2[k] {
				k = nexts[k]
			}
			nexts[i] = k
		}
	}
	caches := make(map[int]int, len(nums2))
	for i := len(nums2) - 1; i >= 0; i-- {
		caches[nums2[i]] = i
	}
	for i := 0; i < len(nums1); i++ {
		t := nexts[caches[nums1[i]]]
		if t != -1 {
			res[i] = nums2[t]
		} else {
			res[i] = -1
		}

	}
	return res
}

// 另一种方案
// 找到nums1中每个元素在nums2的位置，然后找到右边第一个比它大的值，因为nums1和nums2中元素各不相同
// 前者可以通过map来完成映射结果，时间复杂度为O(n)，后者可以通过单调栈的思想，复杂度为O(n)
// 处理后者的过程中记录前者，单调栈从后往前处理，维护当前位置右边的更大元素的列表
func nextGreaterElementCore3(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int)
	s := make([]int, 0, len(nums2))
	for i := len(nums2) - 1; i >= 0; i-- {
		num:=nums2[i]
		for len(s)>0 && num >=s[len(s)-1]{
			s=s[:len(s)-1]
		}
		if len(s)>0{
			cache[num]=s[len(s)-1]
		}else{
			cache[num]=-1
		}
		s=append(s,num)
	}
	res:=make([]int,len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i]=cache[nums1[i]]
	}
	return res
}
