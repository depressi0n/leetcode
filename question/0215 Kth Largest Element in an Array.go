package question

func findKthLargest(nums []int, k int) int {
	return findKthLargestCore2(nums, k)
}

// 基于快排
func findKthLargestCore1(nums []int, k int) int {
	if len(nums) == 1 && k == 1 {
		return nums[0]
	}
	// 选择第一个作为pivot
	pivot := nums[0]
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[right] <= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] > pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	if left < k-1 {
		return findKthLargestCore1(nums[left+1:], k-left-1)
	} else if left == k-1 {
		return pivot
	} else {
		return findKthLargestCore1(nums[:left], k)
	}
}

// 基于堆排序，手动实现堆
func findKthLargestCore2(nums []int, k int) int {
	// 建立大根堆，然后做k-1次删除操作，之后堆顶为所求
	var maxHeapify func(a []int, i int, heapSize int)
	maxHeapify = func(a []int, i int, heapSize int) {
		l, r, largest := i*2+1, i*2+2, i
		if l < heapSize && a[l] > a[largest] {
			largest=l
		}
		if r< heapSize && a[r]>a[largest]{
			largest=r
		}
		if largest!=i{
			a[i],a[largest]=a[largest],a[i]
			maxHeapify(a,largest,heapSize)// 向上传递
		}
	}
	buildMaxHeap := func(a []int, heapSize int) {
		// 向下调整
		for i := heapSize / 2; i >= 0; i-- {
			maxHeapify(a, i, heapSize)
		}
	}
	heapSize:=len(nums)
	buildMaxHeap(nums,heapSize)
	// 弹出k个
	for i:=len(nums)-1;i>=len(nums)-k+1;i--{
		nums[0],nums[i]=nums[i],nums[0]
		heapSize--
		maxHeapify(nums,0,heapSize)
	}
	return nums[0]
}
