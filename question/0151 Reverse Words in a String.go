package question

func reverseWords(s string) string {
	//	去掉前后空格
	start := 0
	for s[start] == ' ' {
		start++
	}
	end := len(s) - 1
	for s[end] == ' ' {
		end--
	}
	words := make([]string, 0)
	for i := start; i < end+1; i++ {
		if s[i] == ' ' {
			//单词结束
			words = append(words, s[start:i])

			//去掉连续空格
			for s[i] == ' ' {
				i++
			}

			//重置
			start = i
		}
	}
	words = append(words, s[start:end+1])
	res := ""
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i]
		res += " "
	}
	return res[:len(res)-1]
}
