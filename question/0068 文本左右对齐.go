package question

func fullJustify(words []string, maxWidth int) []string {
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
