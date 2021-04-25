package question

func reverseBits190_1(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		res = res*2 + num%2
		num = num / 2
	}
	return res
}

// TODO：仅仅使用位操作来实现翻转，分治的思想
//           (b1,b2)
//     xor 10       xor 01
//     (b1,0)       (0,b2)
//     (0,b1)       (b2,0)
//           (b2,b1)
func reverseBits190_2(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}
