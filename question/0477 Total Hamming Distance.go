package question

func TotalHammingDistance1(nums []int) int {
	sum := 0
	hammingDistance := func(a int) {
		for a != 0 {
			a = a & (a - 1)
			sum++
		}
	}
	for length := 1; length < len(nums); length++ {
		for i := 0; i+length < len(nums); i++ {
			hammingDistance(nums[i] ^ nums[i+length])
		}
	}
	return sum
}
func totalHammingDistance2(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < 30; i++ {
		c := 0
		for _, val := range nums {
			c += val >> i & 1
		}
		ans += c * (n - c)
	}
	return
}
