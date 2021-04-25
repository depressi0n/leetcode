package question

func IsSelfCrossing(distance []int) bool {
	//dir := [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
	// 模拟穿插行为？ 每四个一轮？但是先前的怎么考虑？
	//// 可以看成矩形的叠加？然后每次确定一个矩形，只要从矩形中通过，就肯定会产生交叉
	//// 也就是说用矩形记录之前的线的结果，进一步简化这个矩形，其实只要记录之前的线段就可以
	//// 用一个区间记录横线的范围，用另一个区间记录竖线的范围
	//// 下一次穿插时在范围内，则说明相交
	//// 初始化
	//horizontal:=[]int{0,0}
	//vertical:=[]int{0,0}
	//// 当前位置
	//cur:=[]int{0,0}
	//// 下一次方向 i mod 4
	//// i mod 2 == 1 上下走
	//// i mod 2 == 0 左右走
	//for i := 0; i < len(distance); i++ {
	//	cur[0],cur[1]=cur[0]+dir[i][0]*distance[i],cur[1]+dir[i][1]*distance[i]
	//	// 更新范围 和 检查
	//	switch i & 0x11 {
	//	case 0: // 往上走，是否穿越了
	//		vertical[1]=max(vertical[1],cur[1]) // 比较当前的y=cur[1] 上值 v[1]
	//
	//	case 1:
	//		horizontal[0]=min(horizontal[0],cur[0]) // 比较当前的x=cur[0]左值 h[0]
	//	case 2:
	//		vertical[0]=min(vertical[0],cur[1]) // 比较当前的y=cur[1] 下值 v[0]
	//	case 3:
	//		horizontal[1]=min(horizontal[1],cur[0]) // 比较当前的x=cur[0] 右值h[1]
	//	}
	//}

	// 为每个点记录两个方向的长度，第一个点是特殊点，只记录一个方向，进入前就要记录横向的方向，长度为0
	// 每次画出一条新线段时，检查在这个线段上曾经出现所有的值，是否和这条线段冲突了（最长的值超过了距离)
	recordXs := make(map[int][][]int, 500)
	recordYs := make(map[int][][]int, 500)
	cur := []int{0, 0}
	recordYs[cur[1]] = append(recordYs[cur[1]], []int{cur[0], cur[0]})
	for i := 0; i < len(distance); i++ {
		// 更新记录
		switch i & 0x3 {
		case 0: // 向上走，x=cur[0], y range cur[1],cur[1]+distance[i]
			recordXs[cur[0]] = append(recordXs[cur[0]], []int{cur[1], cur[1] + distance[i]})
			for j := cur[1] + 1; j <= cur[1]+distance[i]; j++ {
				// 判断
				if _, ok := recordYs[j]; ok {
					for _, interval := range recordYs[j] {
						if interval[0] <= cur[0] && cur[0] <= interval[1] {
							return true
						}
					}
				}

			}
			cur[1] = cur[1] + distance[i]
		case 1: // 向左走，y=cur[1], x range cur[0]-distance[i],cur[0]
			recordYs[cur[1]] = append(recordYs[cur[1]], []int{cur[0] - distance[i], cur[0]})
			for j := cur[0] - 1; j >= cur[0]-distance[i]; j-- {
				if _, ok := recordXs[j]; ok {
					for _, interval := range recordXs[j] {
						if interval[0] <= cur[1] && cur[1] <= interval[1] {
							return true
						}
					}
				}
			}
			cur[0] = cur[0] - distance[i]
		case 2: // 向下走，x=cur[0], y range cur[1],cur[1]+distance[i]
			recordXs[cur[0]] = append(recordXs[cur[0]], []int{cur[1] - distance[i], cur[1]})
			for j := cur[1] - 1; j >= cur[1]-distance[i]; j-- {
				// 判断
				if _, ok := recordYs[j]; ok {
					for _, interval := range recordYs[j] {
						if interval[0] <= cur[0] && cur[0] <= interval[1] {
							return true
						}
					}
				}

			}
			cur[1] = cur[1] - distance[i]
		case 3: // 向右走，y=cur[1], x range cur[0]-distance[i],cur[0]
			recordYs[cur[1]] = append(recordYs[cur[1]], []int{cur[0], cur[0] + distance[i]})
			for j := cur[0] + 1; j <= cur[0]+distance[i]; j++ {
				if _, ok := recordXs[j]; ok {
					for _, interval := range recordXs[j] {
						if interval[0] <= cur[1] && cur[1] <= interval[1] {
							return true
						}
					}
				}
			}
			cur[0] = cur[0] + distance[i]
		}
	}
	return false
}

// TODO: 只有一个外圈，外圈加内圈（水平）外圈加内圈（竖直） 合法的情况，和下面等价
// TODO：第四条线段只能和第一条线段相交，第5条线段只能与第1，2条线段相交，第6条线段只能与第3，2，1条线段相交，第7条只能与第4，3，2，1条线段相交 滑动窗口大小为6
