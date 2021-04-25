package question

// 求汉明重量，可以用那些很有意思的二进制操作来求
func hammingWeight191_1(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

// TODO 利用 n & (n-1) 将最低位的1置0，而其他位不变的效果
func hammingWeight191_2(num uint32) int {
	sum := 0
	for num != 0 {
		sum++
		num &= num - 1
	}
	return sum
}
