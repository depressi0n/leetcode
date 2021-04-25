package question

import "fmt"

func MaxRepOpt1(text string) int {
	cnt := make(map[byte][]int)
	for i := 0; i < len(text); i++ {
		if _, ok := cnt[text[i]]; !ok {
			cnt[text[i]] = []int{i}
		} else {
			cnt[text[i]] = append(cnt[text[i]], i)
		}
	}
	res := 1
	start := -1
	end := -1
	// 对每一字符都尝试交换一次的处理
	for _, loc := range cnt {
		// 窗口记录为[start:end]
		start = 0
		end = 1
		for start < len(loc) {
			// 可以往后扩,后一个值比前一个大1
			for end < len(loc) && loc[end] == loc[start]+end-start {
				end++
			}
			oldEnd := end
			flag := false
			// 不能往后扩了 只允许交换一次
			// 尝试交换 尽量交换前面的
			if start > 0 || end < len(loc) {
				tail := len(loc)
				if start == 0 {
					tail = len(loc) - 1
				}
				// 交换完成继续尝试往后扩
				for end < tail && loc[end] == loc[start]+end-start+1 {
					end++
				}
				flag = true
			}
			// 计算长度
			length := end - start
			if flag {
				length++
			}
			if length > res {
				res = length
			}
			start = oldEnd
			end = oldEnd + 1
		}
	}
	return res
}

type nonEmpty interface {
	A()
	B()
}
type abc struct {
	a int
}

func (a *abc) A() {
	fmt.Println("implement A")
}

func (a *abc) B() {
	fmt.Println("implement B")
}

func Test1() {
	abcTest := &abc{
		a: 1,
	}
	m := map[string]nonEmpty{}
	m["a"] = abcTest
}

// TODO：滑动窗口，窗口条件：要么窗口内只有一种字符，要么有两种字符且其中一种字符只能出现一次
