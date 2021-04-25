package question

func MinSumOfLengths(arr []int, target int) int {
	// 两个思路
	// 一个是通过滑动窗口找出所有和为target的子数组，然后选择不重叠的两个
	// startWith[i]记录以i开始的子数组为target的长度，若不能则为len(arr)+1
	startWith := make([]int, len(arr))
	start, end := 0, 1
	sum := arr[0]
	for start < len(arr) {
		if sum >= target {
			if sum == target {
				startWith[start] = end - start
			} else {
				startWith[start] = len(arr) + 1
			}
			sum -= arr[start]
			start++
		} else if sum < target {
			if end < len(arr) {
				sum += arr[end]
				end++
			} else {
				for i := start; i < len(arr); i++ {
					startWith[i] = len(arr) + 1
				}
				break
			}
		}
	}

	res := len(arr) + 1
	// 这么做没有完全解决重叠的问题
	//minLength := len(arr) + 1
	//for i := 0; i < len(arr); i++ {
	//	if minLength != len(arr)+1 && startWith[i] != len(arr)+1 {
	//		res = min(res, minLength+startWith[i])
	//	}
	//	minLength = min(minLength, startWith[i])
	//}
	endWith := make([]int, len(arr))
	endWith[len(arr)-1] = startWith[len(arr)-1]
	for i := len(arr) - 2; i > 0; i-- {
		endWith[i] = min(endWith[i+1], startWith[i])
	}
	for i := 0; i < len(arr); i++ {
		if i+startWith[i] < len(arr) {
			res = min(res, startWith[i]+endWith[i+startWith[i]])
		}
	}
	if res >= len(arr)+1 {
		return -1
	} else {
		return res
	}
	// 一个是通过前缀和，找到增量为target的长度然后往后更新
	// 本质上和上面的思想没有什么区别

}
