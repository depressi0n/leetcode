package question

import (
	"strconv"
	"strings"
)

func Answer1736(time string) string {
	return maximumTime(time)
}

func maximumTime(time string) string {
	// 分类讨论即可
	res := ""
	split := strings.Split(time, ":")
	if split[0][0] == '?' {
		if split[0][1] == '?' {
			res += "23:"
		} else if '0' <= split[0][1] && split[0][1] <= '3' {
			res += "2" + strconv.Itoa(int(split[0][1]-'0')) + ":"
		} else {
			res += "1" + strconv.Itoa(int(split[0][1]-'0')) + ":"
		}
	} else if split[0][0] == '1' || split[0][0] == '0' {
		if split[0][1] == '?' {
			res += strconv.Itoa(int(split[0][0]-'0')) + "9:"
		} else {
			res += split[0] + ":"
		}
	} else {
		if split[0][1] == '?' {
			res += "23:"
		} else {
			res += split[0] + ":"
		}
	}
	if split[1][0] == '?' {
		if split[1][1] == '?' {
			res += "59"
		} else {
			res += "5" + strconv.Itoa(int(split[1][1]-'0'))
		}
	} else {
		if split[1][1] == '?' {
			res += strconv.Itoa(int(split[1][0]-'0')) + "9"
		} else {
			res += split[1]
		}
	}
	return res
}
