package question

func Answer0645(nums []int) []int {
	return findErrorNums(nums)
}
func findErrorNums(nums []int) []int {
	// 有一个数字重复，原始数字是1,..,n
	// 找到这个重复数字，修复数字
	// （1）期望序列，真实序列
	// （2）哈希表
	// （1）通过重复1-n,使重复数字出现3次，丢失数字出现1次
	//  计算xor，lowbit= xor & (-xor)表示x和y的二进制表示中最低不同位
	//  可以用lowbit区分x和y
	//  将数字分为两组，a&lowbit==0,b&lowbit!=0
	//  如果a & lowbit = 0 num1=num1 ^ a
	//  否则 num2=num2 ^ a
	//  num1表示第一组数字的异或结果，num2表示第二组数字的异或结果
	//  同一个数字只可能出现在其中一组，此时其他数字出现两次
	//  则 num1和num2与x和y一一对应
	//  再遍历nums数组，如果有元素等于num1，则说明num1是重复数字

	xor := 0
	for i := 0; i < len(nums); i++ {
		xor = xor ^ nums[i]
		xor = xor ^ (i + 1)
	}
	// 此时xor= a ^ b
	lowbit := xor & (-xor) // a和b最低不同位
	num1, num2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&lowbit == 0 {
			num1 ^= nums[i]
		} else {
			num2 ^= nums[i]
		}
		if (i+1)&lowbit == 0 {
			num1 ^= i + 1
		} else {
			num2 ^= i + 1
		}
	}
	flag := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == num1 {
			flag = true
			break
		}
	}
	if flag {
		return []int{num1, num2}
	} else {
		return []int{num2, num1}
	}
	//（2）
	//m:=map[int]int{}
	//for i := 0; i < len(nums); i++ {
	//	m[nums[i]]++
	//}
	//res:=make([]int,2)
	//for i := 0; i <= len(nums); i++ {
	//	if cnt:=m[i];cnt==2{
	//		res[0]=i
	//	}else if cnt==0{
	//		res[1]=i
	//	}
	//}
	//return res
}
