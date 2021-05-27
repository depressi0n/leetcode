package question

func IntToRoman(num int) string {
	//roman2Num:=map[string]int{"I":1,"IV":4,"V":5,"IX":9,"X":10,"XL":40,"L":50,"XC":90,"C":100,"CD":400,"D":500,"CM":900,"M":1000}
	num2Roman := map[int]string{1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M"}
	if res, ok := num2Roman[num]; ok {
		return res
	}
	res := ""
	factor := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	cur := 0
	for num != 0 && cur < len(factor) {
		cnt := num / factor[cur]
		num = num - cnt*factor[cur]
		for cnt > 0 {
			res += num2Roman[factor[cur]]
			cnt--
		}

		cur++
		if cur < len(factor) {
			if num >= factor[cur] {
				num -= factor[cur]
				res += num2Roman[factor[cur]]
			}
			cur++
		}
	}
	return res
}

// TODO:看懂下面的理论，硬编码表示所有的可能性
//M := []string{"", "M", "MM", "MMM"}
//C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
//X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
//I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
//return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
