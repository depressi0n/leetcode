package question

import "strings"

func RepeatedStringMatch(a string, b string) int {
	// 目的 让b成为a的子串
	// 通过重复a达成这个目的， 可以重复0词
	// 分类讨论：
	// （1）b是""，则返回0
	// （2）b本身就已经是a子串，则返回1
	// （3）用两个指针，从b的第一个字母开始匹配，如果能匹配则双指针往后移动，当a的指针到了末尾，则从前面开始，计数器增，继续下去
	// （4）讨论失败的情况，当遇到一个匹配的，说明从这里是不行的，然后找到下一个匹配的位置
	// （5）可以利用类似于KMP的next数组进行优化？
	if b == "" {
		return 0
	}
	p, q := 0, 0
	for i := 0; i < len(a); i++ {
		// 匹配b的第一个字母
		if a[i] == b[0] {
			p = i
			q = 0
			for q < len(b) && a[p%len(a)] == b[q] {
				p++
				q++
			}
			if q == len(b) {
				if p%len(a) == 0 {
					return p / len(a)
				}
				return p/len(a) + 1
			}
		}
	}
	return -1
}

// 很慢...优化方式：可以使用一个类似于hash的想法直接一步跳过中间很多的重复值
func RepeatedStringMatch2(a string, b string) int {
	// 目的 让b成为a的子串
	// 通过重复a达成这个目的， 可以重复0词
	// 分类讨论：
	// （1）b是""，则返回0
	// （2）b本身就已经是a子串，则返回1
	// （3）用两个指针，从b的第一个字母开始匹配，如果能匹配则双指针往后移动，当a的指针到了末尾，则从前面开始，计数器增，继续下去
	// （4）讨论失败的情况，当遇到一个匹配的，说明从这里是不行的，然后找到下一个匹配的位置
	// （5）可以利用类似于KMP的next数组进行优化？
	if b == "" {
		return 0
	}
	help := make(map[string]bool)
	help[a] = true
	q := 0
	for i := 0; i < len(a); i++ {
		// 匹配b的第一个字母
		if a[i] == b[0] {
			q = 0
			// 一步走到末尾?应该是看看b后面的长度够不够
			if len(b)+i < len(a) {
				help[a[i:len(b)+i]] = true
				if _, ok := help[b]; ok {
					return 1
				} else {
					continue
				}
			}
			// 匹配后半段
			help[a[i:]] = true
			if _, ok := help[b[:len(a)-i]]; !ok {
				continue
			}
			q = len(a) - i
			// 开始计数
			cnt := 1
			flag := true
			// 一次跳过len(a)长度，看看能跳几次
			for flag && q+len(a) <= len(b) {
				if _, ok := help[b[q:q+len(a)]]; ok {
					q = q + len(a)
					cnt++
				} else {
					flag = false
					break
				}
			}
			if !flag {
				continue
			}
			// 最后长度不够跳了,则检查最后的长度数量是否相同
			if q < len(b) {
				help[a[:len(b)-q]] = true
				if _, ok := help[b[q:]]; ok {
					return cnt + 1
				}
			} else if q == len(b) {
				return cnt
			}
		}
	}
	return -1
}

// TODO：看懂用KMP去做的
//int repeatedStringMatch3(string a, string b) { // java
//        int n = b.size();
//        int m = a.size();
//        b = " " + b;
//        // next数组
//        for (int i = 2, j = 0; i <= n ; i ++ ) {
//            while (j && b[i] != b[j + 1]) j = ne[j];
//            if (b[i] == b[j + 1]) j ++œ ;
//            ne[i] = j;
//        }
//        int i = 0, j = 0;
//        // 终止条件
//        while(i - j < m) {
//            while (j && b[j + 1] != a[i % m]) j = ne[j];
//            if (a[i % m] == b[j + 1]) j ++ ;
//            i ++ ;
//            if (j == n) break;
//        }
//        if (j == n) {
//            return (i % m == 0) ? i / m : i / m + 1;
//        }
//        else return - 1;
//    }
//
// 反复匹配的方式
func repeatedStringMatch4(a string, b string) int {
	lenA, lenB := len(a), len(b)
	thres := (lenA << 1) + lenB
	t, count := "", 0
	for len(t) < thres {
		if strings.Contains(t, b) {
			return count
		}
		t += a
		count++
	}
	return -1
}
