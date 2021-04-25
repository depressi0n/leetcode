package question

//TODO: 用剪枝+回溯的方法
//全排列+斜线检查，但是很慢...
func solveNQueens(n int) [][]string {
	//本质上也是全排列问题
	if n == 0 {
		return [][]string{}
	}
	if n == 1 {
		return [][]string{{"Q"}}
	}
	//if n<4 {
	//	return [][]string{}
	//}
	var res [][]string
	origin := make([]int, n)
	//run:= make([]int,n)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = i
		origin[i] = i
	}

	AddToRes := func(a []int) {
		var t []string
		for i := 0; i < len(a); i++ {
			var tmp []byte
			for j := 0; j < n; j++ {
				tmp = append(tmp, '.')
			}
			tmp[a[i]] = 'Q'
			t = append(t, string(tmp))
		}
		res = append(res, t)
		return
	}
	Attack := func(a []int) bool {
		for k := 1; k < len(a); k++ {
			for i := 0; i < len(a)-k; i++ {
				if a[i]+k == a[i+k] || a[i] == a[i+k]+k {
					return true
				}
			}
		}
		return false
	}
	Equal := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	if !Attack(temp) {
		AddToRes(temp)
	}
	temp = nextPermutationUnique(temp)
	for !Equal(origin, temp) { //这里一定要防止res==0的死循环问题，如果是第一次进来，可以用res==0，后续进来不能用
		if !Attack(temp) {
			AddToRes(temp)
		}
		temp = nextPermutationUnique(temp) //run 存储新的一组，origin存储原始的一组
	}
	return res
}
