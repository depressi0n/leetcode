package question

func singleNumber137(nums []int) int {
	var so, st int
	for i := 0; i < len(nums); i++ {
		so = ^st & (so ^ nums[i])
		st = ^so & (st ^ nums[i])
	}
	return so
}
