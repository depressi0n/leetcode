package question

func longestPalindrome(s string) string {
	return longestPalindromeCore2(s)
}

// 暴力，长度减小，依次判断
func longestPalindromeCore1(s string) string {
	justify := func(s string, start int, length int) bool {
		left, right := start, start+length-1
		for left <= right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	for length := len(s); length >= 1; length-- {
		for start := 0; start+length-1 < len(s); start++ {
			if justify(s, start, length) {
				return s[start : start+length]
			}
		}
	}
	return ""
}

// 动态规划的思路：
// dp[i][len] 表示s[i:i+len]是否是一个回文串
// dp[i][0] dp[i][1] 都是true
// dp[i][len+1] = dp[i+1][len-2] && s[i] == s[i+len-1]
func longestPalindromeCore2(s string) string {
	if len(s)<=1{
		return s
	}
	res:=s[0:1]
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s)+1)
		// 初始化
		dp[i][0] = true
		dp[i][1] = true
	}
	for length:=2;length<=len(s);length++{
		for i:=0;i+length<=len(s);i++{
			dp[i][length]=dp[i+1][length-2] && s[i] == s[i+length-1]
			if dp[i][length] {
				res=s[i:i+length]
			}
		}
	}
	return res
}

// 中心扩展法：以一个字母为中心，往两边不断扩充，得到的最长结果
// Manacher算法：对中心扩展法的优化，在往两边不断扩充的基础上，借助上一次扩充的结果
func longestPalindromeCore3(s string) string {
	min:=func(x,y int)int{
		if x<y{
			return x
		}
		return y
	}
	expand:=func(left,right int)int{
		for ;left>=0 && right<len(s) && s[left]==s[right];left,right=left-1,right+1{
			// pass
		}
		return (right-left-2)/2
	}
	start,end:=0,-1
	// 借助特殊符号，抹平奇偶长度的差异
	// 可以借助strings.Builder代替，提升性能
	t:="#"
	for i:=0;i<len(s);i++{
		t+=string(s[i])+"#"
	}
	s=t
	// 定义臂长集合
	armLen:=make([]int,0,2*len(s)+1)
	// 记录最右边边界和对应的中心位置
	// 初始为空
	right,j:=-1,-1
	// 以i为中心，开始扩展
	for i:=0;i<len(s);i++{
		var curArmLen int
		// 在之前的扩展范围内，可以借助之间的结果
		if right>=i {
			// 找到j关于i的对称位置
			ii:=j*2-i
			// ii曾经的扩展结果，受限与right-i的长度，因为right之后的部分还未曾触及到
			minArmLen:=min(armLen[ii],right-i)
			// 用之前的结果继续扩展
			curArmLen=expand(i-minArmLen,i+minArmLen)
		}else{
			// 未曾触及到的区域只能是硬扩展
			curArmLen=expand(i,i)
		}
		// 将当前的臂长记录
		armLen=append(armLen,curArmLen)
		// 检查当前是否扩展了右边边界，是则更新
		if i+curArmLen>right{
			j=i
			right=i+curArmLen
		}
		// 检查当前得到的长度是否比已知的长，是则更新
		if curArmLen*2+1>end-start{
			start=i-curArmLen
			end=i+curArmLen
		}
	}
	// 去掉特殊符号
	res:=""
	for i:=start;i<=end;i++{
		if s[i]!='#'{
			res+=string(s[i])
		}
	}
	return res
}