package question

//  有效数字（按顺序）可以分成以下几个部分：
//一个 小数 或者 整数
//（可选）一个 'e' 或 'E' ，后面跟着一个 整数
//小数（按顺序）可以分成以下几个部分：
//（可选）一个符号字符（'+' 或 '-'）
//下述格式之一：
//至少一位数字，后面跟着一个点 '.'
//至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
//一个点 '.' ，后面跟着至少一位数字
//整数（按顺序）可以分成以下几个部分：
//（可选）一个符号字符（'+' 或 '-'）
//至少一位数字

func isNumber(s string) bool {
	return isNumberCore2(s)
}
func isNumberCore1(s string) bool {
	//去掉两边的字母
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	s = s[i:]
	if len(s) == 0 {
		return false
	}
	i = len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	s = s[:i+1]
	if len(s) == 0 {
		return false
	}
	//检查第一个是否是正负号,并去掉
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	if s[0] < '0' || s[0] > '9' { //去掉正负号以后必须是数字
		return false
	}
	i = 0
	for i < len(s) && (s[i] <= '9' && s[i] >= '0') {
		i++
	}
	if i == len(s) { //纯数字
		return true
	}
	//接下来应该遇到.或者e
	if !(s[i] == '.' || s[i] == 'e') {
		return false
	}
	//如果是. 可能以后会碰到e,可以往下走
	if s[i] == '.' {
		j := i + 1
		for j < len(s) && (s[j] <= '9' && s[j] >= '0') { //跳过数字，找e
			j++
		}
		if j == len(s) { //全部是数字，认为没问题
			return true
		}
		if j == i+1 { //点以后没有数字，认为是错误的
			return false
		}
		if s[j] != 'e' { //如果.之后出现了不是e的其他非数字字符，认为是错误的
			return false
		}
		s = s[j+1:] //检查e后面的
	} else {
		s = s[i+1:]
	}
	//如果是e，之后可能有符号
	j := 0
	for j < len(s) && (s[j] <= '9' && s[j] >= '0') {
		j++
	}
	if j == len(s) {
		return true
	}
	if !(s[j] == '+' || s[j] == '-') {
		return false
	}
	s = s[j+1:]
	//之后必须为数字
	j = 0
	for j < len(s) && (s[j] <= '9' && s[j] >= '0') {
		j++
	}
	if j == len(s) {
		return true
	} else {
		return false
	}
}
func isNumberCore2(s string) bool {
	// 首先去掉两边的空格
	start := 0
	for ; start < len(s) && s[start] == ' '; start++ {
		// pass
	}
	end := len(s)
	for ; end > start && s[end-1] == ' '; end-- {

	}
	if start == end{
		return false
	}
	//s = s[start:end]
	//flag := true // 默认正数
	if s[start] == '-'  || s[start] == '+' {
		//flag = false
		start++
	}
	// 寻找e
	cur := start
	for cur < end {
		if s[cur] == 'e' || s[cur] == 'E' {
			break
		}
		cur++
	}
	if cur == start {
		return false
	}
	// 前面部分
	// 寻找小数点
	dot := start
	for dot < cur {
		if s[dot] == '.' {
			break
		}
		dot++
	}
	// 小数点前
	i := start
	for ; i < dot && '0' <= s[i] && s[i] <= '9'; i++ {

	}
	if i < dot {
		return false
	}
	if dot == start && start == cur-1{
		return false
	}
	// 小数点后
	i++
	for ; i < cur && '0' <= s[i] && s[i] <= '9'; i++ {

	}
	if i < cur {
		return false
	}
	// 指数部分
	// 如果有e/E出现，后面必须有一位数字
	// 接收一个符号位置
	cur++
	if cur < end && (s[cur] == '+' || s[cur] == '-') {
		cur++
	}
	if cur == end{
		return false
	}
	for ; cur < end && '0' <= s[cur] && s[cur] <= '9'; cur++ {

	}
	if cur < end {
		return false
	}
	return true
}

// TODO: 有限状态自动机