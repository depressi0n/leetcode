package question

import "strings"

// 给定一个单词数组和一个长度maxWidth，重新排版单词，使其成为每行恰好有maxWidth个字符，且左右两端对齐的文本。
//你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格' '填充，使得每行恰好有maxWidth个字符。
//要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
//文本的最后一行应为左对齐，且单词之间不插入额外的空格。

func fullJustify(words []string, maxWidth int) []string {
	return fullJustifyCore2(words, maxWidth)
}
func fullJustifyCore(words []string, maxWidth int) []string {
	if len(words) == 1 {
		for len(words[0]) < maxWidth {
			words[0] += " "
		}
		return words
	}
	var res []string
	groups := [][]string{{}}
	var lengthCounter, index int
	for i := 0; i < len(words); i++ {
		if lengthCounter+len(words[i])+len(groups[index]) <= maxWidth { //添加之后每个单词之间至少一个空格
			groups[index] = append(groups[index], words[i])
			lengthCounter += len(words[i])
		} else {
			//处理这个组
			res = append(res, "")
			spaceNum := maxWidth - lengthCounter
			positions := make([]int, len(groups[index]))
			ave := 0
			if len(groups[index]) == 1 { //只有一个单词
				positions[0] = spaceNum
			} else {
				ave = spaceNum / (len(groups[index]) - 1)
				more := spaceNum - ave*(len(groups[index])-1)
				for j := 0; j < len(groups[index])-1; j++ {
					if j < more {
						positions[j] = ave + 1
					} else {
						positions[j] = ave
					}
				}
				for j := 0; j < len(groups[index])-1; j++ {
					res[index] += groups[index][j]
					for k := 0; k < positions[j]; k++ {
						res[index] += " "
					}
				}
			}
			res[index] += groups[index][len(groups[index])-1] //这一行处理结束
			for len(res[index]) < maxWidth {
				res[index] += " "
			}
			index++
			lengthCounter = 0
			groups = append(groups, []string{})
			i-- //重新处理刚才没处理完的
		}
	}
	//处理最后一个组,左对齐,最后有没有空格取决长度
	res = append(res, "")
	for j := 0; j < len(groups[index]); j++ {
		res[index] += groups[index][j]
		if j < len(groups[index])-1 {
			res[index] += " "
		}
	}
	for len(res[index]) < maxWidth {
		res[index] += " "
	}
	return res
}

func fullJustifyCore2(words []string, maxWidth int) []string {
	res := make([]string, 0, len(words))
	// 尽可能多多选择单词
	start := 0
	curLen := 0
	t := strings.Builder{}
	for start < len(words) {
		end := start
		curLen = 0
		// 每加一个单词至少有一个空格，此次使用的单词数为
		for end < len(words) && curLen+len(words[end])+end-start <= maxWidth {
			curLen += len(words[end])
			end++
		}
		// 添加到res中
		// 计算空格
		totalSpaceNum := maxWidth - curLen
		remainSpaceNum := 0
		if end == start+1 {
			// 只有一个单词
			remainSpaceNum = 0
		} else {
			remainSpaceNum = totalSpaceNum % (end - start - 1)
		}
		t.Reset()
		if end == len(words) {
			// 最后一行要左对齐
			for i := 0; i < end-start; i++ {
				t.WriteString(words[start+i])
				if i<end-start-1{
					t.WriteByte(' ')
				}
			}
			if totalSpaceNum>end-start-1{
				t.WriteString(strings.Repeat(" ", totalSpaceNum-end+start+1))
			}
			res = append(res, t.String())
			break
		}
		for i := 0; i < end-start; i++ {
			t.WriteString(words[start+i])
			// 写入空格
			if i < remainSpaceNum {
				t.WriteByte(' ')
			}
			if i < end-start-1 {
				t.WriteString(strings.Repeat(" ", totalSpaceNum/(end-start-1)))
			}
		}
		if end == start+1 {
			t.WriteString(strings.Repeat(" ", totalSpaceNum))
		}
		res = append(res, t.String())
		// 下一轮
		start = end
	}
	return res
}
