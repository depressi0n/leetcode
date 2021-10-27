package question

// 题目描述：给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

func isPalindrome0009(x int) bool {
	return isPalindromeCore1(x)
}
// 反转一半的数字，首先负数和尾数为0的可以直接排除
// 怎么判断反转了一半的数字呢，通过反转得到的数值和剩下的值进行比较，如果长度已经超过了，则说明超过一半
// 当是偶数的时候，判断相等，当是奇数的时候，判断revX大10倍
func isPalindromeCore1(x int) bool {
	if x < 0 || (x != 0 && x%10==0){
		return false
	}
	revX:=0
	t:=x
	for t>revX{
		revX=revX*10+t%10
		t=t/10
	}
	if revX == t || t == revX/10 {
		return true
	}
	return false
}

//  全部反转看结果
func isPalindromeCore2(x int) bool {
	if x < 0{
		return false
	}
	revX:=0
	t:=x
	for t!=0{
		revX=revX*10+t%10
		t=t/10
	}
	if revX == x{
		return true
	}
	return false
}