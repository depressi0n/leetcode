package question

import "math/bits"

// 内置函数 bits.OnesCount()
func hammingDistance1(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

// 移位操作
func hammingDistance2(x int, y int) int {
	res := 0
	for i := 0; i < 31; i++ {
		if ((x>>i)&1)^((y>>i)&1) == 1 {
			res++
		}
	}
	return res
}

// Brian Kernighan 算法 : f(x) = x & (x-1) 则f(x)恰为x删去其二进制表示最右侧的1的结果
// Reference: 《The C Programming Language (Second Edition)》
// Peter Wegner 在1960年的CACM3上出版
func hammingDistance3(x int, y int) int {
	x = x ^ y
	res := 0
	for x != 0 {
		x &= x - 1
		res++
	}
	return res
}
