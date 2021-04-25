package question

import "strconv"

//要考虑负数...
func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 {
		return ""
	}
	if numerator == 0 {
		return "0"
	}
	res := ""
	if numerator < 0 && denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	} else {
		if numerator < 0 {
			numerator = -numerator
			res = "-"
		}
		if denominator < 0 {
			denominator = -denominator
			res = "-"
		}
	}

	if numerator >= denominator {
		t := numerator / denominator
		numerator = numerator - denominator*t
		res += strconv.Itoa(t)
	} else {
		res += "0"
	}
	if numerator == 0 {
		return res
	}
	res += "."
	m := make(map[int]int)
	//以numerator为正数，并且<denominator进入下面的循环
	for numerator < denominator {
		numerator *= 10
		if index, ok := m[numerator]; ok {
			return res[:index] + "(" + res[index:] + ")"
		}
		m[numerator] = len(res)
		t := numerator / denominator
		res += strconv.Itoa(t)
		numerator = numerator - denominator*t
		if numerator == 0 {
			break
		}
	}
	return res
}
