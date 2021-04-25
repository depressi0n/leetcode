package question

import (
	"strconv"
	"strings"
)

// 前导零需要去除
// 每个revision至少有一个字母
func compareVersion(version1 string, version2 string) int {
	// 有效性是保证的，所以不用处理
	split1 := strings.Split(version1, ".")
	split2 := strings.Split(version2, ".")
	// 去除前导零
	for i := 0; i < len(split1); i++ {
		for j := 0; j < len(split1[i]) && split1[i][j] == '0'; j++ {
		}
	}
	for i := 0; i < len(split2); i++ {
		for j := 0; j < len(split2[i]) && split2[i][j] == '0'; j++ {
		}
	}
	index := 0
	rev1, rev2 := 0, 0
	for {
		if index < len(split1) && index < len(split2) {
			rev1, _ = strconv.Atoi(split1[index])
			rev2, _ = strconv.Atoi(split2[index])
		} else {
			rev1 = 0
			if index < len(split1) {
				rev1, _ = strconv.Atoi(split1[index])
			}
			rev2 = 0
			if index < len(split2) {
				rev2, _ = strconv.Atoi(split2[index])
			}
		}
		if rev1 < rev2 {
			return -1
		} else if rev1 > rev2 {
			return 1
		} else {
			index++
		}
		if index >= len(split1) && index >= len(split2) {
			return 0
		}
	}
}

// TODO：也可以省点空间，自己计算一下就行，然后每次跳到下一个.
