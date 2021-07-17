package question

func IsPowerOfFour(n int) bool {
	// 1必须摆在2*i+1的位置上
	// 而且只能有一个
	//(1) 二分法
	//start,end:=0,15
	//for start<end{
	//	mid:=(end-start)/2+start
	//	if 1 << (2*mid) <n{
	//		start=mid+1
	//	}else if 1 <<1 << (2*mid+1) == n {
	//		return true
	//	}else{
	//		end=mid
	//	}
	//}
	//if 1 <<(2*start) == n {
	//	return true
	//}
	//return false

	//(2) 2的幂 AND 偶数位置(1010 1010 1010 1010 1010 1010 1010 1010 = 0xAAAA AAAA)
	return n > 0 && (n&(n-1)) == 0 && n&0xAAAA_AAAA == 0
}
