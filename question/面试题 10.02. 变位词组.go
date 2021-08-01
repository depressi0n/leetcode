package question

var index int

type TreeNodeMST1002 struct {
	items [26]*TreeNodeMST1002
	index int
}

func (root *TreeNodeMST1002) Insert(str string) int {
	t := [26]int{}
	for i := 0; i < len(str); i++ {
		t[str[i]-'a']++
	}
	cur := root
	for i := 0; i < 26; i++ {
		num := t[i]
		for num != 0 {
			if cur.items[i] == nil {
				cur.items[i] = &TreeNodeMST1002{items: [26]*TreeNodeMST1002{}, index: -1}
			}
			cur = cur.items[i]
			num--
		}
	}
	if cur.index == -1 {
		cur.index = index
		index++
	}
	return cur.index
}
func AnswerMST1002(strs []string) [][]string {
	return groupAnagrams(strs)
}
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0, len(strs))
	index = 0
	root := &TreeNodeMST1002{items: [26]*TreeNodeMST1002{}, index: -1}
	for i := 0; i < len(strs); i++ {
		loc := root.Insert(strs[i])
		if loc >= len(res) {
			res = append(res, []string{strs[i]})
			continue
		}
		res[loc] = append(res[loc], strs[i])
	}
	return res
}
