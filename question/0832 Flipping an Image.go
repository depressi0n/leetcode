package question

func FlipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i])/2; j++ {
			image[i][j], image[i][len(image[i])-j-1] = image[i][len(image[i])-j-1]^1, image[i][j]^1
		}
		//  长度为奇数则中间特殊处理
		if len(image[i])&1 == 1 {
			image[i][len(image[i])>>1] ^= 1
		}
	}
	return image
}
