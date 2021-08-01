package main

import (
	"fmt"
	"github.com/depressi0n/leetcode/question"
)

// 对边界条件的技巧：
// （1）输入可能为空值或零值
// （2）输入可能很大
// （3）单向性的情况（如全部是负数）
// （4）申请的变量一定要注意取值范围，不要越界
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	length1 := len(nums1)
//	length2 := len(nums2)
//	if length1 == 0 {
//		if length2%2 == 0 {
//			return float64(nums2[length2/2-1]+nums2[length2/2]) / 2
//		}
//		return float64(nums2[(length2+1)/2-1])
//	}
//	if length2 == 0 {
//		if length1%2 == 0 {
//			return float64(nums1[length1/2-1]+nums1[length1/2]) / 2
//		}
//		return float64(nums1[(length1+1)/2-1])
//	}
//	length := length1 + length2
//	if length%2 == 0 {
//		return float64(findMinK(length/2, nums1, nums2)+findMinK(length/2+1, nums1, nums2)) / 2
//	} else {
//		return float64(findMinK((length+1)/2, nums1, nums2))
//	}
//
//}
//func findMinK(k int, num1 []int, num2 []int) int {
//	length1 := len(num1)
//	length2 := len(num2)
//	if length1 == 0 {
//		return num2[k-1]
//	}
//	if length2 == 0 {
//		return num1[k-1]
//	}
//	if k == 1 {
//		if num1[0] < num2[0] {
//			return num1[0]
//		} else {
//			return num2[0]
//		}
//	}
//	i, j := 0, 0
//	if length1 < k/2 {
//		i = length1
//	} else {
//		i = k / 2
//	}
//	if length2 < k/2 {
//		j = length2
//	} else {
//		j = k / 2
//	}
//	if num1[i-1] < num2[j-1] {
//		return findMinK(k-i, num1[i:], num2)
//	} else {
//		return findMinK(k-j, num1, num2[j:])
//	}
//}
//func longestPalindrome(s string) string {
//	if len(s) == 0 {
//		return ""
//	}
//	var newS []byte
//	length := 0
//	t := []byte(s)
//	for i := 0; i < len(t); i++ {
//		newS = append(newS, t[i])
//		length += 1
//		if i != len(t)-1 {
//			newS = append(newS, byte('/'))
//			length += 1
//		}
//	}
//	//默认第一个元素自己成为回文串，故
//	id, R := 0, 0
//	var arr []int
//	arr = append(arr, 0)
//	maxID := 0 //记录最长回文串的中心点 //可以用一个变量解决，在arr中的下标
//	maxR := 0  //记录最长回文串的半径
//	for i := 1; i < length; i++ {
//		//r := 0
//		if i >= id+R { //暴力扩
//			r := 0
//			for i+r < length && i-r >= 0 && newS[i-r] == newS[i+r] {
//				r++
//				if i+r > length || i-r < 0 {
//					break
//				}
//			}
//			r -= 1
//			if i+r > id+R {
//				id = i
//				R = r
//			}
//			arr = append(arr, r) //更新半径数组
//			if r > maxR {        //更新最大长度
//				maxID = id
//				maxR = R
//			}
//		} else {
//			r := arr[2*id-i]
//			if i+r < id+R {
//				arr = append(arr, r)
//			} else {
//				for i+r < length && i-r >= 0 && newS[i+r] == newS[i-r] {
//					r++
//					if i+r > length || i-r < 0 {
//						break
//					}
//				}
//				r -= 1
//				arr = append(arr, r)
//				if i+r > id+R {
//					R = r
//					id = i
//				}
//				if r > maxR {
//					maxID = i
//					maxR = r
//				}
//			}
//		}
//	}
//	//fmt.Printf("%s\n",string(newS[maxID-maxR:maxID+maxR+1]))
//	if maxID%2 == 0 {
//		return s[maxID/2-maxR/2 : maxID/2+maxR/2+1]
//	} else {
//		return s[(maxID-1)/2-(maxR-1)/2 : (maxID+1)/2+(maxR-1)/2+1]
//	}
//	//return string(newS)
//}
//func convert(s string, numRows int) string {
//	if numRows == 1 {
//		return s
//	}
//	b := make([]string, numRows)
//	i, k := 0, 0
//	flag := true //down,otherwise up
//	for k < len(s) {
//		b[i] += string(s[k])
//		if flag {
//			i++
//			if i == numRows-1 {
//				flag = false
//			}
//		} else {
//			i--
//			if i == 0 {
//				flag = true
//			}
//		}
//		k++
//	}
//	res := ""
//	for i := 0; i < numRows; i++ {
//		res += b[i]
//	}
//	return res
//}
//func reverse(x int) int {
//	var res int
//
//	for x != 0 {
//		if int(res)*10/10 != res {
//			return 0
//		}
//		res = res*10 + x%10
//		x = x / 10
//	}
//	return res
//}
//func myAtoi(str string) int {
//
//	s := strings.TrimSpace(str)
//	if len(s) == 0 {
//		return 0
//	}
//	flag := true
//	res := 0
//	if '+' == s[0] || '-' == s[0] {
//		if '-' == s[0] {
//			flag = false
//		}
//		s = s[1:]
//	}
//	if len(s) == 0 || s[0] < '0' || s[0] > '9' {
//		return 0
//	}
//	for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
//		res = res*10 + int(s[0]-'0')
//		if !flag && res*(-1) < math.MinInt32 {
//			return math.MinInt32
//		}
//		if flag && res > math.MaxInt32 {
//			return math.MaxInt32
//		}
//		s = s[1:]
//	}
//
//	if flag {
//		return res
//	} else {
//		return -res
//	}
//}
//func isPalindrome(x int) bool {
//	if (x != 0 && x%10 == 0) || x < 0 {
//		return false
//	}
//	p := 0
//	for x > p {
//		p = p*10 + x%10
//		x = x / 10
//	}
//	return x == p || x == p/10
//
//}
//func isMatch1(s string, p string) bool {
//	m, n := len(s), len(p)
//	matches := func(i, j int) bool {
//		if i == 0 {
//			return false
//		}
//		if p[j-1] == '.' {
//			return true
//		}
//		return s[i-1] == p[j-1]
//	}
//	dp := make([][]bool, m+1)
//	for i := 0; i < m+1; i++ {
//		dp[i] = make([]bool, n+1)
//	}
//	dp[0][0] = true
//	for i := 0; i < m+1; i++ {
//		for j := 1; j < n+1; j++ {
//			if p[j-1] == '*' {
//				dp[i][j] = dp[i][j] || dp[i][j-2]
//				if matches(i, j-1) {
//					dp[i][j] = dp[i][j] || dp[i-1][j]
//				}
//			} else if matches(i, j) {
//				dp[i][j] = dp[i][j] || dp[i-1][j-1]
//			}
//		}
//	}
//	return dp[m][n]
//}
//func maxArea(height []int) int {
//	i, j := 0, len(height)-1
//	max := 0
//	t := 0
//	for {
//		if i >= j {
//			break
//		}
//		if height[i] < height[j] {
//			t = (j - i) * height[i]
//			if t > max {
//				max = t
//			}
//			i++
//		} else if height[i] > height[j] {
//			t = (j - i) * height[j]
//			if t > max {
//				max = t
//			}
//			j--
//		} else {
//			t = (j - i) * height[j]
//			if t > max {
//				max = t
//			}
//			j--
//			i++
//		}
//	}
//	return max
//}
//
//var m = map[int]string{0: "", 1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M", 2000: "MM", 3000: "MMM"}
//var c = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
//
//func intToRoman(num int) string {
//	s, ok := m[num]
//	if ok {
//		return s
//	}
//	var res string
//	for i := len(c) - 1; i >= 0; i-- {
//		for num >= c[i] {
//			res += m[c[i]]
//			num -= c[i]
//			if num == 0 {
//				break
//			}
//		}
//		if num == 0 {
//			break
//		}
//	}
//	return res
//}
//
//var rm = map[string]int{"": 0, "I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000, "MM": 2000, "MMM": 3000}
//var rs = []string{"IV", "IX", "XL", "XC", "CD", "CM"}
//
//func romanToInt(s string) int {
//	res, ok := rm[s]
//	if ok {
//		return res
//	}
//	res = 0
//OK:
//	for i := 0; i < len(s); {
//		if s[i] == 'I' || s[i] == 'X' || s[i] == 'C' {
//			for j := 0; j < len(rs); j++ {
//				if i+1 < len(s) && s[i:i+2] == rs[j] {
//					res += rm[rs[j]]
//					i += 2
//					continue OK
//				}
//			}
//		}
//		res += rm[s[i:i+1]]
//		i++
//	}
//	return res
//}
//
//func longestCommonPrefix(strs []string) string {
//	if len(strs) == 0 {
//		return ""
//	}
//	if len(strs) == 1 {
//		return strs[0]
//	}
//	common := longestCommonPrefixTwo(strs[0], strs[1])
//	if common == "" {
//		return common
//	}
//	for i := 2; i < len(strs); i++ {
//		common = longestCommonPrefixTwo(common, strs[i])
//		if common == "" {
//			return common
//		}
//	}
//	return common
//}
//func longestCommonPrefixTwo(str1, str2 string) string {
//	if str1 == "" || str2 == "" {
//		return ""
//	}
//	var res string
//	i := 0
//	for {
//		if i < len(str1) && i < len(str2) && str1[i] == str2[i] {
//			res += str1[i : i+1]
//			i++
//		} else {
//			return res
//		}
//	}
//}
//
//// 常规
//func threeSum(nums []int) [][]int {
//	m := make(map[int]int)
//	for i := 0; i < len(nums); i++ {
//		v, ok := m[nums[i]]
//		if ok {
//			m[nums[i]] = v + 1
//		} else {
//			m[nums[i]] = 1
//		}
//	}
//	res := *new([][]int)
//	for i := 0; i < len(nums); i++ {
//		for j := 0; j < len(nums); j++ {
//			a, b, c := nums[i], nums[j], -nums[i]-nums[j]
//			if a == b && b == c { //a=b=c
//				if m[a] >= 3 {
//					res = appendInto(res, []int{a, b, c})
//				}
//			} else if a == b { //a=b!=c
//				if m[a] >= 2 && m[c] >= 1 {
//					res = appendInto(res, []int{a, b, c})
//				}
//			} else if b == c { //b=c!=a
//				if m[b] >= 2 && m[a] >= 1 {
//					res = appendInto(res, []int{a, b, c})
//				}
//			} else if a == c { //a=c!=b
//				if m[a] >= 2 && m[b] >= 1 {
//					res = appendInto(res, []int{a, b, c})
//				}
//			} else if m[a] >= 1 && m[b] >= 1 && m[c] >= 1 { //都不相等
//				res = appendInto(res, []int{a, b, c})
//			}
//		}
//	}
//	return res
//}
//func appendInto(r [][]int, s []int) [][]int {
//	sort.Ints(s)
//	flag := false
//	for i := 0; i < len(r); i++ {
//		if r[i][0] == s[0] && r[i][1] == s[1] {
//			flag = true
//			break
//		}
//	}
//	if !flag {
//		r = append(r, s)
//	}
//	return r
//}
//func threeSum1(nums []int) [][]int {
//	res := *new([][]int)
//	sort.Ints(nums)
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > 0 {
//			break
//		}
//		j := i + 1
//		k := len(nums) - 1
//		for {
//			if j >= k {
//				break
//			}
//			if nums[i]+nums[j]+nums[k] == 0 {
//				res = append(res, []int{nums[i], nums[j], nums[k]})
//				for k-1 >= j && nums[k-1] == nums[k] {
//					k--
//				}
//				k--
//				for j+1 <= k && nums[j+1] == nums[j] {
//					j++
//				}
//				j++
//			} else if nums[i]+nums[j]+nums[k] > 0 {
//				for k-1 >= j && nums[k-1] == nums[k] {
//					k--
//				}
//				k--
//			} else {
//				for j+1 <= k && nums[j+1] == nums[j] {
//					j++
//				}
//				j++
//			}
//		}
//		for i+1 < len(nums) && nums[i+1] == nums[i] {
//			i++
//		}
//	}
//	return res
//}
//func threeSumClosest(nums []int, target int) int {
//	best := math.MaxInt32
//	sort.Ints(nums)
//	update := func(cur int) {
//		if abs(cur-target) < abs(best-target) {
//			best = cur
//		}
//	}
//	for i := 0; i < len(nums); i++ {
//		//跳过重复值，不重复枚举
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		j, k := i+1, len(nums)-1
//		for j < k {
//			sum := nums[i] + nums[j] + nums[k]
//			if sum == target {
//				return target
//			}
//			update(sum)
//			if sum > target {
//				k0 := k - 1
//				//跳过重复只
//				for j < k0 && nums[k0] == nums[k] {
//					k0--
//				}
//				k = k0
//			}
//			if sum < target {
//				j0 := j + 1
//				// 跳过重复值
//				for j0 < k && nums[j0] == nums[j] {
//					j0++
//				}
//				j = j0
//			}
//		}
//	}
//	return best
//}
//func abs(x int) int {
//	if x < 0 {
//		return -1 * x
//	}
//	return x
//}
//
//func letterCombinations(digits string) []string {
//	if len(digits) == 0 {
//		return []string{""}
//	}
//	var NumberToLetter = map[byte]byte{
//		'2': 'a',
//		'3': 'd',
//		'4': 'g',
//		'5': 'j',
//		'6': 'm',
//		'7': 'p',
//		'8': 't',
//		'9': 'w',
//	}
//	var NumberToLen = map[byte]int{
//		'2': 3,
//		'3': 3,
//		'4': 3,
//		'5': 3,
//		'6': 3,
//		'7': 4,
//		'8': 3,
//		'9': 4,
//	}
//	// 根据数字2，对应找到a，长度为3，3对应找到d，长度为3...7对应p，长度为4，9对应w，长度4
//	// 给定字符串对应的输出的字符串长度肯定相同，所以这里可以利用这个相同，依次往后变化
//	res := *new([]string)
//	//首先将输入的字符串转成输出的第一个字符串
//	var tmp string
//	for i := 0; i < len(digits); i++ {
//		tmp += string(NumberToLetter[digits[i]])
//	}
//	//将这个字符串及其最后一个字母的扩展放入结果中
//	res = append(res, tmp)
//	for i := 1; i < NumberToLen[digits[len(digits)-1]]; i++ {
//		t := []byte(tmp)
//		t[len(digits)-1] += byte(i)
//		res = append(res, string(t))
//	}
//	index := len(digits) - 2 //记录这个index是为了记录该到了第几个字母
//	prevLen := len(res)      //记录这个长度是为了记住下一次遍历到那里结束
//	//每次更改一个字母
//	for ; index >= 0; index-- {
//		for i := 1; i < NumberToLen[digits[index]]; i++ { //这次要改多少个
//			for j := 0; j < prevLen; j++ {
//				t := []byte(res[j])
//				t[index] = byte(int(t[index]) + i)
//				res = append(res, string(t))
//			}
//		}
//		prevLen = len(res)
//	}
//	return res
//}
//func fourSum(nums []int, target int) [][]int {
//	if len(nums) < 4 {
//		return nil
//	}
//	sort.Ints(nums)
//	res := *new([][]int)
//	for i := 0; i < len(nums)-3; i++ { //最后必须有三个数
//		for j := i + 1; j < len(nums)-2; j++ { //最后必须有两个数
//			p, q := j+1, len(nums)-1
//			for p < q {
//				if nums[i]+nums[j]+nums[p]+nums[q] == target {
//					res = append(res, []int{nums[i], nums[j], nums[p], nums[q]}) //等于的话就记录然后减小q增大p
//					for {
//						if p < q && nums[p+1] == nums[p] { //跳过相同的元素
//							p = p + 1
//						} else {
//							break
//						}
//					}
//					p++
//					for {
//						if p < q && nums[q-1] == nums[q] { //跳过相同的元素
//							q = q - 1
//						} else {
//							break
//						}
//					}
//					q--
//				} else if nums[i]+nums[j]+nums[p]+nums[q] < target { //比目标值小，增大p
//					for {
//						if p < q && nums[p+1] == nums[p] { //跳过相同的元素
//							p = p + 1
//						} else {
//							break
//						}
//					}
//					p++
//				} else { // 比目标值大，减小q
//					for {
//						if p < q && nums[q-1] == nums[q] { //跳过相同的元素
//							q = q - 1
//						} else {
//							break
//						}
//					}
//					q--
//				}
//			}
//			for j+1 < len(nums)-2 && nums[j+1] == nums[j] {
//				j++
//			}
//		}
//		for i+1 < len(nums)-3 && nums[i+1] == nums[i] {
//			i++
//		}
//	}
//	return res
//}
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	p := make([]*ListNode, 2)
//	p[1] = head
//	for i := 0; i < n; i++ {
//		p[1] = p[1].Next
//	}
//	for p[1].Next != nil {
//		p[1] = p[1].Next
//		p[0] = p[0].Next
//	}
//	if p[0] == head {
//		head = head.Next
//		p[0] = nil
//		return head
//	}
//	p[0].Next = p[0].Next.Next
//	return head
//}
//func isValid(s string) bool {
//	stack := make([]byte, len(s))
//	cur := 0
//	i := 0
//	for i < len(s) {
//		switch s[i] {
//		case '(':
//			stack[cur] = '('
//			cur++
//		case ')':
//			if cur-1 < 0 {
//				return false
//			}
//			if stack[cur-1] != '(' {
//				return false
//			} else {
//				cur--
//			}
//		case '[':
//			stack[cur] = '['
//			cur++
//		case ']':
//			if cur-1 < 0 {
//				return false
//			}
//			if stack[cur-1] != '[' {
//				return false
//			} else {
//				cur--
//			}
//		case '{':
//			stack[cur] = '{'
//			cur++
//		case '}':
//			if cur-1 < 0 {
//				return false
//			}
//			if stack[cur-1] != '{' {
//				return false
//			} else {
//				cur--
//			}
//		}
//		if cur < 0 {
//			return false
//		}
//		i++
//	}
//	if cur == 0 {
//		return true
//	} else {
//		return false
//	}
//}
//
//// 设置一个头节点之后把这个头节点删除会更简单一些，最开始的长if就不用了
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	var res, cur *ListNode
//	res = &ListNode{
//		Val:  -1,
//		Next: nil,
//	}
//	cur = res
//	for l1 != nil && l2 != nil {
//		if minNode(l1, l2) == l1 { //这里也可以直接判断值
//			cur.Next = l1
//			l1 = l1.Next
//			cur = cur.Next
//			//cur.Next=nil
//
//		} else {
//			cur.Next = l2
//			l2 = l2.Next
//			cur = cur.Next
//		}
//	}
//	if l1 != nil {
//		cur.Next = l1
//
//	}
//	if l2 != nil {
//		cur.Next = l2
//	}
//	l1 = nil
//	l2 = nil
//	cur = nil
//	return res.Next
//}
//func minNode(p, q *ListNode) *ListNode {
//	if p.Val < q.Val {
//		return p
//	}
//	return q
//}
//
//// 另外一条思路貌似使用dp做的，因为这里确实包含了很多的重复计算，可以把这些重复计算省下来，大致的方向是，先得到之前m-1的结果，然后遍历i：1-n/2，通过拼接就可以得到答案
//func generateParenthesis(n int) []string {
//	// 用递归会更好,但怎么解决重复的问题？ 解决方案是通过对比左右括号数目，无论如何右括号肯定小于左括号，而当没达到最大左括号，可以继续先插左，但是当插入完左括号以后说明可以在
//	// 之后的任意位置插入右括号
//	if n == 0 {
//		return []string{}
//	}
//	if n == 1 {
//		return []string{"()"}
//	}
//	var res []string
//	generateParenthesisDFS(n, n, "", &res)
//	return res
//}
//
//func generateParenthesisDFS(left int, right int, s string, res *[]string) {
//	if right == 0 { //右括号插入完毕说明完毕
//		*res = append(*res, s)
//		return
//	}
//	if left > 0 { //当左括号没有被插完时，可以继续插入左括号
//		generateParenthesisDFS(left-1, right, s+"(", res)
//	}
//	if right > left { //这个表示之前插入过左括号，可在之后的任意位置插入右括号，
//		generateParenthesisDFS(left, right-1, s+")", res)
//	}
//}
//func mergeKLists(lists []*ListNode) *ListNode { //败者树或者归并合并
//	k := len(lists)
//	for k != -1 {
//		k = mergeList(lists, k)
//	}
//	return lists[0]
//}
//func mergeList(lists []*ListNode, length int) int { //两两合并
//	if length == 1 {
//		return -1
//	}
//	k := 0
//	for i := 0; i < length; i += 2 {
//		if i+1 < length {
//			lists[k] = mergeTwoLists(lists[i], lists[i+1])
//		} else {
//			lists[k] = lists[i]
//		}
//		k++
//	}
//	return k
//}
//func swapPairs(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//	}
//	p := head
//	q := p.Next
//	p.Next = q.Next
//	q.Next = p
//	head = q
//	q = p.Next
//	for {
//		if q == nil || q.Next == nil {
//			return head
//		}
//		r := q.Next.Next
//		q.Next.Next = q
//		p.Next = q.Next
//		q.Next = r
//		p = q
//		q = p.Next
//	}
//}
//
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	if k == 1 { //如果k为1直接return
//		return head
//	}
//	s := &ListNode{ //加上一个哨兵节点，统一操作
//		Val:  -1,
//		Next: head,
//	}
//	pre := s
//	r := pre
//	var p, q, ppre *ListNode
//	for r.Next != nil {
//		for i := 0; i < k; i++ {
//			if r.Next == nil { //如果在往后移动过程中出现了nil，说明是最后一组且不满k个，可以直接返回
//				return s.Next
//			}
//			r = r.Next
//		}
//		//能走到这里没有return的话说明满K个,完成换位
//		p = pre.Next
//		q = p.Next
//		pre.Next = r.Next //将中间这一组断开
//		//r=r.Next//将r指向可能的下一组
//		ppre = pre               //更新之前的pre
//		pre = p                  // 下一组的新的pre
//		for i := 0; i < k; i++ { //头插法
//			p.Next = ppre.Next
//			ppre.Next = p
//			if i == k-1 {
//				p = nil
//				q = nil
//			} else {
//				p = q
//				q = q.Next
//			}
//		}
//		r = pre //将r回到已改好的最后一个，和pre一致
//	}
//	return s.Next
//}
//func removeDuplicates(nums []int) int {
//	if len(nums) == 0 || len(nums) == 1 {
//		return len(nums)
//	}
//	k := 1
//	for i := 1; i < len(nums); i++ {
//		if nums[i] != nums[k-1] {
//			nums[k] = nums[i]
//			k++
//		}
//	}
//	return k
//}
//func removeElement(nums []int, val int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	k := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] != val {
//			nums[k] = nums[i]
//			k++
//		}
//	}
//	return k
//}
//func strStr(haystack string, needle string) int {
//	if len(needle) == 0 {
//		return 0
//	}
//	for i := 0; i < len(haystack) && len(haystack)-i >= len(needle); i++ {
//		j := 0
//		for i < len(haystack) && j < len(needle) && haystack[i] == needle[j] {
//			i++
//			j++
//		}
//		if j != len(needle) {
//			i = i - j
//		} else {
//			return i - j
//		}
//	}
//	return -1
//}
//func divide(dividend int, divisor int) int { //在移位运算和加减法的条件下进行计算，使用二分查找
//	if dividend == 0 || divisor == 0 {
//		return 0
//	}
//	if divisor == -1 && dividend == -math.MaxInt32 {
//		return math.MaxInt32
//	}
//	flag := 0
//	dividendAbs := dividend
//	if dividend < 0 {
//		dividendAbs = -dividend
//		flag += 1
//	}
//
//	divisorAbs := divisor
//	if divisor < 0 {
//		divisorAbs = -divisor
//		flag += 1
//	}
//	result := div(dividendAbs, divisorAbs)
//	if flag == 1 {
//		if result > math.MaxInt32+1 {
//			return -(math.MaxInt32) + 1
//		}
//		return -result
//	}
//	if result > math.MaxInt32 {
//		return math.MaxInt32
//	}
//	return result
//}
//func div(dividend int, divisor int) int {
//	result := 0
//	for dividend >= divisor {
//		i := 0
//		j := 1
//		for dividend >= (divisor << i) {
//			j <<= 1
//			i += 1
//		}
//		result += j >> 1
//		dividend -= divisor << (i - 1)
//	}
//	return result
//}
//func findSubstring(s string, words []string) []int {
//	//words中长度都相同，这个应该是可以利用的
//	//做一个map，然后计数，因为一个string可能重复出现，然后匹配的时候就用这个map去匹配
//	//主串后移时应该一个一个后移
//	// 特殊的输入情况：s=""，words=[]; s="",words=[""]; s="sd",word=[]; s="sd",word=[""];
//	//
//	var res []int
//	if len(words) == 0 {
//		return nil
//	}
//	length := len(words[0]) //查看words的长度
//	if length == 0 {
//		if len(s) == 0 {
//			return []int{0}
//		}
//		for i := 0; i <= len(s); i++ { //这个地方为什么一定要加上等于号？
//			res = append(res, i)
//		}
//		return res
//	}
//	m := make(map[string]int)
//	for i := 0; i < len(words); i++ {
//		if k, ok := m[words[i]]; ok {
//			m[words[i]] = k + 1
//		} else {
//			m[words[i]] = 1
//		}
//	}
//	for i := 0; i < len(s); i++ {
//		windows := make(map[string]int)
//		for k, v := range m {
//			windows[k] = v
//		}
//		flag := true
//		if i+length*len(words) > len(s) {
//			break
//		}
//		for j := i; j < i+length*len(words); j += length {
//			substr := s[j : j+length]
//			if k, ok := windows[substr]; ok {
//				windows[substr] = k - 1
//				if k == 1 {
//					delete(windows, substr)
//				}
//			} else {
//				flag = false
//				break
//			}
//		}
//		if flag {
//			res = append(res, i)
//		}
//	}
//	return res
//}
//func findSubstring1(s string, words []string) []int {
//	var result []int // 最终结果
//	// 特判
//	if len(s) == 0 || len(words) == 0 {
//		return result
//	}
//	// 统计words中每个单词出现的次数
//	need := make(map[string]int)
//	for _, word := range words {
//		need[word]++
//	}
//	wordLen, wordNum, ls := len(words[0]), len(words), len(s)
//	totalLen := wordLen * wordNum              // 子串总长度
//	for start := 0; start < wordLen; start++ { //可能的起跳长度都需要遍历一遍
//		// 窗口左边，窗口右边，词频匹配的单词数
//		left, right, match := start, start, 0
//		// 窗口中每个单词出现的次数
//		window := make(map[string]int)
//		for right+wordLen <= ls {
//			rightWord := s[right : right+wordLen] // 当前新加入的单词
//			right += wordLen
//			if need[rightWord] > 0 {
//				window[rightWord]++
//				if window[rightWord] == need[rightWord] {
//					match++
//				}
//			}
//			// 如果满足了长度，判断是否满足词频
//			if right-left == totalLen {
//				// 词频也匹配，加入结果
//				if match == len(need) {
//					result = append(result, left)
//				}
//				leftWord := s[left : left+wordLen] // 当前新移出的单词，这里跳跃都是一次一个单词
//				left += wordLen
//				if need[leftWord] > 0 {
//					if window[leftWord] == need[leftWord] {
//						match--
//					}
//					window[leftWord]--
//				}
//			}
//		}
//	}
//	return result
//}
//
//func longestValidParentheses(s string) int {
//	//去掉左边的右括号,这里是可以肯定没有问题的，因为左边的右括号肯定是无法匹配的,反之亦然,但这一步可以不用管，因为下面的算法会跳过这种情况
//	//from left to right
//	left := 0
//	right := 0
//	max := 0
//	for i := 0; i < len(s); i++ {
//		if s[i] == '(' {
//			left++
//		}
//		if s[i] == ')' {
//			right++
//		}
//		if right > left {
//			left = 0
//			right = 0
//		} else if left == right && max < 2*left {
//			max = 2 * left
//		}
//	}
//	//from right to left
//	left = 0
//	right = 0
//	for i := len(s) - 1; i >= 0; i-- {
//		if s[i] == '(' {
//			left++
//		}
//		if s[i] == ')' {
//			right++
//		}
//		if left > right {
//			left = 0
//			right = 0
//		} else if left == right && max < 2*left {
//			max = 2 * left
//		}
//	}
//	return max
//}
//func isMatch_wenhao_and_dian(s string, p string) bool {
//	dp := *new([][]bool)
//	for i := 0; i <= len(s); i++ {
//		tmp := make([]bool, len(p)+1)
//		dp = append(dp, tmp)
//	}
//	match := func(i, j int) bool {
//		if i == 0 {
//			return false
//		}
//		if p[j-1] == '.' {
//			return true
//		}
//		return s[i-1] == p[j-1]
//	}
//	//初始化数组，dp[i][j]表示以s[:i]和p[:j]是否匹配,所以最后结果应该是dp[len(s),len(p)]
//	//当i=0或者j=0时，字符串为空必定不匹配，但是空字符串和空字符串是匹配的
//	dp[0][0] = true //其他情况即空串和非空串，在匹配时认为是false
//	//状态转移方程
//	//当p[j-1]='*'时，需要看前面一位p[j-2],若s[i-1]=p[j-2],也就是能匹配，可以让这个*发挥表示多个的作用，则dp[i][j]=dp[i-1][j]；不能匹配时，发挥0个的作用，dp[i][j]=dp[i][j-2]
//	//当p[j-1]!='*'时，若s[i-1]=p[j-1]，dp[i][j]=dp[i-1][j-1];否则就为false
//	for i := 0; i <= len(s); i++ { //这里起始条件从0开始的目的是，可能会有p能匹配上空串
//		for j := 1; j <= len(p); j++ {
//			if p[j-1] == '*' {
//				dp[i][j] = dp[i][j] || dp[i][j-2]
//				if match(i, j-1) {
//					dp[i][j] = dp[i][j] || dp[i-1][j]
//				}
//			} else { //p[i-1]不是*的情况
//				if match(i, j) {
//					dp[i][j] = dp[i][j] || dp[i-1][j-1]
//				}
//			}
//		}
//	}
//	return dp[len(s)][len(p)]
//}
//func search(nums []int, target int) int {
//	if nums == nil || len(nums) == 0 {
//		return -1
//	}
//	start := 0
//	end := len(nums) - 1
//	var mid int
//	for start <= end {
//		mid = (start + end) / 2
//		if target == nums[mid] {
//			return mid
//		}
//		if nums[start] <= nums[mid] {
//			//左半有序
//			if nums[start] <= target && target <= nums[mid] {
//				end = mid - 1
//			} else {
//				start = mid + 1
//			}
//		} else {
//			//右半有序
//			if nums[mid] <= target && target <= nums[end] {
//				start = mid + 1
//			} else {
//				end = mid - 1
//			}
//		}
//	}
//	return -1
//}
//func searchRange(nums []int, target int) []int {
//	// 分析情况
//	// 1. 数组中不存在目标值，示例{1,1,1},2
//	// 2. 数组中存在唯一目标值且位于开头或且位于结尾，或且位于中间{1,2,2,2},1;{1,1,1,2},2;{1,2,3,4},2
//	// 3. 数组中存在多个目标值{1,2,2,2,2,3},2;{2,2,2,2},2
//	left := -1
//	right := -1
//	start := 0
//	end := len(nums) - 1
//	var mid int
//	for start <= end {
//		mid = (start + end) / 2
//		if nums[mid] < target {
//			start = mid + 1
//		} else if nums[mid] > target {
//			end = mid - 1
//		} else {
//			left = findLeft(nums, start, mid, target)
//			right = findRight(nums, mid, end, target)
//			return []int{left, right}
//		}
//	}
//	return []int{left, right}
//}
//func findLeft(nums []int, start, end int, target int) int {
//	var mid int
//	for start < end { //可以肯定的是这部分值必须小于等于目标值
//		mid = (start + end) / 2 //因为这里会向下取整
//		if nums[mid] == target {
//			end = mid
//		} else {
//			start = mid + 1
//		}
//	}
//	return start
//}
//func findRight(nums []int, start, end int, target int) int {
//	var mid int
//	for start < end {
//		mid = (start + end + 1) / 2 //因为这里需要向上取整
//		if nums[mid] == target {    //可以肯定的是这部分值必须大于等于目标值
//			start = mid
//		} else {
//			end = mid - 1
//		}
//	}
//	return end
//}
//func searchInsert(nums []int, target int) int {
//	if nums == nil || len(nums) == 0 || nums[0] > target {
//		return 0
//	}
//	if target > nums[len(nums)-1] {
//		return len(nums)
//	}
//	start := 0
//	end := len(nums)
//	var mid int
//	for start < end {
//		mid = (start + end) / 2
//		if nums[mid] == target {
//			return mid
//		} else if nums[mid] > target {
//			end = mid
//		} else {
//			start = mid + 1
//		}
//	}
//	if nums[mid] > target {
//		return mid
//	} else {
//		return mid + 1
//	}
//}
//func isValidSudoku(board [][]byte) bool {
//	m := make(map[int]int)
//	for i := 1; i < 10; i++ {
//		m[i] = 0
//	}
//	for i := 0; i < 9; i++ {
//		for j := 0; j < 9; j++ {
//			if board[i][j] != 0 && m[int(board[i][j])] != 0 {
//				return false
//			}
//			m[int(board[i][j])] = 1
//		}
//		for i := 1; i < 10; i++ {
//			m[i] = 0
//		}
//	}
//	for i := 0; i < 9; i++ {
//		for j := 0; j < 9; j++ {
//			if board[j][i] != 0 && m[int(board[j][i])] != 0 {
//				return false
//			}
//			m[int(board[j][i])] = 1
//		}
//		for i := 1; i < 10; i++ {
//			m[i] = 0
//		}
//	}
//
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			for k := 3 * i; k < 3*(i+1); k++ {
//				for n := 3 * j; n < 3*(j+1); n++ {
//					if board[k][n] != 0 && m[int(board[k][n])] != 0 {
//						return false
//					}
//					m[int(board[k][n])] = 1
//				}
//			}
//			for i := 1; i < 10; i++ {
//				m[i] = 0
//			}
//		}
//	}
//	return true
//}
//func solveSudoku(board [][]byte) {
//	h := make([][][]byte, 9)
//	for i := 0; i < 9; i++ {
//		h[i] = make([][]byte, 9)
//		for j := 0; j < 9; j++ {
//			h[i][j] = *new([]byte)
//		}
//	}
//	for i := 0; i < 9; i++ {
//		for j := 0; j < 9; j++ {
//			if board[i][j] != 0 {
//				h[i][j] = append(h[i][j], board[i][j])
//			} else {
//				for k := 0; k < 9; k++ {
//					h[i][j] = append(h[i][j], byte(k+1))
//				}
//			}
//		}
//	}
//	finish := func(b [][]byte) bool {
//		for i := 0; i < 9; i++ {
//			for j := 0; j < 9; j++ {
//				if b[i][j] == 0 {
//					return false
//				}
//			}
//		}
//		return true
//	}
//	for !finish(board) {
//		for i := 0; i < 9; i++ {
//			for j := 0; j < 9; j++ {
//				if len(h[i][j]) == 1 {
//					board[i][j] = h[i][j][0]
//					for k := 0; k < 9; k++ {
//						if len(h[i][k]) != 1 {
//							for m, n := range h[i][k] {
//								if n == h[i][j][0] {
//									h[i][k] = append(h[i][k][:m], h[i][k][m+1:]...)
//									if len(h[i][k]) == 1 {
//										board[i][k] = h[i][k][0]
//									}
//									break
//								}
//							}
//						}
//						if len(h[k][j]) != 1 {
//							for m, n := range h[k][j] {
//								if n == h[i][j][0] {
//									h[k][j] = append(h[k][j][:m], h[k][j][m+1:]...)
//									if len(h[k][j]) == 1 {
//										board[k][j] = h[k][j][0]
//									}
//									break
//								}
//							}
//						}
//					}
//				}
//			}
//		}
//		for i := 0; i < 3; i++ {
//			for j := 0; j < 3; j++ {
//				for m := 3 * i; m < 3*(i+1); m++ {
//					for n := 3 * j; n < 3*(j+1); n++ {
//						if len(h[m][n]) == 1 {
//							board[m][n] = h[m][n][0]
//							for p := 3 * i; p < 3*(i+1); p++ {
//								for q := 3 * j; q < 3*(j+1); q++ {
//									if len(h[p][q]) != 1 {
//										for a, b := range h[p][q] {
//											if b == h[m][n][0] {
//												h[p][q] = append(h[p][q][:a], h[p][q][a+1:]...)
//												if len(h[p][q]) == 1 {
//													board[p][q] = h[p][q][0]
//												}
//												break
//											}
//										}
//									}
//								}
//							}
//						}
//					}
//				}
//			}
//		}
//	}
//}
//func countAndSay(n int) string {
//	oldres := "1"
//	res := oldres
//	m := map[byte]string{
//		'1': "1",
//		'2': "2",
//		'3': "3",
//		'4': "4",
//		'5': "5",
//		'6': "6",
//		'7': "7",
//		'8': "8",
//		'9': "9",
//	}
//	for i := 1; i < n; i++ {
//		res = ""
//		for j := 0; j < len(oldres); j++ {
//			cnt := 1
//			for j+1 < len(oldres) && oldres[j+1] == oldres[j] {
//				cnt++
//				j++
//			}
//			res += m[byte('0'+cnt)] + oldres[j:j+1]
//		}
//		oldres = res
//	}
//	return res
//}
//func dfsCombinationsum(candidates []int, sum, index int, target int, have []int) [][]int {
//	res := new([][]int)
//	if sum+candidates[index] == target { //剪枝并结束
//		have = append(have, candidates[index])
//		newHave := make([]int, len(have))
//		copy(newHave, have)
//		*res = append(*res, newHave)
//		have = have[:len(have)-1]
//		return *res
//	}
//	if sum+candidates[index] < target { //可以继续递归
//		for i := index; i < len(candidates); i++ {
//			if sum+candidates[index]+candidates[i] <= target {
//				have = append(have, candidates[index])
//				*res = append(*res, dfsCombinationsum(candidates, sum+candidates[index], i, target, have)...)
//				have = have[:len(have)-1]
//			} else {
//				break
//			}
//		}
//	}
//	return *res
//}
//func combinationSum(candidates []int, target int) [][]int {
//	sort.Ints(candidates)
//	res := new([][]int)
//	for i := 0; i < len(candidates); i++ {
//		have := new([]int)
//		if candidates[i] > target {
//			break
//		}
//		*res = append(*res, dfsCombinationsum(candidates, 0, i, target, *have)...)
//	}
//	return *res
//}
//func dfsCombinationSum2(candidates []int, index int, target int, have *[]int, res *[][]int) {
//	if target == 0 { //存放结果,但这里可能重复，所以要检查是否是重复的
//		newHave := make([]int, len(*have))
//		copy(newHave, *have)
//		*res = append(*res, newHave)
//		return
//	}
//	for i := index; i >= 0; i-- {
//		if candidates[i] > target {
//			continue
//		} else {
//			*have = append(*have, candidates[i])
//			dfsCombinationSum2(candidates, i-1, target-candidates[i], have, res)
//			*have = (*have)[:len(*have)-1]
//			for i-1 >= 0 && candidates[i-1] == candidates[i] {
//				i--
//			}
//		}
//	}
//}
//func combinationSum2(candidates []int, target int) [][]int {
//	sort.Ints(candidates)
//	var have []int
//	have = *new([]int)
//	var res [][]int
//	res = *new([][]int)
//	dfsCombinationSum2(candidates, len(candidates)-1, target, &have, &res)
//	return res
//}
//func firstMissingPositive(nums []int) int {
//	for i := 0; i < len(nums); i++ {
//		if nums[i] <= 0 {
//			nums[i] = len(nums) + 1
//		}
//	}
//	for i := 0; i < len(nums); i++ {
//		if nums[i] < 0 && nums[i] >= -len(nums) && nums[-nums[i]-1] > 0 {
//			nums[-nums[i]-1] *= -1
//		} else if nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] > 0 {
//			nums[nums[i]-1] *= -1
//		}
//	}
//	i := 0
//	for ; i < len(nums); i++ {
//		if nums[i] > 0 {
//			break
//		}
//	}
//	return i + 1
//}
//func trap(height []int) int {
//	if height == nil || len(height) == 0 || len(height) == 1 || len(height) == 2 {
//		return 0
//	}
//	// 动态规划
//	//leftMaxIndex := make([]int, len(height))//记录index而不是value是为了之后的比较和压缩空间
//	//rightMaxIndex := make([]int, len(height))
//	//for i := 0; i < len(height); i++ {
//	//	if i == 0 { 	// 这个标志为此位置不可能存储雨水
//	//		leftMaxIndex[i] = len(height)
//	//	} else if i == 1 {	//左边最大值肯定是第1个
//	//		leftMaxIndex[i] = 0
//	//	} else if height[leftMaxIndex[i-1]] <= height[i-1] {//如果前一个位置的值大于等于前一个位置左边的最大值，则左侧最大值为前一个位置的值
//	//		leftMaxIndex[i] = i - 1
//	//	}else {// 否则和之前的最大值相同
//	//		leftMaxIndex[i]=leftMaxIndex[i-1]
//	//	}
//	//}
//	//for i := len(height) - 1; i >= 0; i-- {
//	//	if i == len(height)-1 {
//	//		rightMaxIndex[i] = len(height)
//	//	} else if i == len(height)-2 {
//	//		rightMaxIndex[i] = len(height) - 1
//	//	} else if height[rightMaxIndex[i+1]] <= height[i+1] {
//	//		rightMaxIndex[i] = i + 1
//	//	}else{
//	//		rightMaxIndex[i]=rightMaxIndex[i+1]
//	//	}
//	//计算局部最大值
//	//for i:=0;i<len(height);i++{
//	//	if i == 0 {
//	//		leftMaxIndex[i]=0
//	//	}else if i == len(height)-1 {
//	//		leftMaxIndex[i]=0
//	//	}else{
//	//		if height[leftMaxIndex[i]] > height[rightMaxIndex[i]] {//修改局部最大值的index
//	//			leftMaxIndex[i]=rightMaxIndex[i]
//	//		}
//	//		if height[i]>height[leftMaxIndex[i]] {//赋差值即雨水高度
//	//			leftMaxIndex[i]=0
//	//		}else{
//	//			leftMaxIndex[i]=height[leftMaxIndex[i]]-height[i]
//	//		}
//	//	}
//	//	leftMaxIndex[0]+=leftMaxIndex[i]
//	//}
//	//var res int
//	//for i:=0;i<len(height);i++{
//	//	res+=leftMaxIndex[i]
//	//}
//	//return leftMaxIndex[0]
//	//动态规划的优化
//	// 空间可以压缩为n，而不是2n,在求出右边最大值时直接得到一个差值放入leftMaxIndex中然后直接加起来即可
//	//for i := len(height) - 1; i >= 0; i-- {
//	//	if i == len(height)-1 || i == 0 {//右侧最大值肯定为0，所以不用比较
//	//		leftMaxIndex[i] = 0
//	//		continue
//	//	}
//	//	if i == len(height)-2 {//右侧最大值肯定为最后一个元素
//	//		if height[len(height)-1] <= height[leftMaxIndex[i]] {//右侧最大值比左侧最大值小
//	//			//局部最大值为右侧值，累加
//	//			if height[i]<height[len(height)-1]{
//	//				leftMaxIndex[len(height)-1]+=height[len(height)-1]-height[i]
//	//			}
//	//		}else{
//	//			//局部最大值为左侧最大值
//	//			if height[i]<height[leftMaxIndex[i]]{
//	//				leftMaxIndex[len(height)-1]+=height[leftMaxIndex[i]]-height[i]
//	//			}
//	//		}
//	//		leftMaxIndex[i]=len(height)-1//更新右侧最大值的index，为了后续的比较
//	//		continue
//	//	}
//	//	if height[leftMaxIndex[i+1]] <= height[i+1] {// 右侧最大值为右边元素
//	//		if height[i+1] > height[leftMaxIndex[i]] {//局部最大值为左侧最大值
//	//			if height[leftMaxIndex[i]]>height[i]{
//	//				leftMaxIndex[len(height)-1]+=height[leftMaxIndex[i]]-height[i]
//	//			}
//	//		}else {//局部最大值为右侧元素
//	//			if height[i+1]>height[i]{
//	//				leftMaxIndex[len(height)-1]+=height[i+1]-height[i]
//	//			}
//	//		}
//	//		leftMaxIndex[i]=i+1
//	//	}else{// 右侧最大值为右边元素的右侧最大值
//	//		if  height[leftMaxIndex[i+1]]> height[leftMaxIndex[i]] { //局部最大值为左侧最大值
//	//			if height[leftMaxIndex[i]]-height[i] >0{
//	//				leftMaxIndex[len(height)-1]+=height[leftMaxIndex[i]]-height[i]
//	//			}
//	//		}else{//局部最大值为右侧元素的右侧最大值
//	//			if height[leftMaxIndex[i+1]]-height[i] >0{
//	//				leftMaxIndex[len(height)-1]+=height[leftMaxIndex[i+1]]-height[i]
//	//			}
//	//		}
//	//		leftMaxIndex[i]=leftMaxIndex[i+1]
//	//	}
//	//}
//	//return leftMaxIndex[len(height)-1]
//	//TODO： 单调递减栈
//
//	// TOODO：双指针法
//	leftMax := 0
//	rightMax := 0
//	left := 0
//	right := len(height) - 1
//	res := 0
//	for left < right {
//		if height[left] < height[right] {
//			if leftMax < height[left] {
//				leftMax = height[left]
//			} else {
//				res += leftMax - height[left]
//			}
//			left++
//		} else {
//			if rightMax < height[right] {
//				rightMax = height[right]
//			} else {
//				res += rightMax - height[right]
//			}
//			right--
//		}
//	}
//	return res
//}
//func multiply(num1 string, num2 string) string {
//	if num1 == "0" || num2 == "0" || len(num1) == 0 || len(num2) == 0 {
//		return "0"
//	}
//	if num1 == "1" {
//		return num2
//	}
//	if num2 == "1" {
//		return num1
//	}
//	var short, long []byte
//	if len(num1) < len(num2) {
//		short = []byte(num1)
//		long = []byte(num2)
//	} else {
//		short = []byte(num2)
//		long = []byte(num1)
//	}
//	for i := 0; i < len(short)/2; i++ {
//		short[i], short[len(short)-1-i] = short[len(short)-1-i], short[i]
//	}
//	for i := 0; i < len(long)/2; i++ {
//		long[i], long[len(long)-1-i] = long[len(long)-1-i], long[i]
//	}
//	var res []byte
//	var temp []byte
//	var tmp int
//	var carry int
//	for i := 0; i < len(short); i++ {
//		for k := 0; k < i; k++ {
//			temp = append(temp, '0')
//		}
//		for j := 0; j < len(long); j++ {
//			tmp = int(long[j]-'0')*int(short[i]-'0') + carry
//			carry = tmp / 10
//			tmp %= 10
//			temp = append(temp, byte(tmp+'0'))
//		}
//		if carry != 0 {
//			temp = append(temp, byte(carry+'0'))
//		}
//		// res+temp,
//		carry = 0
//		for k := i; k < len(temp); k++ {
//			if k < len(res) {
//				tmp = int(res[k]-'0') + int(temp[k]-'0') + carry
//				carry = tmp / 10
//				tmp %= 10
//				res[k] = byte(tmp + '0')
//			} else {
//				tmp = int(temp[k]-'0') + carry
//				if tmp >= 10 {
//					carry = 1
//					tmp -= 10
//				} else {
//					carry = 0
//				}
//				res = append(res, byte(tmp+'0'))
//			}
//		}
//		if carry != 0 {
//			res = append(res, byte(carry+'0'))
//		}
//		carry = 0
//		//清空temp
//		temp = []byte{}
//	}
//	for i := 0; i < len(res)/2; i++ {
//		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
//	}
//	return string(res)
//}
//func isMatch(s string, p string) bool {
//	//if (len(s) == 0 && len(p) == 0) || p=="*" {// 可以没有，但是加上可以加快速度
//	//	return true
//	//}
//	dp := make([][]bool, len(p)+1) // dp[i][j]表示s[:i]和p[:j]是否匹配
//	for i := 0; i <= len(p); i++ {
//		dp[i] = make([]bool, len(s)+1)
//	}
//	dp[0][0] = true // 空串匹配空串
//	//为了方便调试dp[i][j]表示p[:i]和s[:j]是否匹配
//	for i := 1; i <= len(p) && p[i-1] == '*'; i++ { // 全*和空串是匹配的
//		dp[i][0] = true
//	}
//	for i := 1; i <= len(p); i++ {
//		for j := 1; j <= len(s); j++ {
//			if p[i-1] == '*' {
//				for k := 0; k <= j; k++ {
//					if dp[i-1][k] == true {
//						dp[i][j] = true
//						break
//					}
//				}
//				//如果匹配上，应该是可以剪枝的
//				if dp[i][j] == true {
//					for j++; j <= len(s); j++ {
//						dp[i][j] = true
//					}
//				}
//
//			} else if p[i-1] == '?' || p[i-1] == s[j-1] { //不是*，则是否成功匹配，取决于之前是否匹配
//				dp[i][j] = dp[i-1][j-1]
//			}
//		}
//	}
//	return dp[len(p)][len(s)]
//	//for j := 1; j <= len(p) && p[j-1] == '*'; j++ {
//	//	dp[0][j] = true
//	//}
//	//for j := 1; j <= len(p); j++ { // j=0时只考虑s=""是匹配的，其余都认为是false，所以从j=1开始，竖向填充
//	//	for i := 1; i <= len(s); i++ {
//	//		if p[j-1] == '*' {
//	//			for k := 0; k <= i; k++ { //*可以表示空串也可以不表示空串，当k=i时*表示空串，其余表示非空串
//	//				if dp[k][j-1] == true {
//	//					dp[i][j] = true
//	//					break
//	//				}
//	//			}
//	//			if dp[i][j] == true { // 如果这里匹配成功了，*表示后面的都能匹配成功，全部都设为true
//	//				for k := i + 1; k <= len(s); k++ {
//	//					dp[k][j] = true
//	//				}
//	//			}
//	//			break
//	//		} else if p[j-1] == '?' || p[j-1] == s[i-1] { //如果不是*号，则看是否能匹配上，如果能匹配上，则取决于除了这次比较之前的能否匹配上
//	//			dp[i][j] = dp[i-1][j-1]
//	//		} else {
//	//			dp[i][j] = false //最后一个必然没有匹配上
//	//		}
//	//	}
//	//}
//	//return dp[len(s)][len(p)]
//}
//func jump(nums []int) int {
//	// 动态规划
//	//if len(nums)==0 || len(nums) ==1 {
//	//	return 0
//	//}
//	//if nums[0]>=len(nums)-1{
//	//	return 1
//	//}
//	//minStep:= make([]int,len(nums))
//	//minStep[len(nums)-1]=0
//	//for i:=0;i<len(nums)-1;i++{
//	//	minStep[i]=len(nums)+1
//	//}
//	//for i:=len(nums)-2;i>=0;i--{
//	//	for j:=1;j<=nums[i];j++{
//	//		if i+j<len(nums) && minStep[i+j]+1<minStep[i]{
//	//			minStep[i]=minStep[i+j]+1
//	//		}
//	//	}
//	//}
//
//	// 贪心
//	maxPosition := 0
//	step := 0
//	end := 0
//	for i := 0; i < len(nums)-1; i++ {
//		if i+nums[i] >= maxPosition {
//			maxPosition = i + nums[i]
//		}
//		if i == end {
//			end = maxPosition
//			step++
//		}
//		if maxPosition > len(nums)-1 {
//			break
//		}
//	}
//
//	return step
//}
//func permute(nums []int) [][]int {
//	if len(nums) == 1 {
//		return [][]int{nums}
//	}
//	res := *new([][]int)
//	var t []int
//	for i := 0; i < len(nums); i++ {
//		t = append(t, nums[i])
//	}
//	for len(res) == 0 || !Equal(nums, t) {
//		t2 := make([]int, len(t))
//		t = nextPermutation(t)
//		copy(t2, t)
//		res = append(res, t2)
//	}
//	return res
//}
//func Equal(a, b []int) bool {
//	if len(a) != len(b) {
//		return false
//	}
//	for i := 0; i < len(a); i++ {
//		if a[i] != b[i] {
//			return false
//		}
//	}
//	return true
//}
//func nextPermutationUnique(nums []int) []int {
//	var i, j int
//	j = len(nums) - 1
//	for j > 0 && nums[j] <= nums[j-1] { // 找到第一个递减的
//		j--
//	}
//	i = j
//	for i < len(nums)-1 && nums[i+1] == nums[i] { //找到将被交换的位置
//		i++
//	}
//	if j == 0 {
//		for n := 0; n <= (len(nums)-1)/2; n++ {
//			nums[n], nums[len(nums)-n-1] = nums[len(nums)-n-1], nums[n]
//		}
//		return nums
//	}
//	j--
//	k := len(nums) - 1
//	for k >= i && nums[k] <= nums[j] { // 找到第一个比nums[j]大的
//		k--
//	}
//	nums[j], nums[k] = nums[k], nums[j] //交换
//	for n := j + 1; n <= (len(nums)+j)/2; n++ {
//		nums[n], nums[len(nums)+j-n] = nums[len(nums)+j-n], nums[n]
//	}
//	return nums
//}
//func permuteUnique(nums []int) [][]int {
//	if len(nums) == 1 {
//		return [][]int{nums}
//	}
//	res := *new([][]int)
//	var t []int
//	for i := 0; i < len(nums); i++ {
//		t = append(t, nums[i])
//	}
//	for len(res) == 0 || !Equal(nums, t) {
//		t2 := make([]int, len(t))
//		t = nextPermutationUnique(t)
//		copy(t2, t)
//		res = append(res, t2)
//	}
//	return res
//}
//func rotate1(matrix [][]int) {
//	if len(matrix) == 0 {
//		return
//	}
//	for i := 0; i < len(matrix)/2; i++ {
//		for j := i; j < len(matrix[i])-i-1; j++ {
//			matrix[i][j], matrix[j][len(matrix[i])-i-1] = matrix[j][len(matrix[i])-i-1], matrix[i][j]
//			matrix[i][j], matrix[len(matrix[i])-j-1][i] = matrix[len(matrix[i])-j-1][i], matrix[i][j]
//			matrix[len(matrix[i])-j-1][i], matrix[len(matrix[i])-i-1][len(matrix[i])-j-1] = matrix[len(matrix[i])-i-1][len(matrix[i])-j-1], matrix[len(matrix[i])-j-1][i]
//		}
//	}
//	return
//
//}
//func groupAnagrams(strs []string) [][]string {
//	if len(strs) == 0 || len(strs) == 1 {
//		return [][]string{strs}
//	}
//	var res [][]string
//	var patterns []map[byte]bool
//	pattern := make(map[byte]bool)
//	for i := 0; i < len(strs[0]); i++ {
//		pattern[strs[0][i]] = true
//	}
//	res = append(res, []string{strs[0]})
//	patterns = append(patterns, pattern)
//	for i := 1; i < len(strs); i++ { // 匹配第i个string
//		flag := false                        //默认匹配不上
//		for j := 0; j < len(patterns); j++ { //第j个模式
//			if len(strs[i]) != len(patterns[j]) {
//				break
//			}
//			k := 0
//			for ; k < len(strs[i]); k++ {
//				if _, ok := patterns[j][strs[i][k]]; !ok {
//					break
//				}
//			}
//			if k == len(strs[i]) { // 和第j个模式匹配上了
//				res[j] = append(res[j], strs[i])
//				flag = true
//				break
//			}
//		}
//		if !flag { //不适合于任何模式，自己建立一个
//			pattern := make(map[byte]bool)
//			for j := 0; j < len(strs[i]); j++ {
//				pattern[strs[i][j]] = true
//			}
//			res = append(res, []string{strs[i]})
//			patterns = append(patterns, pattern)
//		}
//	}
//	return res
//}
func main() {
	//var b map[byte]interface{}
	//fmt.Print(findMedianSortedArrays([]int{1,2, 3}, []int{5}))
	//fmt.Printf("%s\n", longestPalindrome("bbbcccc"))
	//fmt.Printf("%s", convert("PAYPALISHIRING", 4))
	//fmt.Printf("%d",reverse(120))
	//fmt.Printf("%d\n", myAtoi(  "-2147483647"))
	//fmt.Printf("%v\n", isPalindrome(121))
	//fmt.Printf("%v\n", isMatch("aa","a"))
	//fmt.Printf("%v\n", maxArea([]int{1,1}))
	//fmt.Printf("%s\n", intToRoman(1994))
	//fmt.Printf("%d\n", romanToInt("LVIII"))
	//fmt.Printf("%s\n", longestCommonPrefix([]string{"dog","racecar","car"}))
	//fmt.Printf("%#v\n", threeSum1([]int{0,0,0,0}))
	//fmt.Printf("%#v\n", threeSumClosest([]int{-1,2,1,-4},1))
	//fmt.Printf("%#v\n", letterCombinations(""))
	//fmt.Printf("%#v\n", fourSum([]int{0, 0, 0, 0, 0, 0}, 0))
	//fmt.Printf("%#v\n", isValid("[[[]"))
	//res := mergeTwoLists(&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}, &ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//})
	//for res != nil {
	//	fmt.Printf("%d->", res.Val)
	//	res = res.Next
	//}
	//fmt.Printf("%v\n", generateParenthesis(3))
	//res:=mergeKLists([]*ListNode{&ListNode{
	//	Val:  1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: &ListNode{
	//			Val:  2,
	//			Next: nil,
	//		},
	//	},
	//},&ListNode{
	//	Val:  1,
	//	Next: &ListNode{
	//		Val:  3,
	//		Next: &ListNode{
	//			Val:  7,
	//			Next: nil,
	//		},
	//	},
	//},&ListNode{
	//	Val:  3,
	//	Next: &ListNode{
	//		Val:  5,
	//		Next: nil,
	//	},
	//},&ListNode{
	//	Val:  4,
	//	Next: &ListNode{
	//		Val:  9,
	//		Next: nil,
	//	},
	//},&ListNode{
	//	Val:  9,
	//	Next: nil,
	//}})
	//for res != nil {
	//	fmt.Printf("%d->", res.Val)
	//	res = res.Next
	//}
	//res := swapPairs(nil)
	//for res != nil {
	//	fmt.Printf("%d->", res.Val)
	//	res=res.Next
	//}
	//res := reverseKGroup(&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: nil,
	//	}},2)
	//for res != nil {
	//	fmt.Printf("%d->", res.Val)
	//	res = res.Next
	//}
	//fmt.Printf("%d\n",removeDuplicates([]int{1}))
	//fmt.Printf("%d\n",removeElement([]int{1,2,3,7,4,5,5,5,8},5))
	//fmt.Printf("%d\n",strStr("hello","heo"))
	//fmt.Printf("%d\n", divide(2147483648, -1))
	//fmt.Printf("%v\n", findSubstring1("ababaab",[]string{"ab","ba","ba"}))
	//n:=[]int{3,2,1}
	//nextPermutation(n)
	//for i:=0;i<len(n);i++{
	//	fmt.Printf("%d",n[i])
	//}
	//fmt.Printf("%v\n", longestValidParentheses("((()))((())()"))
	//fmt.Printf("%v\n", isMatch("aa",".*"))
	//fmt.Printf("%v\n", search([]int{4,5,1,2,3},3))
	//fmt.Printf("%v\n", searchRange([]int{2,2,2,2}, 2))
	//fmt.Printf("%v\n", searchInsert([]int{2, 3, 5}, 6))
	//fmt.Printf("%v\n", isValidSudoku([][]byte{
	//	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	//	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	//	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	//	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	//	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	//	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	//	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	//	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	//	{0, 0, 0, 0, 8, 0, 0, 7, 9},
	//}))
	//TODO: finish this function
	//t:=[][]byte{
	//	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	//	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	//	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	//	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	//	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	//	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	//	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	//	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	//	{0, 0, 0, 0, 8, 0, 0, 7, 9},
	//}
	//solveSudoku(t)
	//for i:=0;i<9;i++{
	//	for j:=0;j<9;j++{
	//		fmt.Printf("%d\t",t[i][j])
	//	}
	//	fmt.Println()
	//}
	//fmt.Printf("%s\n",countAndSay(5))
	//res := combinationSum2([]int{2, 3, 4, 14, 16}, 18)
	//for i := 0; i < len(res); i++ {
	//	for j := 0; j < len(res[i]); j++ {
	//		fmt.Printf("%d,", res[i][j])
	//	}
	//	fmt.Println()
	//}
	//fmt.Printf("%d",firstMissingPositive([]int{2,2,6,1}))
	//fmt.Printf("%d", trap([]int{4,3,4}))
	//fmt.Printf("%v", multiply("98","9"))
	//fmt.Printf("%v\n", isMatch("aa", "*"))
	//fmt.Printf("%v\n", jump([]int{5,1,1,1,1,1}))
	//fmt.Printf("%v\n", permute([]int{0,1,2}))
	//fmt.Printf("%v\n", permuteUnique([]int{1,2,2,3,3}))
	//in := [][]int{
	//	{5, 1, 9, 11},
	//	{2, 4, 8, 10},
	//	{13, 3, 6, 7},
	//	{15, 14, 12, 16},
	//}
	//rotate(in)
	//for i := 0; i < len(in); i++ {
	//	for j := 0; j < len(in[i]); j++ {
	//		fmt.Printf("%d\t", in[i][j])
	//	}
	//	fmt.Println()
	//}
	//fmt.Printf("%v\n", groupAnagras([]string{"acc","ac","c"}))
	//fmt.Printf("%v",pow(2,-3))
	//res:=solveNQueens(3)
	//for i:=0;i<len(res);i++{
	//	for j:=0;j<len(res[i]);j++{
	//		fmt.Println(res[i][j])
	//	}
	//	fmt.Println()
	//}
	//fmt.Printf("%d",maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
	//res:=spiralOrder([][]int{{1,11},{2,12},{3,13},{4,14},{5,15},{6,16},{7,17},{8,18},{9,19},{10,20}})
	//fmt.Printf("%v",res)
	//fmt.Printf("%v",canJump([]int{1,0,1,0,4}))
	//fmt.Printf("%v",merge([][]int{{4,5},{1,4},{2,6},{7,9},{11,22}}))
	//fmt.Printf("%v",insert([][]int{{23,25},{30,34},{41,43},{45,49}},[]int{29,32}))
	//fmt.Printf("%v",lengthOfLastWord("    ffff   "))
	//fmt.Printf("%v",generateMatrix(4))
	//fmt.Printf("%v",getPermutation(3,5))
	//res:=&ListNode{
	//	Val:  1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: &ListNode{
	//				Val:  4,
	//				Next: &ListNode{
	//					Val:  5,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}
	//res=rotateRight(res,5)
	//for res!=nil{
	//	fmt.Printf("%d->",res.Val)
	//	res=res.Next
	//}
	//fmt.Printf("%v",uniquePaths(7,3))
	//fmt.Printf("%v",uniquePathsWithObstacles([][]int{{0,1}}))
	//fmt.Printf("%v",minPathSum([][]int{{1,3,1},{1,5,1},{4,2,1}}))
	//fmt.Printf("%v",isNumber("   2e-10 "))
	//fmt.Printf("%v",plusOne([]int{8}))
	//fmt.Printf("%v",addBinary("1010","1011"))
	//fmt.Printf("%v", fullJustify([]string{
	//	"ask","not","what","your","country","can","do","for","you","ask","what","you","can","do","for","your","country",
	//},16))
	//fmt.Printf("%v\n", computeValue(
	//	[]int{500,200,100,50,20,10,5,2,1},
	//	[]int{10,10,10,10,10,10,10,10,10},
	//	670, 10))
	//fmt.Printf("%v\n",mySqrt(17))
	//fmt.Printf("%v\n",climbStairs(4))
	//fmt.Printf("%v\n",simplifyPath("/a/a/b/c/./../"))
	//fmt.Printf("%v\n",minDistance("intention","execution"))
	//m:=[][]int{{1,2,3,4},{5,0,7,8},{0,10,11,12},{13,14,15,0}}
	//setZeroes(m)
	//fmt.Printf("%v\n",m)
	//fmt.Printf("%v\n",searchMatrix([][]int{{2}},1))
	//nums := []int{1,2,0,0,1,0,1,0,2,2,0,2,1,0,1,2,0,1,2,1,0,1,2,1,0}
	//sortColors(nums)
	//fmt.Printf("%s",minWindow("ADOBECODEBANCDEAB","ADD"))
	//fmt.Printf("%s",minWindow1("ADOBECODEBANC","ABC"))
	//fmt.Printf("%v",combine1(5,3))
	//fmt.Printf("%v",subsets([]int{1,2,3,4,5}))
	//fmt.Printf("%v", exist([][]byte{
	//	{'A','B','C','E'},
	//	{'S','F','C','S'},
	//	{'A','D','E','E'},
	//}, "ABCCED"))
	//fmt.Printf("%v",removeDuplicatesII([]int{0,0,0,0,1,1,1,1,2,2,2,3,3}))
	//fmt.Printf("%v",searchII([]int{1,3,1,1,1},1))
	//leetcode:=&ListNode{
	//	Val:  1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: &ListNode{
	//			Val:  2,
	//			Next: &ListNode{
	//				Val:  3,
	//				Next: &ListNode{
	//					Val: 4,
	//					Next: &ListNode{
	//						Val:  4,
	//						Next: &ListNode{
	//							Val:  5,
	//							Next: nil,
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//res:=deleteDuplicates(leetcode)
	//p:=res
	//for p!=nil{
	//	fmt.Printf("%d->",p.Val)
	//	p=p.Next
	//}
	//fmt.Printf("%d",largestRectangleArea([]int{2,1,5,6,2,3}))
	//fmt.Printf("%d", maximalRectangle([][]byte{
	//	//{'1', '0', '1', '0', '0'},
	//	//{'1', '0', '1', '1', '1'},
	//	//{'1', '1', '1', '1', '1'},
	//	//{'1', '0', '0', '1', '0'},
	//	{'0','1'},
	//	{'0','1'},
	//}))
	//_=&ListNode{
	//	Val:  1,
	//	Next: &ListNode{
	//		Val:  4,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: &ListNode{
	//				Val:  2,
	//				Next: &ListNode{
	//					Val:  5,
	//					Next: &ListNode{
	//						Val:  2,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//res:=partition(&ListNode{
	//	Val:  1,
	//	Next: &ListNode{
	//		Val:  1,
	//		Next: nil,
	//	},
	//},0)
	//p:=res
	//for p!=nil{
	//	fmt.Printf("%d->",p.Val)
	//	p=p.Next
	//}

	//head:=&ListNode{
	//	Val:  0,
	//	Next: &ListNode{
	//		Val:  1,
	//		Next: &ListNode{
	//			Val:  2,
	//			Next: &ListNode{
	//				Val:  3,
	//				Next: &ListNode{
	//					Val:  4,
	//					Next: &ListNode{
	//						Val:  5,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//rear:=head
	//for rear.Next!=nil{
	//	rear=rear.Next
	//}
	////rear.Next=head.Next.Next
	//fmt.Printf("%t",hasCycle(head))

	//fmt.Printf("%t",canPartition([]int{100,100,100,100,100,100,100,100}))

	//fmt.Printf("%t",isScramble("abcde","caebd"))

	//nums1:=[]int{1,2,3,0,0,0}
	//nums2:=[]int{4,5,6}
	//mergeIn88(nums1,len(nums1)-len(nums2),nums2,len(nums2))

	//grayCode(1)

	//subsetsWithDup([]int{1,2,2})

	//fmt.Println(numDecodings("0"))

	//head:=&ListNode{
	//	Val:  1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: &ListNode{
	//				Val:  4,
	//				Next: &ListNode{
	//					Val:  5,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}
	//res:=reverseBetween(head,1,5)
	//p:=res
	//for p!=nil{
	//	fmt.Printf("%d->",p.Val)
	//	p=p.Next
	//}

	//fmt.Printf("%v",restoreIpAddresses("101023"))

	//generateTrees(3)
	//fmt.Printf("%t",isInterleave("ab","aa","abaa"))

	//fmt.Printf("%v",generate(5))
	//fmt.Printf("%v",getRow(5))
	//fmt.Printf("%v",maxProfit([]int{1,2,4,5,3,7}))
	//fmt.Printf("%v",isPalindrome1("0P"))
	//fmt.Printf("%v",findLadders("hit", "cog", []string{"hot","dot","dog","lot","log","cog"}))
	//fmt.Println(longestConsecutive([]int{}))
	//fmt.Println(sumNumbers(&TreeNode{
	//	Val:   0,
	//	Left:  &TreeNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: nil,
	//}))
	//solve([][]byte{{'X', 'O', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}})
	//partition("baaab")
	//minCut("ababab")
	//fmt.Println(candy([]int{1,4,2,5,7,1,7,8,3,2}))
	//fmt.Println(wordBreak("leetcode",[]string{"leet","code"}))
	//fmt.Println(maxPoints([][]int{{1,1},{2,2},{3,3}}))
	//fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	//fmt.Println(reverseWords("   Alice does not even like bob"))
	//fmt.Println(maxProduct1([]int{0,-2}))
	//fmt.Println(findMin([]int{1,2,3,4,1,1,1,1}))
	//fmt.Println(findPeakElement([]int{2,1}))
	//fmt.Println(maximumGap([]int{3,6,2,9,1}))
	//fmt.Println(compareVersion("0.0.0","0"))
	//fmt.Println(fractionToDecimal(-60,-5))
	//fmt.Println(convertToTitle(44))
	//fmt.Println(majorityElement([]int{3,2,2,3}))
	//fmt.Println(titleToNumber("AA"))
	//fmt.Println(trailingZeroes1(25))
	//a:=[]int{0,2,11,14,17,21,25}
	//start,end:=0,a[len(a)-1]+1
	//for start+1<end{
	//	i:=0
	//	del:=0
	//	mid:=(start+end)/2
	//	for i<=len(a)-2{
	//		j:=i+1
	//		for j<len(a) && a[j]-a[i]<mid{
	//			j++
	//		}
	//		del+=j-i-1
	//		i=j
	//	}
	//	if del<=2{
	//		start=mid
	//	}else{
	//		end=mid
	//	}
	//	fmt.Println(start)
	//}
	//fmt.Println(start)
	//fmt.Println(calculateMinimumHP([][]int{
	//	{-2,-3,3},
	//	{-5,-10,1},
	//	{10,30,-5},
	//}))
	//fmt.Printf(largestNumber([]int{3,30,34,5,9}))
	//rotate([]int{1,2,3,4},2)
	//fmt.Println(reverseBits(1))
	//fmt.Println(hammingWeight(15))
	//fmt.Println(rangeBitwiseAnd(5,7))
	//fmt.Println(isHappy(2))
	//fmt.Println(countPrimes(10))
	//fmt.Println(isIsomorphic("title","paper"))
	//fmt.Println(canFinish(2,[][]int{{0,1}}))
	//t:=TrieConstructor()
	//t.Insert("apple")
	//fmt.Println(t.Search("app"))
	//fmt.Println(t.Search("apple"))
	//fmt.Println(t.StartsWith("apple"))
	//fmt.Println(t.StartsWith("app"))
	//t.Insert("app")
	//fmt.Println(t.Search("app"))
	//fmt.Println(t.StartsWith("app"))
	//fmt.Println(minSubArrayLen(16, []int{2, 3, 1, 2, 2, 3}))
	//fmt.Println(findOrder(4,[][]int{{1,0},{2,0},{3,1},{3,2}}))
	//fmt.Println(findWords(
	//	[][]byte{
	//		{'o', 'a', 'a', 'n'},
	//		{'e', 't', 'a', 'e'},
	//		{'i', 'h', 'k', 'r'},
	//		{'i', 'f', 'l', 'v'},
	//	},
	//	[]string{"oath", "pea", "eat", "rain"}))
	//fmt.Println(rob([]int{2, 5, 1, 3, 4, 3, 2, 6, 1, 8, 9}))
	//fmt.Println(question.ShortestBridge([][]int{
	//	{1,1,1,1,1},{1,0,0,0,1},{1,0,1,0,1},{1,0,0,0,1},{1,1,1,1,1},
	//}))
	//inputs := [][][]int{
	//	[][]int{
	//		{1,2,7},
	//		{3,6,7},
	//	},
	//	[][]int{
	//		{7,12},
	//		{4,5,15},
	//		{6},
	//		{15,19},
	//		{9,12,13},
	//	},
	//	[][]int{
	//		{24},
	//		{3, 6, 11, 14, 22},
	//		{1, 23, 24},
	//		{0, 6, 14},
	//		{1, 3, 8, 11, 20},
	//	},
	//	[][]int{
	//		{1, 9, 12, 20, 23, 24, 35, 38},
	//		{10, 21, 24, 31, 32, 34, 37, 38, 43},
	//		{10, 19, 28, 37},
	//		{8},
	//		{14, 19},
	//		{11, 17, 23, 31, 41, 43, 44},
	//		{21, 26, 29, 33},
	//		{5, 11, 33, 41},
	//		{4, 5, 8, 9, 24, 44},
	//	},
	//	[][]int{
	//		{3, 16, 33, 45, 59, 79, 103, 135},
	//		{3, 35, 39, 54, 56, 78, 96, 101, 120, 132, 146, 148},
	//		{13, 72, 98},
	//		{37, 70, 107},
	//		{0, 12, 31, 37, 41, 68, 78, 94, 100, 101, 113, 123},
	//		{11, 32, 52, 85, 135},
	//		{43, 50, 128},
	//		{0, 13, 49, 51, 53, 55, 60, 65, 66, 80, 82, 87, 92, 99, 112, 118, 120, 125, 128, 131, 137},
	//		{15, 19, 34, 37, 45, 52, 56, 97, 108, 123, 142},
	//		{7, 9, 20, 28, 29, 33, 34, 38, 43, 46, 47, 48, 53, 59, 65, 72, 74, 80, 88, 92, 110, 111, 113, 119, 135, 140},
	//		{15, 41, 64, 83},
	//		{7, 13, 26, 31, 57, 85, 101, 108, 110, 115, 119, 124, 149},
	//		{47, 61, 67, 70, 74, 75, 77, 84, 92, 101, 124, 132, 133, 142, 147},
	//		{0, 2, 5, 6, 12, 18, 34, 37, 47, 58, 77, 98, 99, 109, 112, 131, 135, 149},
	//		{6, 7, 8, 9, 14, 17, 21, 25, 33, 40, 45, 50, 56, 57, 58, 60, 68, 92, 93, 100, 108, 114, 130, 149},
	//		{7},
	//		{5, 16, 22, 48, 77, 82, 108, 114, 124},
	//		{34, 71},
	//		{8, 16, 32, 48, 104, 108, 116, 134, 145},
	//		{3, 10, 16, 19, 35, 45, 64, 74, 89, 101, 116, 149},
	//		{1, 5, 7, 10, 11, 18, 40, 45, 50, 51, 52, 54, 55, 69, 71, 81, 82, 83, 85, 89, 96, 100, 114, 115, 124, 134, 138, 148},
	//		{0, 2, 3, 5, 6, 9, 15, 52, 64, 103, 108, 114, 146},
	//		{5, 33, 39, 40, 44, 45, 66, 67, 68, 69, 84, 102, 106, 115, 120, 128, 133},
	//		{17, 26, 49, 50, 55, 58, 60, 65, 88, 90, 102, 121, 126, 130, 137, 139, 144},
	//		{6, 12, 13, 37, 41, 42, 48, 50, 51, 55, 64, 65, 68, 70, 73, 102, 106, 108, 120, 123, 126, 127, 129, 135, 136, 149},
	//		{6, 7, 12, 33, 37, 41, 47, 53, 54, 80, 107, 121, 126},
	//		{15, 75, 91, 103, 107, 110, 125, 139, 142, 149},
	//		{18, 24, 30, 52, 61, 64, 75, 79, 85, 95, 100, 103, 105, 111, 128, 129, 142},
	//		{3, 14, 18, 32, 45, 52, 57, 63, 68, 78, 85, 91, 100, 104, 111, 114, 142},
	//		{4, 7, 11, 20, 21, 31, 32, 33, 48, 61, 62, 65, 66, 73, 80, 92, 93, 97, 99, 108, 112, 116, 136, 139},
	//	},
	//	[][]int{
	//	{148,167,216},
	//	{6,23,25,40,43,58,63,69,77,86,94,96,106,117,119,127,139,151,153,155,157,186,191,196,200,204,210,216,219},
	//	{2,6,7,16,27,30,42,47,49,68,69,77,93,94,96,102,104,111,114,126,131,137,150,161,167,171,174,193,198,199,200,223},
	//	{46,131,211},
	//	{25,36,51,52,65,78,90,102,103,105,108,114,123,151,152,153,162,174,175},
	//	{217},
	//	{9,10,15,27,37,38,41,43,46,51,67,74,81,82,83,94,95,107,113,120,122,123,124,132,149,160,162,169,170,171,174,177,185,192,193,195,196,198,213,217,220,221},
	//	{74,78,85,95,130,136,145,152,173,175,180,181,184,193,199,202},
	//	{13,18,28,38,41,42,47,75,87,91,106,151,158,166,181,182,199,216},
	//	{44,63,71,74,144,162,169,220},
	//	{2,23,115,185,208},
	//	{0,8,13,14,35,46,67,89,91,122,124,126,130,156,177,193,212,214},
	//	{2,4,24,37,40,43,55,68,81,92,106,107,109,127,132,138,145,159,163,165,170,172,183,184,209,213,215,220},
	//	{5,16,17,34,38,48,55,59,60,65,69,84,86,94,100,103,109,110,112,127,130,131,134,145,148,149,154,161,166,169,182,183,201,203,208,214,223},
	//	{0,2,5,6,8,19,49,50,53,79,92,94,97,109,110,112,121,129,132,135,138,139,144,160,166,170,194,197,198,201,212},
	//	{27,52,61,112,118,133,142,159,175,186,216},
	//	{2,20,34,64,65,77,87,91,95,96,97,125,126,131,144,146,149,152,154,164,165,170,179,205,207},
	//	{24,85,123,132,172,173,194,222},
	//	{2,4,5,15,23,36,44,47,63,64,78,80,84,97,99,102,104,114,120,130,132,143,161,162,163,167,171,172,176,179,180,194,196,199,202,204,209,214,216,221},
	//	{8,22,26,31,38,39,41,59,78,90,102,108,110,138,141,146,176,185,190,198,200,219,220},
	//	{5,24,30,46,55,64,67,74,78,136,194,216},
	//	{133,142,202},
	//	{13,40,49,57,63,75,76,85,91,107,116,121,128,135,137,141,154,193,198,200,204,223},
	//	{4,13,14,26,28,33,39,49,58,65,67,74,77,81,90,96,122,124,144,156,158,166,169,170,179,203,204,208,215,223},
	//	{6,20,28,36,46,90,107,115,124,131,135,144,147,148,149,161,162,174,176,214,221},
	//	{10,20,21,29,35,36,62,65,67,70,72,87,89,92,100,103,107,109,113,126,129,139,140,145,146,147,174,176,180,184,189,190,193,196,198,199,200,209,217},
	//	{19,22,27,54,59,63,77,102,122,126,140,143,154,164,165,175,212,216,217,218},
	//	{11,13,16,18,27,31,46,49,69,77,88,109,111,119,121,146,161,169,193,194,198,200,204},
	//	{1,7,28,58,73,91,98,138,150,173,182,186,213},
	//	{3,25,28,33,46,68,70,74,78,97,141,146,149,169,172,178,185,188,202,212,223},
	//	{3,4,19,22,24,37,38,43,54,55,56,57,58,62,66,72,75,77,88,106,114,119,127,132,133,137,144,146,150,156,161,164,165,179,181,195,200,213,214,215,222},
	//	{9,11,14,15,38,46,55,61,66,68,69,75,76,79,82,91,100,101,102,113,135,141,142,171,175,180,198,208,210,215,218,221},
	//	{2,30,33,62,93,104,124,127,128,147,158,160,161,173,181,189,192,199,201,215,223},
	//	{4,26,29,38,47,58,61,69,78,93,94,112,114,131,136,144,182,193,198,203,206,209},
	//	{5,13,14,16,17,22,30,32,45,47,49,55,63,64,68,77,82,84,86,92,98,100,104,107,117,119,122,127,134,153,164,179,185,197,201,209,212,213,220,223},
	//	{2,4,5,6,42,55,75,81,84,93,102,111,112,113,118,129,142,149,159,169,191,193,200,214,223},
	//	{10,12,15,19,20,24,33,34,40,47,54,64,93,104,115,121,123,124,155,172,189,190,193,196,202,212,219,222},
	//	{104,108,143},
	//	{14,15,20,21,31,47,48,59,67,70,74,82,94,102,109,121,125,128,148,162,165,171,180,196,199,202,205,212,214},
	//	{2,6,17,18,41,50,60,70,118,151,155,158,166,167,172,180,182,186,188,195},
	//	{1,23,25,30,39,41,42,48,58,65,67,94,100,121,126,135,145,152,163,164,171,174,206,210,220,224},
	//	{18,25,96,123,172},
	//	{5,7,9,12,13,19,22,25,34,51,62,64,74,79,81,85,88,101,102,119,123,140,143,149,155,165,166,167,178,182,189,204,213,222,223},
	//	{1,5,18,21,23,50,54,59,62,67,68,72,87,94,95,96,110,116,118,122,133,135,151,155,156,158,171,178,183,184,192,198,208,212,222,224},
	//	{18,20,24,34,47,52,56,68,77,82,89,91,97,101,105,106,107,109,118,123,139,141,143,152,153,162,174,180,184,187,188,192,198,202,206,216,224},
	//	},
	//}
	//sources:=[]int{6,15,20,37,85,180}
	//targets:=[]int{2,12,8,28,112,143}
	//for i, input := range inputs {
	//	fmt.Println(question.NumBusesToDestination815_2(
	//		input,
	//		sources[i],
	//		targets[i],
	//	))
	//	fmt.Println(question.NumBusesToDestination815_1(
	//		input,
	//		sources[i],
	//		targets[i],
	//	))
	//}
	//input:=[]string{
	//	"",
	//	"x+5-3+x=6+x-2",
	//	"x=x",
	//	"2x=x",
	//	"2x+3x-6x=x+2",
	//	"x=x+2",
	//	"-12+21x-3-5x=2x-9",
	//}
	//for _, s := range input {
	//	fmt.Println(question.SolveEquation(s))
	//}
	//fmt.Println(question.IsPossible([]int{1, 1, 1, 2}))
	//newmodule.Ouput()
	//remain := 7
	//flipped := []int{5, 5}
	//loc := 6
	//for i := len(flipped) - 1; i >= 0; i-- {
	//	loc = loc%remain + flipped[i] + 1
	//	remain++
	//}
	//loc = loc % remain
	//fmt.Println(loc)
	//fmt.Println(classify([]string{"cat","act","ns","sn","and","tac","ttac","dna"}))
	//fmt.Println(question.BrokenCalc(1,1000))
	//fmt.Println(question.AllCellsDistOrder(100,100,51,42))
	//fmt.Println(question.MaxSumTwoNoOverlap([]int{2,1,5,6,0,9,5,0,3,8},3,4))
	//fmt.Println(question.MaxRepOpt1("aaaaa"))
	//fmt.Println(question.KthLargestValue([][]int{{5,2},{1,6}},1))
	//fmt.Println(question.KthLargestValue([][]int{{5,2},{1,6}},2))
	//fmt.Println(question.KthLargestValue([][]int{{5,2},{1,6}},3))
	//fmt.Println(question.KthLargestValue([][]int{{5,2},{1,6}},4))
	//fmt.Println(question.KthLargestValue2([][]int{{5,2},{1,6}},1))
	//fmt.Println(question.KthLargestValue2([][]int{{5,2},{1,6}},2))
	//fmt.Println(question.KthLargestValue2([][]int{{5,2},{1,6}},3))
	//fmt.Println(question.KthLargestValue2([][]int{{5,2},{1,6}},4))
	//fmt.Println(question.LetterCasePermutation("a"))
	//fmt.Println(2^0)
	//fmt.Println(question.IsBipartite3([][]int{{1},{0,3},{3},{1,2}}))
	//fmt.Println(question.IsBipartite3([][]int{{1,2,3},{0,2},{0,1,3},{0,2}}))
	//fmt.Println(question.IsBipartite3([][]int{{1,3},{0,2},{1,3},{0,2}}))
	//fmt.Println(question.MinSumOfLengths([]int{3,2,2,4,3},3))
	//fmt.Println(question.MinSumOfLengths([]int{7,3,4,7},7))
	//fmt.Println(question.MinSumOfLengths([]int{4,3,2,6,2,3,4},6))
	//fmt.Println(question.MinSumOfLengths([]int{5,5,4,4,5},3))
	//fmt.Println(question.MinSumOfLengths([]int{3,1,1,1,5,1,2,1},3))
	//fmt.Println(question.MinSumOfLengths([]int{1,1,1,2,2,2,4,4},6))
	//fmt.Println(question.MinSwaps([][]int{{0,0,1},{1,1,0},{1,0,0}}))
	//fmt.Println(question.MinSwaps([][]int{{0,1,1,0},{0,1,1,0},{0,1,1,0},{0,1,1,0}}))
	//fmt.Println(question.MinSwaps([][]int{{1,0,0},{1,1,0},{1,1,1}}))
	//fmt.Println(question.MinSwaps([][]int{{0}})) //0
	//fmt.Println(question.MinSwaps([][]int{{1,0,0,0,0,0},{0,1,0,1,0,0},{1,0,0,0,0,0},{1,1,1,0,0,0},{1,1,0,1,0,0},{1,0,0,0,0,0}}))
	//fmt.Println(question.MinSwaps([][]int{{1,0,0,0,0,0},{0,0,0,1,0,0},{0,0,0,1,0,0},{0,1,0,0,0,0},{0,0,1,0,0,0},{0,0,0,0,0,1}}))
	//fmt.Println(question.IsSelfCrossing([]int{2,1,1,2}))
	//fmt.Println(question.IsSelfCrossing([]int{1,2,3,4}))
	//fmt.Println(question.IsSelfCrossing([]int{1,1,1,1}))
	//fmt.Println(question.IsSelfCrossing([]int{5,5,10,10}))
	//fmt.Println(question.IsSelfCrossing([]int{5,5,10,10,4}))
	//fmt.Println(question.IsSelfCrossing([]int{5,5,10,10,4,9}))
	//fmt.Println(question.IsSelfCrossing([]int{5,5,10,10,4,9,3}))
	//fmt.Println(question.IsSelfCrossing([]int{5,5,10,10,4,9,3,8}))
	//fmt.Println(question.IsSelfCrossing([]int{5,5,10,10,4,9,3,8,3}))
	//fmt.Println(question.IsSelfCrossing([]int{1,1,2,1,1}))
	//fmt.Println(question.IsSelfCrossing([]int{2,1,4,4,3,2,2,1,1}))
	//fmt.Println(question.RepeatedStringMatch("abcd","cdabcdab"))
	//fmt.Println(question.RepeatedStringMatch("a","aa"))
	//fmt.Println(question.RepeatedStringMatch("a","a"))
	//fmt.Println(question.RepeatedStringMatch("abc","wxyz"))
	//fmt.Println(question.RepeatedStringMatch("abc","abd"))
	//fmt.Println(question.RepeatedStringMatch("abc","cab"))
	//fmt.Println(question.RepeatedStringMatch("abc","ab"))
	//fmt.Println(question.RepeatedStringMatch("abc","abca"))
	//fmt.Println(question.RepeatedStringMatch("abc","bcabc"))
	//fmt.Println(question.RepeatedStringMatch("abaabaa","abaababaab"))
	//fmt.Println(question.RepeatedStringMatch("aabaa","aaab"))
	//fmt.Println()
	//fmt.Println(question.RepeatedStringMatch2("abcd","cdabcdab"))
	//fmt.Println(question.RepeatedStringMatch2("a","aa"))
	//fmt.Println(question.RepeatedStringMatch2("a","a"))
	//fmt.Println(question.RepeatedStringMatch2("abc","wxyz"))
	//fmt.Println(question.RepeatedStringMatch2("abc","abd"))
	//fmt.Println(question.RepeatedStringMatch2("abc","cab"))
	//fmt.Println(question.RepeatedStringMatch2("abc","ab"))
	//fmt.Println(question.RepeatedStringMatch2("abc","abca"))
	//fmt.Println(question.RepeatedStringMatch2("abc","bcabc"))
	//fmt.Println(question.RepeatedStringMatch2("abaabaa","abaababaab"))
	//fmt.Println(question.RepeatedStringMatch2("aabaa","aaab"))
	//fmt.Println(solve([]int{9,9,15,7,10,8,12}))
	//fmt.Println(solve([]int{10,10,9,8,13}))
	//fmt.Prntln(question.FlipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
	//fmt.Println(question.FindMinDifference([]string{"11:59","00:00"}))
	//fmt.Println(question.FindMinDifference([]string{"12:01","00:00"}))
	//fmt.Println(question.FindMinDifference([]string{"02:00","00:00","22:01"}))
	//fmt.Println(question.ShortestPathBinaryMatrix([][]int{
	//	{0,1},
	//	{1,0},
	//}))
	//fmt.Println(question.ShortestPathBinaryMatrix([][]int{
	//	{0,0,0},
	//	{1,1,0},
	//	{1,1,0},
	//}))
	//fmt.Println(question.ShortestPathBinaryMatrix([][]int{
	//	{1,0,0},
	//	{1,1,0},
	//	{1,1,0},
	//}))
	//fmt.Println(question.ShortestPathBinaryMatrix([][]int{
	//	{0,1,0,0,0},
	//	{0,1,0,1,0},
	//	{0,1,0,1,0},
	//	{0,1,0,1,0},
	//	{0,0,0,1,0},
	//}))
	//fmt.Println(question.FallingSquares([][]int{
	//	{1,2},
	//	{2,3},
	//	{6,1},
	//}))
	//fmt.Println(question.FallingSquares([][]int{
	//	{100,100},
	//	{200,100},
	//}))
	//fmt.Println(question.MaxChunksToSorted([]int{0,1,2,3,4}))// 5「0,1,2,3,4」
	//fmt.Println(question.MaxChunksToSorted([]int{4,3,2,1,0}))// 1「4,4,4,4,4」
	//fmt.Println(question.MaxChunksToSorted([]int{1,0,2,3,4}))// 4「1,1,2,3,4」
	//fmt.Println(question.MaxChunksToSorted([]int{0,3,2,1,4}))// 3「0,3,3,3,4」
	//fmt.Println(question.MaxChunksToSorted([]int{0})) // 1「0」
	//fmt.Println(question.MaxChunksToSorted([]int{1,2,0,3})) // 2「2,2,2,3」
	//fmt.Println(question.MaxChunksToSorted([]int{2,0,1})) // 1「1,1,2」
	//fmt.Println(question.CheckPalindromeFormation("x","y"))
	//fmt.Println(question.CheckPalindromeFormation("adbef","fecab"))
	//fmt.Println(question.CheckPalindromeFormation("ulacfd","jizalu"))
	//fmt.Println(question.CheckPalindromeFormation("xbdef","xecab"))
	//fmt.Println(question.IsPossibleDivide([]int{1, 2, 3, 3, 4, 4, 5, 6}, 4))
	//fmt.Println(question.IsPossibleDivide([]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3))
	//fmt.Println(question.IsPossibleDivide([]int{3, 3, 2, 2, 1, 1}, 3))
	//fmt.Println(question.IsPossibleDivide([]int{1, 2, 3, 4}, 3))
	//fmt.Println(question.IsPossibleDivide([]int{1, 2, 3, 4, 3, 6}, 3))
	//fmt.Println(question.IsPossibleDivide([]int{1, 2}, 3))
	//fmt.Println(question.IsPossibleDivide([]int{1, 1, 2, 2, 3, 3}, 2))

	//fmt.Println(question.NumberToWords(0))
	//fmt.Println(question.NumberToWords(1))
	//fmt.Println(question.NumberToWords(10))
	//fmt.Println(question.NumberToWords(19))
	//fmt.Println(question.NumberToWords(20))
	//fmt.Println(question.NumberToWords(21))
	//fmt.Println(question.NumberToWords(99))
	//fmt.Println(question.NumberToWords(100))
	//fmt.Println(question.NumberToWords(101))
	//fmt.Println(question.NumberToWords(109))
	//fmt.Println(question.NumberToWords(110))
	//fmt.Println(question.NumberToWords(111))
	//fmt.Println(question.NumberToWords(119))
	//fmt.Println(question.NumberToWords(120))
	//fmt.Println(question.NumberToWords(121))
	//fmt.Println(question.NumberToWords(998))
	//fmt.Println(question.NumberToWords(999))
	//fmt.Println(question.NumberToWords(1000))
	//fmt.Println(question.NumberToWords(1001))
	//fmt.Println(question.NumberToWords(1099))
	//fmt.Println(question.NumberToWords(1100))
	//fmt.Println(question.NumberToWords(1101))
	//fmt.Println(question.NumberToWords(1109))
	//fmt.Println(question.NumberToWords(1110))
	//fmt.Println(question.NumberToWords(1111))
	//fmt.Println(question.NumberToWords(9990))
	//fmt.Println(question.NumberToWords(9991))
	//fmt.Println(question.NumberToWords(9999))
	//fmt.Println(question.NumberToWords(10000))
	//fmt.Println(question.NumberToWords(10001))
	//fmt.Println(question.NumberToWords(1000000))
	//fmt.Println(question.NumberToWords(100000))
	//fmt.Println(question.NumberToWords(1000000000000))
	//fmt.Println(question.NumberToWords(10000000000000))
	//fmt.Println(question.NumberToWords(100000000000000))
	//fmt.Println(question.NumberToWords(100000000000001))
	//fmt.Println(question.NumberToWords(100000000100001))
	//fmt.Println(question.NumberToWords(100001))
	//fmt.Println()
	//fmt.Println(question.NumberToWords(123))
	//fmt.Println(question.NumberToWords(12345))
	//fmt.Println(question.NumberToWords(1234567))
	//fmt.Println(question.NumberToWords(1234567891))

	//fmt.Println(question.MinMoves2([]int{1})) //0
	//fmt.Println(question.MinMoves2([]int{1,2,3})) //2
	//fmt.Println(question.MinMoves2([]int{1,10,2,9})) // 16
	//fmt.Println(question.MinMoves2([]int{1,100,2,9})) //106

	//fmt.Println(question.DiffWaysToCompute("2-1-1"))
	//fmt.Println(question.DiffWaysToCompute("2*3-4*5"))
	//fmt.Println(question.DiffWaysToCompute("12*32+4-42*5"))

	//	fmt.Println(question.CountOrders(1))
	//	fmt.Println(question.CountOrders(2))
	//	fmt.Println(question.CountOrders(3))
	//	fmt.Println(question.CountOrders(4))
	//	fmt.Println(question.CountOrders(5))

	//fmt.Println(question.MaximumElementAfterDecrementingAndRearranging([]int{2,2,1,2,1}))
	//fmt.Println(question.MaximumElementAfterDecrementingAndRearranging([]int{100,1,1000}))
	//fmt.Println(question.MaximumElementAfterDecrementingAndRearranging([]int{1,2,3,4,5}))
	//fmt.Println(question.MaximumElementAfterDecrementingAndRearranging([]int{1,2,3,4,7}))
	//fmt.Println(question.MaximumElementAfterDecrementingAndRearranging([]int{5,5,5,5,5,5,5,5,5}))
	//fmt.Println(question.MaximumElementAfterDecrementingAndRearranging([]int{209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,10000,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,1,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209,209}))

	//room:=question.Constructor0855(1000000)
	//fmt.Println(room.Seat())
	//room.Leave(0)
	//fmt.Println(room.Seat())
	//room.Leave(0)
	//fmt.Println(room.Seat())
	//room.Leave(0)
	//fmt.Println(room.Seat())
	//room.Leave(0)
	//fmt.Println(room.Seat())
	//room.Leave(0)
	//
	//fmt.Println()
	//room=question.Constructor0855(10)
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//room.Leave(4)
	//fmt.Println(room.Seat())
	//
	//fmt.Println()
	//room=question.Constructor0855(10)
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//room.Leave(0)
	//room.Leave(4)
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//fmt.Println(room.Seat())
	//room.Leave(0)

	//fmt.Println(question.NumWays(3,2))
	//fmt.Println(question.NumWays(2,4))
	//fmt.Println(question.NumWays(4,2))
	//fmt.Println(question.NumWays(7,8))
	//fmt.Println(question.NumWays(27,7))

	//fmt.Println(question.IntToRoman(3))
	//fmt.Println(question.IntToRoman(4))
	//fmt.Println(question.IntToRoman(9))
	//fmt.Println(question.IntToRoman(11))
	//fmt.Println(question.IntToRoman(58))
	//fmt.Println(question.IntToRoman(79))
	//fmt.Println(question.IntToRoman(92))
	//fmt.Println(question.IntToRoman(193))
	//fmt.Println(question.IntToRoman(493))
	//fmt.Println(question.IntToRoman(593))
	//fmt.Println(question.IntToRoman(893))
	//fmt.Println(question.IntToRoman(993))
	//fmt.Println(question.IntToRoman(1093))
	//fmt.Println(question.IntToRoman(1493))
	//fmt.Println(question.IntToRoman(1593))
	//fmt.Println(question.IntToRoman(1693))
	//fmt.Println(question.IntToRoman(1893))
	//fmt.Println(question.IntToRoman(1993))
	//fmt.Println(question.IntToRoman(3994))

	//fmt.Println(question.StoneGame([]int{5,3,4,5}))
	//fmt.Println(question.StoneGame([]int{4,1,7,3,8,4,7,4,5,10}))
	//fmt.Println(question.StoneGame([]int{3,7,2,3}))

	//fmt.Println(question.StoneGameII([]int{1,1,1}))
	//fmt.Println(question.StoneGameII([]int{2,7,9,4,4}))
	//fmt.Println(question.StoneGameII([]int{1,2,3,4,5,100}))

	//fmt.Println(question.StoneGameIII([]int{1,2,3,6}))
	//fmt.Println(question.StoneGameIII([]int{1,2,3,-1,-2,-3,7}))
	//
	//fmt.Println(question.StoneGameIII([]int{1,2,3,7}))
	//fmt.Println(question.StoneGameIII([]int{1,2,3,-9}))

	//fmt.Println(question.IsCousins(&question.TreeNode{
	//	Val:   1,
	//	Left:  &question.TreeNode{
	//		Val:   2,
	//		Left:  &question.TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//	Right: &question.TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//},4,3))
	//
	//fmt.Println(question.IsCousins(&question.TreeNode{
	//	Val:   1,
	//	Left:  &question.TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: &question.TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &question.TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//},4,3))

	//fmt.Println(question.WinnerSquareGame(1))
	//fmt.Println(question.WinnerSquareGame(2))
	//fmt.Println(question.WinnerSquareGame(3))
	//fmt.Println(question.WinnerSquareGame(4))
	//fmt.Println(question.WinnerSquareGame(5))
	//fmt.Println(question.WinnerSquareGame(6))
	//fmt.Println(question.WinnerSquareGame(7))
	//fmt.Println(question.WinnerSquareGame(8))

	//fmt.Println(question.StoneGameV([]int{6,2}))
	//fmt.Println(question.StoneGameV([]int{6,2,3}))
	//fmt.Println(question.StoneGameV([]int{6,2,4}))
	//fmt.Println(question.StoneGameV([]int{6,2,3,4}))
	//fmt.Println(question.StoneGameV([]int{6,2,3,4,5}))
	//fmt.Println(question.StoneGameV([]int{6,2,3,4,5,5}))
	//fmt.Println(question.StoneGameV([]int{7,7,7,7,7,7}))
	//fmt.Println(question.StoneGameV([]int{4}))

	//fmt.Println(question.CountTriplets([]int{2,3,1,6,7}))
	//fmt.Println(question.CountTriplets([]int{1,1,1,1,1}))
	//fmt.Println(question.CountTriplets([]int{2,3}))
	//fmt.Println(question.CountTriplets([]int{1,3,5,7,9}))
	//fmt.Println(question.CountTriplets([]int{7,11,12,9,5,2,7,17,22}))

	//	fmt.Println(question.StoneGameVI([]int{1,3},[]int{2,1}))
	//	fmt.Println(question.StoneGameVI([]int{1,2},[]int{3,1}))
	//	fmt.Println(question.StoneGameVI([]int{2,4,3},[]int{1,6,7}))

	//fmt.Println(question.StoneGameVII([]int{5,3,1,4,2}))
	//fmt.Println(question.StoneGameVII([]int{7,90,5,1,100,10,10,2}))

	//fmt.Println(question.TopKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"},2))
	//fmt.Println(question.TopKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},2))
	//fmt.Println(question.TopKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},4))

	//fmt.Println(question.MaxUncrossedLines([]int{1,4,2},[]int{1,2,4})) //2
	//fmt.Println(question.MaxUncrossedLines([]int{2,5,1,2,5},[]int{10,5,2,1,5,2})) //3
	//fmt.Println(question.MaxUncrossedLines([]int{1,3,7,1,7,5},[]int{1,9,2,5,1})) //2
	//fmt.Println(question.MaxUncrossedLines([]int{1,4,1,2,4},[]int{1,2,4})) //3
	//fmt.Println(question.MaxUncrossedLines([]int{1,5,2,4,4,1,2,4},[]int{1,2,4,2,4,5,4}))//5
	//fmt.Println(question.MaxUncrossedLines([]int{1},[]int{1,3})) //1

	//fmt.Println(question.MaximizeXor([]int{0,1,2,3,4},[][]int{{3,1},{1,3},{5,6}}))
	//fmt.Println(question.MaximizeXor([]int{5,2,4,6,6,3},[][]int{{12,4},{8,1},{6,3}}))
	//fmt.Println(question.MaximizeXor([]int{536870912,0,534710168,330218644,142254206},[][]int{{558240772,1000000000},{307628050,1000000000},{3319300,1000000000},{2751604,683297522},{214004,404207941}}))

	//fmt.Println(question.StrangePrinter("aaabbb"))
	//fmt.Println(question.StrangePrinter("aba"))
	//fmt.Println(question.StrangePrinter("abc"))
	//fmt.Println(question.StrangePrinter("abcacca"))
	//fmt.Println(question.StrangePrinter("aaabbdwohdaihoiwqgsagfogwofhiahofwb"))

	//fmt.Println(question.MinChanges([]int{1,2,0,3,0},1)) //3
	//fmt.Println(question.MinChanges([]int{3,4,5,2,1,7,3,4,7},3)) //3
	//fmt.Println(question.MinChanges([]int{1,2,4,1,2,5,1,2,6},3)) // 3
	//fmt.Println(question.MinChanges([]int{1,2,4,1,2,5,1,2,6},4))
	//fmt.Println(question.MinChanges([]int{1,2,4,1,2,5,1,2,6},5))

	//fmt.Println(question.ReverseParentheses("(abcd)"))
	//fmt.Println(question.ReverseParentheses("(u(love)i)"))
	//fmt.Println(question.ReverseParentheses("(ed(et(oc))el)"))
	//fmt.Println(question.ReverseParentheses("a(bcdefghijkl(mno)p)q"))
	//fmt.Println(question.ReverseParentheses("a((bcdefghijkl)(mno)p)q"))

	//solution := question.Constructor0710(2, []int{})
	//fmt.Println(solution.Pick())
	//fmt.Println(solution.Pick())
	//fmt.Println(solution.Pick())
	//
	//solution = question.Constructor0710(1, []int{})
	//fmt.Println(solution.Pick())
	//fmt.Println(solution.Pick())
	//fmt.Println(solution.Pick())
	//
	//solution = question.Constructor0710(3, []int{1})
	//fmt.Println(solution.Pick())
	//fmt.Println(solution.Pick())
	//fmt.Println(solution.Pick())
	//
	//solution = question.Constructor0710(4, []int{2})
	//fmt.Println(solution.Pick())
	//fmt.Println(solution.Pick())
	//fmt.Println(solution.Pick())
	//fmt.Println(solution.Pick())
	//fmt.Println(solution.Pick())

	//fmt.Println(question.TotalHammingDistance1([]int{4,14,2}))
	//fmt.Println(question.TotalHammingDistance1([]int{4,14,4}))
	//fmt.Println(question.TotalHammingDistance1([]int{4,14,4,5}))

	//fmt.Println(question.IsPowerOfTwo(1))
	//fmt.Println(question.IsPowerOfTwo(2))
	//fmt.Println(question.IsPowerOfTwo(3))
	//fmt.Println(question.IsPowerOfTwo(4))

	//fmt.Println(question.DistributeCoins(&question.TreeNode{
	//	Val:   0,
	//	Left:  &question.TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &question.TreeNode{
	//		Val:   0,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}))
	//fmt.Println(question.DistributeCoins(&question.TreeNode{
	//	Val:   1,
	//	Left:  &question.TreeNode{
	//		Val:   0,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &question.TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}))
	//fmt.Println(question.DistributeCoins(&question.TreeNode{
	//	Val:   1,
	//	Left:  &question.TreeNode{
	//		Val:   0,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &question.TreeNode{
	//		Val:   0,
	//		Left:  nil,
	//		Right: &question.TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}))
	//fmt.Println(question.DistributeCoins(&question.TreeNode{
	//	Val:   0,
	//	Left:  &question.TreeNode{
	//		Val:   5,
	//		Left:  &question.TreeNode{
	//			Val:   0,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &question.TreeNode{
	//			Val:   1,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &question.TreeNode{
	//		Val:   2,
	//		Left:  &question.TreeNode{
	//			Val:   0,
	//			Left:  &question.TreeNode{
	//				Val:   0,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: nil,
	//		},
	//		Right: &question.TreeNode{
	//			Val:   0,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}))
	//fmt.Println(question.IsPowerOfFour(-1)) // F
	//fmt.Println(question.IsPowerOfFour(0)) // F
	//fmt.Println(question.IsPowerOfFour(1)) // T
	//fmt.Println(question.IsPowerOfFour(2)) // F
	//fmt.Println(question.IsPowerOfFour(3)) // F
	//fmt.Println(question.IsPowerOfFour(4)) // T
	//fmt.Println(question.IsPowerOfFour(5)) // F
	//fmt.Println(question.IsPowerOfFour(8)) // F
	//fmt.Println(question.IsPowerOfFour(12)) // F
	//fmt.Println(question.IsPowerOfFour(16)) // T

	//fmt.Println(question.CanEat([]int{7,4,5,3,8},[][]int{{0,2,2},{4,2,4},{2,13,10000}}))
	//fmt.Println(question.CanEat([]int{5,2,6,4,1},[][]int{{3,1,2},{4,10,3},{3,10,100},{4,100,30},{1,3,1}}))

	//fmt.Println(question.CheckSubarraySum([]int{23,2,4,6,7},6))
	//fmt.Println(question.CheckSubarraySum([]int{23,2,6,4,7},6))
	//fmt.Println(question.CheckSubarraySum([]int{23,2,6,4,7},13))

	//fmt.Println(question.FindMaxLength([]int{0,1}))
	//fmt.Println(question.FindMaxLength([]int{0,1,0}))
	//fmt.Println(question.FindMaxLength([]int{0,1,0,1}))
	//fmt.Println(question.FindMaxLength([]int{0,1,0,1,1}))
	//fmt.Println(question.FindMaxLength([]int{0,1,0,1,1,0}))
	//fmt.Println(question.FindMaxLength([]int{0,1,0,0,0,1,0,1}))
	//fmt.Println(question.FindMaxLength([]int{0,0,0,1,1,1,0}))

	//fmt.Println(question.FindMaxForm([]string{"10","0001","111001","1","0"},5,3))
	//fmt.Println(question.FindMaxForm([]string{"10","0","1"},1,1))

	//fmt.Println(question.FindTargetSumWays([]int{1,1,1,1,1},3))
	//fmt.Println(question.FindTargetSumWays([]int{1},1))
	//fmt.Println(question.FindTargetSumWays([]int{0,0,0,0,0,0,0,0,1},1))

	//fmt.Println(question.LastStoneWeightII([]int{2,7,4,1,8,1}))
	//fmt.Println(question.LastStoneWeightII([]int{31,26,33,21,40}))
	//fmt.Println(question.LastStoneWeightII([]int{1,2}))
	//fmt.Println(question.LastStoneWeightII([]int{1,1,4,2,2}))

	//fmt.Println(question.ProfitableSchemes(5, 3, []int{2, 2}, []int{2, 3}))
	//fmt.Println(question.ProfitableSchemes(10, 5, []int{2, 3, 5}, []int{6, 7, 8}))

	//fmt.Println(question.LargestNumber1449([]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 9))
	//fmt.Println(question.LargestNumber1449([]int{7, 6, 5, 5, 5, 6, 8, 7, 8}, 12))
	//fmt.Println(question.LargestNumber1449([]int{2, 4, 6, 2, 4, 6, 4, 4, 4}, 5))
	//fmt.Println(question.LargestNumber1449([]int{6, 10, 15, 40, 40, 40, 40, 40, 40}, 47))
	//fmt.Println(question.LargestNumber1449([]int{1000,30,105,70,42,1000,1000,1000,1000},503))

	//fmt.Println(question.Answer1833([]int{1,3,2,4,1},7))
	//fmt.Println(question.Answer1833([]int{10,6,8,7,7,8},5))
	//fmt.Println(question.Answer1833([]int{1,6,3,1,2,5},20))

	//fmt.Println(question.Answer0645([]int{1,2,2,4}))
	//fmt.Println(question.Answer0645([]int{1,1}))

	//fmt.Println(question.Answer0646([][]int{
	//	{1,2},{2,3},{3,4},
	//}))
	//fmt.Println(question.Answer0646([][]int{
	//	{1,2},{7,8},{4,5},
	//}))

	//fmt.Println(question.Answer0726("H2O"))
	//fmt.Println(question.Answer0726("H2O2"))
	//fmt.Println(question.Answer0726("Be32"))

	//fmt.Println(question.Answer1418(
	//	[][]string{
	//	{"David","3","Ceviche"},
	//	{"Corina","10","Beef Burrito"},
	//	{"David","3","Fried Chicken"},
	//	{"Carla","5","Water"},
	//	{"Carla","5","Ceviche"},
	//	{"Rous","3","Ceviche"},
	//	}))
	//fmt.Println(question.Answer1418(
	//	[][]string{
	//	{"James","12","Fried Chicken"},
	//	{"Ratesh","12","Fried Chicken"},
	//	{"Amadeus","12","Fried Chicken"},
	//	{"Adam","1","Canadian Waffles"},
	//	{"Brianna","1","Canadian Waffles"},
	//	}))

	//fmt.Println(question.Answer0930([]int{1,0,1,0,1},2))
	//fmt.Println(question.Answer0930([]int{0,0,0,0,0},0))
	//fmt.Println(question.Answer0930([]int{0,0,1,0,0},1))

	//fmt.Println(question.Answer0238([]int{1,2,3,4}))
	//fmt.Println(question.Answer0238([]int{-1,1,0,-3,3}))

	//fmt.Println(question.Answer0274([]int{3,0,6,1,5}))
	//fmt.Println(question.Answer0274([]int{1,3,1}))

	//fmt.Println(question.Answer1818([]int{1,7,5},[]int{2,3,5}))
	//fmt.Println(question.Answer1818([]int{2,4,6,8,10},[]int{2,4,6,8,10}))
	//fmt.Println(question.Answer1818([]int{1,10,4,4,2,7},[]int{9,3,5,1,7,4}))

	//fmt.Println(question.Answer1846([]int{100,1,1000}))
	//fmt.Println(question.Answer1846([]int{2,2,1,2,1}))
	//fmt.Println(question.Answer1846([]int{1,2,3,4,5}))

	//fmt.Println(question.Answer0736("(let x 2 (mult x (let x 3 y 4 (add x y))))"))
	//fmt.Println(question.Answer0736("(let x 3 x 2 x)"))
	//fmt.Println(question.Answer0736("(let x 1 y 2 x (add x y) (add x y))"))
	//fmt.Println(question.Answer0736("(let x 2 (add (let x 3 (let x 4 x)) x))"))
	//fmt.Println(question.Answer0736("(let a1 3 b2 (add a1 1) b2)"))
	//fmt.Println(question.Answer0736("(let x -2 y x y)"))
	//fmt.Println(question.Answer0736("(let x (add 12 -7) (mult x x))"))

	//fmt.Println(question.AnswerMST1002([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	//fmt.Println(question.AnswerMST1002([]string{"hos","boo","nay","deb","wow","bop","bob","brr","hey","rye","eve","elf","pup","bum","iva","lyx","yap","ugh","hem","rod","aha","nam","gap","yea","doc","pen","job","dis","max","oho","jed","lye","ram","pup","qua","ugh","mir","nap","deb","hog","let","gym","bye","lon","aft","eel","sol","jab"}))

	//fmt.Println(question.Answer1838([]int{1,2,4},5))
	//fmt.Println(question.Answer1838([]int{1,4,8,13},5))
	//fmt.Println(question.Answer1838([]int{3,9,6},2))

	//fmt.Println(question.Answer1540("input","ouput",9))
	//fmt.Println(question.Answer1540("abc","bcd",10))
	//fmt.Println(question.Answer1540("aab","bbb",27))
	//fmt.Println(question.Answer1540("iqssxdlb","dyuqrwyr",40))

	//fmt.Println(question.Answer1736("2?:?0"))
	//fmt.Println(question.Answer1736("0?:3?"))
	//fmt.Println(question.Answer1736("1?:22"))

	//fmt.Println(question.Answer1743([][]int{{2,1},{3,4},{3,2}}))
	//fmt.Println(question.Answer1743([][]int{{4,-2},{1,4},{-3,1}}))
	//fmt.Println(question.Answer1743([][]int{{100000,-100000}}))
	//fmt.Println(question.Answer1743([][]int{{-3,-9},{-5,3},{2,-9},{6,-3},{6,1},{5,3},{8,5},{-5,1},{7,2}}))

	//fmt.Println(question.Answer1104(14))
	//fmt.Println(question.Answer1104(26))
	fmt.Println(question.Answer1104(16))
	fmt.Println(question.Answer1104(15))
}

//[["eel"],["aft"],["lon"],["bye"],["gym"],["let"],["hog"],["mir"],["iva"],["brr"],["eve"],["nay"],["sol"],["pup","pup"],["max"],["bum"],["lye"],["gap"],["hey"],["boo"],["aha"],["elf"],["bob"],["hem"],["doc"],["oho"],["wow"],["deb","deb"],["hos"],["rye"],["bop"],["yap"],["ugh","ugh"],["ram"],["rod"],["nam"],["yea"],["nap"],["pen"],["job"],["lyx"],["dis"],["jed"],["jab"],["qua"]]

// Input: arr 表示输入的数组，每个仓内存储的数量
func solve(arr []int) int {
	// 首先求平均值，然后将原有的数组变成与平均值的差值
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	avg := sum / len(arr)
	// 前缀数组记录前i项的累加和
	prefixSum := make([]int, len(arr))
	prefixSum[0] = arr[0] - avg
	step := 0
	if prefixSum[0] != 0 {
		step++
	}
	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = arr[i] - avg + prefixSum[i-1]
		// 多的往后移动，少的往前移
		if prefixSum[i] != 0 {
			step++
		}
	}
	return step
}

var cnt = 0

type PrefixTreeNode struct {
	c     [26]*PrefixTreeNode
	index int
}

func NewPrefixTreeNode() *PrefixTreeNode {
	return &PrefixTreeNode{
		c:     [26]*PrefixTreeNode{},
		index: -1,
	}
}
func (root *PrefixTreeNode) addWords(words []int) int {
	if len(words) != 26 {
		return -1
	}
	p := root
	for i := 0; i < 26; i++ {
		for j := 0; j < words[i]; j++ {
			if p.c[i] == nil {
				p.c[i] = NewPrefixTreeNode()
			}
			p = p.c[i]
		}
	}
	if p.index == -1 {
		p.index = cnt
		cnt++
	}
	return p.index
}
func classify(strs []string) [][]string {
	res := make([][]string, 0, len(strs))
	root := NewPrefixTreeNode()
	// 字母组合
	// 用[26]int记录字母组合
	for i := 0; i < len(strs); i++ {
		tmp := make([]int, 26)
		for j := 0; j < len(strs[i]); j++ {
			tmp[strs[i][j]-'a']++
		}
		index := root.addWords(tmp)
		if index >= len(res) {
			res = append(res, []string{})
		}
		res[index] = append(res[index], strs[i])
	}
	return res
}

// 用values表示币值的大小(已排序），nums表示对应币值的最大数量，target表示目标币值，capacity表示背包容量
// res记录满足要求的结果
//func computeValue(values []int, nums []int, target int, capacity int) (res [][]int) {
//	for i := 0; i < len(values); i++ {
//		dfs(values, nums, target, capacity, i, []int{}, &res)
//	}
//	return res
//}
//func dfs(values []int, nums []int, target int, capacity int, cur int, curRes []int, res *[][]int) {
//	//终止条件
//	if capacity >= 0 && target <= 0 { //满足条件
//		copyCur := make([]int, len(curRes))
//		for i := 0; i < len(curRes); i++ {
//			copyCur[i] = curRes[i]
//		}
//		*res = append(*res, copyCur)
//		return
//	} else if capacity < 0 {
//		return
//	}
//	//递归
//	for i := cur; i < len(values) && nums[i] > 0; i++ { //至少当前值是可选的
//		nums[i]--
//		curRes = append(curRes, values[i])
//		if nums[i] > 0 {
//			dfs(values, nums, target-values[i], capacity-1, i, curRes, res) //当前值还可以选
//		} else {
//			dfs(values, nums, target-values[i], capacity-1, i+1, curRes, res) //当前值没有了
//		}
//		//恢复状态
//		nums[i]++
//		curRes = curRes[:len(curRes)-1] //截掉一个
//	}
//}
