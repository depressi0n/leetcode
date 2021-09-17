package question

//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

func spiralOrder(matrix [][]int) []int {
	return spiralOrderCore(matrix)
}

func spiralOrderCore(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	// 一圈一圈走，收集数据
	res := make([]int, 0, len(matrix)*len(matrix[0]))
	x,y:=0,0
	rlen,clen:=len(matrix),len(matrix[0])
	start:=0
	for{
		for ;y<clen;y++{
			res=append(res,matrix[x][y])
		}
		y--
		x++
		if x>=rlen{
			break
		}
		for ;x<rlen;x++{
			res=append(res,matrix[x][y])
		}
		x--
		y--
		if y<start{
			break
		}
		for ;y>=start;y--{
			res=append(res,matrix[x][y])
		}
		y++
		x--
		if x<start{
			break
		}
		for ;x>start;x--{
			res=append(res,matrix[x][y])
		}
		if start == rlen-2 || start == rlen-1 || start == clen-2 ||start == clen-1{
			return res
		}
		start++
		rlen--
		clen--
		x=start
		y=start
	}
	return res
}
