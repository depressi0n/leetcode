package question

// 00012的反向操作

func romanToInt(s string) int {
	return romanToIntCore(s)
}

// 硬编码会出现前缀编码的情况，所以需要消除，采用模拟的方式，如果后一个值比前一个值大则符号取反，相当于预读一个值
var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToIntCore(s string) int {
	res:=0
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			res -= value
		} else {
			res += value
		}
	}
	return res
}