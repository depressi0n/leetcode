package question

//TODO：条件判断有问题，明天再想想
func numDecodings(s string) int {
	//if len(s) == 0 || s[0] == '0' { //保证了以1-9的数开头
	//	return 0
	//}
	//// 如果出现0，必须和前一个数绑在一起，如果是个无效值，则整体无效
	//isValidZero := func(index int) bool { //s[index]=='0'
	//	if index-1 >= 0 && (s[index-1] == '1' || s[index-1] == '2') {
	//		return true
	//	}
	//	return false
	//}
	//isValid := func(index int) bool { //将index和index-1绑定在一块
	//	value, err := strconv.Atoi(s[index-1 : index+1])
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	if 0 < value && value < 27 {
	//		return true
	//	}
	//	return false
	//}
	////dp[i]表示从[0:i]的解码方式
	//dp := make([]int, len(s)+1)
	//dp[0] = 1
	////dp[1] = 1
	//for i := 1; i <= len(s); i++ {
	//	if i < len(s) && s[i] == '0' { //后一个数是0
	//		if isValidZero(i) { //可以和前面的数绑定
	//			dp[i] = -1        //标记为被绑定
	//			dp[i+1] = dp[i-1] //绑定以后的效果
	//			i++               //直接跳到下一个
	//			continue
	//		} else { //无法绑定，而0是无效值，整体无效
	//			return 0
	//		}
	//	}
	//	//后面的数不是0，肯定可以单独解码
	//	dp[i] = dp[i-1]
	//	//如果前面一个数是0，那么直接往后走就好了
	//	if i-2>=0 && s[i-2] == '0' {
	//		continue
	//	}
	//	if i-2>=0 && isValid(i - 1) { //有效的绑定
	//		dp[i] += dp[i-2]
	//	}
	//}
	//return dp[len(s)]
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	prev, next := 1, 1
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == '0' && s[i-1] != '1' && s[i-1] != '2': //非法
			return 0
		case s[i] == '0': //上面case后半部分的补集
			next = prev
		case s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6'): //加起来
			prev, next = next, prev+next
		default: //一般情况就是前后相等
			prev = next
		}
	}
	return next
}
