package question

func CheckPalindromeFormation(a string, b string) bool {
	//分隔两个字符串，分隔后的串可为空
	//再次交叉拼接，是否能够形成回文串
	// 最暴力的方法是从0开始尝试即可，会超时
	//check := func(a string) bool {
	//	i := 0
	//	for ; i < len(a)/2 && a[i] == a[len(a)-i-1]; i++ {
	//
	//	}
	//	if i == len(a)/2 {
	//		return true
	//	}
	//	return false
	//}
	//for i := 0; i < len(a); i++ {
	//	if check(a[:i]+b[i:]) || check(b[:i]+a[i:]){
	//		return true
	//	}
	//}
	//return false

	//考虑上面的方法中，有很大一部分是在做重复检查
	// a[:i] b[i:] => a[:i+1] b[i+1:]
	// b[:i] a[i:] => b[:i+1] a[i+1:]
	// 变化是将a[i]放到第一个组合中，而b[i]被放到第二个组合中，相当于将他们二者做交换...
	// adbef
	// fecab
	// ad bef => adb ef => adbe f
	// fe cab => fec ab => feca b
	// ad[c]ab fe[b]ef => ad[b]ab fe[c]ef
	// adb[a]b fec[e]f => adb[e]b fec[a]f
	//TODO:看懂这里的内涵
	check := func(a string, b string, left int) int {
		right := len(a) - 1 - left
		for left >= 0 && right < len(a) {
			if a[left] != b[right] {
				break
			}
			left--
			right++
		}
		return left
	}
	// 首先从中间往外扩
	left := len(a)/2 - 1
	left = min(check(a, a, left), check(b, b, left)) // 尽可能往外扩
	left = min(check(a, b, left), check(b, a, left)) // 一次机会拼接
	return left == -1
	//TODO：看懂这里的C++
	//class Solution {
	//    bool check(string a, string b) {
	//        int n = a.size();
	//        bool flag = true;
	//        for (int i = 0; i < n / 2; ++i) {
	//            if (flag) {
	//                if (a[i] != b[n - 1 - i])
	//                    flag = false;
	//            }
	//            if (!flag)
	//                if (a[i] != a[n - 1 - i])
	//                    return false;
	//        }
	//        return true;
	//    }
	//public:
	//    bool checkPalindromeFormation(string a, string b) {
	//        if (check(a, b) || check(b, a))
	//            return true;
	//        reverse(a.begin(), a.end());
	//        reverse(b.begin(), b.end());
	//        if (check(a, b) || check(b, a))
	//            return true;
	//        return false;
	//    }
	//};
	//

}
