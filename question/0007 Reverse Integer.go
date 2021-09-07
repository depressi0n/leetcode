package question

import "math"

//题目描述：给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//如果反转后整数超过 32 位的有符号整数的范围[−2^31,2^31− 1] ，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）

func reverse(x int) int {
	return reverseCore(x)
}

// math.MaxInt32 = 2147483647
// res*10 + digit <= math.MaxInt32 = [math.MaxInt32/10]*10 +  [math.MaxInt32%10]
// (res-[math.MaxInt32/10]*10)*10 < 7 - digit
// =>
// res > [math.MaxInt32/10]*10 时，因为digit大于0，所以不可能成立【非法】
// res = [math.MaxInt32/10]*10 时，digit必须小于7，此时digit必然小于等于2，所以肯定成立
// res < [math.MaxInt32/10]*10 时，digit<=9，成立
func reverseCore(x int) int {
	res:=0
	for x!=0{
		if res < math.MinInt32 / 10 || res > math.MaxInt32 / 10 {
			return 0
		}
		res=res*10+x%10
		x=x/10
	}
	return res
}
