package question

func NumberToWords1(num int) string {
	// 用英文单词表示一个正整数
	// 首先英文单词分为
	// 个 one two three four five six seven eight nine ten
	// 十 ten twenty thirty forty fifty sixty seventy eighty ninety
	// 百 hundred
	// 千 thousand
	// 百万 million
	// 十亿 billion
	// 每三个数字分为一组，内部统一以百、十、个位的形式表示
	oneToNineteen := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	mtens := []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	units := []string{"Thousand", "Million", "Billion"} // 1000
	// 特殊情况
	if num < 20 {
		return oneToNineteen[num]
	}
	res := ""
	//var tmp [3]int
	cnt := -1
	for num > 0 {
		// 取出最后三个数
		tmps := num % 1000
		num = num / 1000
		cnt++
		switch {
		case tmps == 0: // 零
			//  是否有更高级的，如果有则不用加
			if 0 < num && num < 1000 {
				if res != "" {
					res = " " + units[cnt] + res
				} else {
					res = " " + units[cnt]
				}
			}
		case tmps < 100:
			t := ""
			if tmps < 20 {
				t = oneToNineteen[tmps] + res
			} else {
				t += mtens[tmps/10-2]
				if tmps%10 != 0 {
					t += " " + oneToNineteen[tmps%10]
				}
			}
			if 0 < num && num < 1000 {
				if res != "" {
					res = " " + units[cnt] + res
				} else {
					res = " " + units[cnt]
				}
			}
		default:
			t := ""
			if tmps/100 != 0 {
				t += oneToNineteen[tmps/100] + " Hundred "
			}
			tmps = tmps % 100
			if tmps < 20 {
				// 这里不用判断0，上面的case已经默认这里不可能是0
				t += oneToNineteen[tmps]
			} else {
				t += mtens[tmps/10-2]
				if tmps%10 != 0 {
					t += " " + oneToNineteen[tmps%10]
				}
			}
			if num != 0 {
				t = " " + units[cnt] + " " + t
			}
			res = t + res
		}
		//// 取出最后三个数
		//tmps := num % 1000
		//num = num / 1000
		//cnt++
		//switch {
		//case tmps%100 == 0: // 整百
		//	if tmps/100 != 0 {
		//		if res!=""{
		//			res = oneToNineteen[tmps/100] + " Hundred " + res
		//		}else{
		//			res = oneToNineteen[tmps/100] + " Hundred"
		//		}
		//	}
		//	//  是否有更高级的，如果有则不用加
		//	if 0 < num && num < 1000 {
		//		if res!=""{
		//			res = " " + units[cnt]+ " " + res
		//		}else{
		//			res = " " + units[cnt]
		//		}
		//	}
		//default:
		//	t := ""
		//	if tmps/100 != 0 {
		//		t += oneToNineteen[tmps/100] + " Hundred "
		//	}
		//	tmps = tmps % 100
		//	if tmps < 20 {
		//		// 这里不用判断0，上面的case已经默认这里不可能是0
		//		t += oneToNineteen[tmps]
		//	} else {
		//		t += mtens[tmps/10-2]
		//		if tmps%10 != 0 {
		//			t += " " + oneToNineteen[tmps%10]
		//		}
		//	}
		//	if num != 0 {
		//		t = " " + units[cnt] + " " + t
		//	}
		//	res = t + res
		//}

		//t := ""
		//// 取出三个数
		//for i := 2; i >= 0; i-- {
		//	tmp[i] = num % 10
		//	num = num / 10
		//	cnt++
		//}
		//if tmp[0] != 0 {
		//	t += oneToNineteen[tmp[0]] + " Hundred "
		//}
		//if tmp[1] == 0 {
		//	if tmp[2] != 0 {
		//		t += oneToNineteen[tmp[2]]
		//	}
		//} else if tmp[1] == 1 {
		//	// 加前面的
		//	t += " " + oneToNineteen[10+tmp[2]]
		//} else {
		//	// 加后面的
		//	t += " " + mtens[tmp[1]-2]
		//	if tmp[2] != 0 {
		//		t += " " + oneToNineteen[tmp[2]]
		//	}
		//}
		////  unit
		//if num != 0 {
		//	t = " " + units[cnt/3] + " " + t
		//}
		//res = t + res
	}
	return res
}

func NumberToWords(num int) string {
	oneToNineteen := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	mtens := []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	//units := []string{"Thousand", "Million", "Billion"} // 1000
	lessThanHundred := func(num int) string {
		if num < 20 {
			return oneToNineteen[num]
		}
		if num%10 == 0 {
			return mtens[num/10-2]
		} else {
			return mtens[num/10-2] + " " + oneToNineteen[num%10]
		}
	}
	lessThanThousand := func(num int) string {
		if num <= 99 {
			return lessThanHundred(num)
		}
		t := lessThanHundred(num % 100)
		if t != "Zero" {
			return oneToNineteen[num/100] + " Hundred " + t
		} else {
			return oneToNineteen[num/100] + " Hundred"
		}
	}
	// 分组
	if num < 1000 {
		return lessThanThousand(num)
	}
	if num < 1000000 { //Thousand
		t := lessThanThousand(num % 1000)
		if t != "Zero" {
			return lessThanThousand(num/1000) + " Thousand " + t
		}
		return lessThanThousand(num/1000) + " Thousand"
	}
	if num < 1000000000 { // Million
		if num%1000000 == 0 {
			return lessThanThousand(num/1000000) + " Million"
		}
		return lessThanThousand(num/1000000) + " Million " + NumberToWords(num%1000000)
	}
	if num < 1000000000000 { // Billion
		if num%1000000000 == 0 {
			return lessThanThousand(num/1000000000) + " Billion"
		}
		return lessThanThousand(num/1000000000) + " Billion " + NumberToWords(num%1000000000)
	}
	res := ""
	a := num % 1000000000000
	if a != 0 {
		res = NumberToWords(a)
	}
	num = num / 1000000000000
	if res != "" {
		return NumberToWords(num) + " Billion " + res
	} else {
		return NumberToWords(num) + " Billion"
	}
}
