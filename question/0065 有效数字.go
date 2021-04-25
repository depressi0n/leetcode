package question

func isNumber(s string) bool {
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
