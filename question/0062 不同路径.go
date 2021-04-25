package question

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 || n == 1 {
		return 1
	}
	//res:=make([][]int,m)
	//for i:=0;i<m;i++{
	//	res[i]=make([]int,n)
	//}
	//res[m-1][n-1]=0
	//for i:=0;i<m-1;i++{
	//	res[i][n-1]=1
	//}
	//for i:=0;i<n-1;i++{
	//	res[m-1][i]=1
	//}
	//for i:=m-2;i>=0;i--{
	//	for j:=n-2;j>=0;j--{
	//		res[i][j]=res[i][j+1]+res[i+1][j]
	//	}
	//}
	//return res[0][0]
	//TODO：这里是完全可以节省空间的
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	for i := 0; i < m-1; i++ {
		for j := n - 2; j >= 0; j-- {
			res[j] = res[j] + res[j+1]
		}
	}
	return res[0]
}
