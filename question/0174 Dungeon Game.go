package question

// 每一步只能右下
func calculateMinimumHP(dungeon [][]int) int {
	// dfs会超时
	//min:=math.MaxInt32
	//rl,cl:=len(dungeon[0]),len(dungeon)
	////visit:=make([][]bool,cl)
	////for i:=0;i<len(visit);i++{
	////	visit[i]=make([]bool,rl)
	////}
	//var dfs func(int,int,int)
	//dfs= func(i int, j int,hp int) {
	//	//visit[i][j]=true
	//	hp-=dungeon[i][j]
	//	if hp<=0{ //如果出现hp比0小的要求则变成1
	//		hp=1
	//	}
	//	if i==0 && j==0 {
	//		if hp>0 && hp<min{
	//			min=hp
	//		}
	//		return
	//	}
	//	// up
	//	if i>0 {
	//		dfs(i-1,j,hp)
	//	}
	//	// left
	//	if j>0{
	//		dfs(i,j-1,hp)
	//	}
	//	//visit[i][j]=false
	//	return
	//}
	//dfs(cl-1,rl-1,1)
	//return min

	//用dp
	rl, cl := len(dungeon[0]), len(dungeon)
	hp := make([][]int, cl)
	for i := 0; i < len(hp); i++ {
		hp[i] = make([]int, rl)
	}
	// 从下往上
	// 初始化
	hp[cl-1][rl-1] = 1 - dungeon[cl-1][rl-1]
	if hp[cl-1][rl-1] <= 0 {
		hp[cl-1][rl-1] = 1
	}
	for i := rl - 2; i >= 0; i-- {
		hp[cl-1][i] = hp[cl-1][i+1] - dungeon[cl-1][i]
		if hp[cl-1][i] <= 0 {
			hp[cl-1][i] = 1
		}
	}
	for i := cl - 2; i >= 0; i-- {
		hp[i][rl-1] = hp[i+1][rl-1] - dungeon[i][rl-1]
		if hp[i][rl-1] <= 0 {
			hp[i][rl-1] = 1
		}
	}
	for i := cl - 2; i >= 0; i-- {
		for j := rl - 2; j >= 0; j-- {
			if hp[i][j+1] < hp[i+1][j] {
				hp[i][j] = hp[i][j+1] - dungeon[i][j]
			} else {
				hp[i][j] = hp[i+1][j] - dungeon[i][j]
			}
			if hp[i][j] <= 0 {
				hp[i][j] = 1
			}
		}
	}
	return hp[0][0]
}
