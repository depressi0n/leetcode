package question

import "strconv"

//给你两个二进制字符串，返回它们的和（用二进制表示）。
//输入为 非空 字符串且只包含数字 1 和 0。

func addBinary(a string, b string) string {
	return addBinaryCore2(a, b)
}
func addBinaryCore1(a string, b string) string {

	var c, d []byte
	//看哪一个长
	if len(a) < len(b) {
		c = []byte(b)
		d = []byte(a)
	} else {
		c = []byte(a)
		d = []byte(b)
	}
	b = ""
	for i := 0; i < len(c); i++ {
		if i < len(c)-len(d) {
			b += "0"
		} else {
			b += string(d[i-len(c)+len(d)])
		}
	}
	d = []byte(b)
	carry := 0
	for i := len(c) - 1; i >= 0; i-- {
		switch {
		case c[i] == '1' && d[i] == '1':
			if carry == 1 {
				c[i] = '1'
			} else {
				c[i] = '0'
			}
			carry = 1
		case c[i] == '1' || d[i] == '1':
			if carry == 1 {
				c[i] = '0'
			} else {
				carry = 0
				c[i] = '1'
			}
		default:
			if carry == 1 {
				c[i] = '1'
			} else {
				c[i] = '0'
			}
			carry = 0
		}
	}
	if carry == 1 {
		c = append([]byte{'1'}, c...)
	}
	return string(c)
}

func addBinaryCore2(a string, b string) string {
	ans := ""
	carry := 0
	lengthA, lengthB := len(a), len(b)
	length := max(lengthA, lengthB)
	for i:=0;i<length;i++{
		if i<lengthA{
			carry+=int(a[lengthA-i-1]-'0')
		}
		if i<lengthB{
			carry+=int(b[lengthB-i-1]-'0')
		}
		ans=strconv.Itoa(carry%2)+ans
		carry/=2
	}
	if carry>0{
		ans="1"+ans
	}
	return ans
}
