package question

// 距离问题可以想一想计数排序和桶排序，他们和比较排序具有不同的复杂度
func maximumGap164_1(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	bucketSize := (max - min) / (len(nums) - 1)
	if bucketSize == 0 {
		bucketSize = 1
	}
	bucketCount := (max-min)/bucketSize + 1
	bucket := make([][]int, bucketCount)
	for i := 0; i < len(bucket); i++ {
		bucket[i] = make([]int, 2)
	}
	for i := 0; i < len(nums); i++ {
		index := (nums[i] - min) / bucketSize
		if bucket[index][0] == 0 {
			bucket[index][0] = nums[i]
			bucket[index][1] = nums[i]
			continue
		}
		if nums[i] < bucket[index][0] { //<min
			bucket[index][0] = nums[i]
		}
		if nums[i] > bucket[index][1] { //>max
			bucket[index][1] = nums[i]
		}
	}
	prevm, res := min, 0
	for i := 0; i < bucketCount; i++ {
		if bucket[i][0] == 0 {
			continue
		}
		if res < bucket[i][0]-prevm {
			res = bucket[i][0] - prevm
		}
		prevm = bucket[i][1]
	}
	return res
}

// 0ms的方法，TODO：理解思想
const BSTEP = 8
const ISIZE = 5
const BSIZE = 1 << (BSTEP - ISIZE)

func maximumGap164_2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	return bktDist(nums, 32-BSTEP, 0xFFFFFFFF)
}

func bktDist(arr []int, bi, mask int) int {
	if len(arr) == 2 {
		return abs(arr[0] - arr[1])
	}

	bpos := [8]byte{0, 1, 0, 2, 0, 0, 0, 3}
	getBpos := func(v uint32) int {
		b := 0
		if v > 0xFFFF {
			b += 16
			v >>= 16
		}
		if v > 0xFF {
			b += 8
			v >>= 8
		}
		if v > 0xF {
			b += 4
			v >>= 4
		}
		return b + int(bpos[v-1])
	}

	ans := 0
	for ; mask > 0; bi, mask = bi-BSTEP, mask>>BSTEP {
		if bi < 0 {
			bi = 0
		}
		bitm := [BSIZE]uint32{}
		bmax, bmin := [1 << BSTEP]int{}, [1 << BSTEP]int{}
		xmax, xmin := 0, 1<<BSTEP
		for _, v := range arr {
			x := (v & mask) >> bi
			xb, xf := x>>5, uint32(1)<<(x&0x1F)
			if (bitm[xb] & xf) == 0 {
				bmax[x], bmin[x] = v, v
				bitm[xb] |= xf
				if x > xmax {
					xmax = x
				}
				if x < xmin {
					xmin = x
				}
			} else if v > bmax[x] {
				bmax[x] = v
			} else if v < bmin[x] {
				bmin[x] = v
			}
		}
		if xmin == xmax {
			continue
		}

		for i, last := 0, -1; i < len(bitm); i++ {
			for bv := bitm[i]; bv > 0; bv &= bv - 1 {
				x := i*32 + getBpos(((bv^(bv-1))>>1)+1)
				if last >= 0 && bmin[x]-last > ans {
					ans = bmin[x] - last
				}
				last = bmax[x]
			}
		}
		if bi == 0 || ans >= (1<<bi) {
			break
		}

		bkts := [1 << BSTEP][]int{}
		for _, v := range arr {
			x := (v & mask) >> bi
			bkts[x] = append(bkts[x], v)
		}
		for i, j := 0, 0; i < len(bitm); i, j = i+1, j+32 {
			for bv := bitm[i]; bv > 0; bv &= bv - 1 {
				x := i*32 + getBpos(((bv^(bv-1))>>1)+1)
				if len(bkts[x]) < 2 {
					continue
				} else {
					d := bktDist(bkts[x], bi-BSTEP, mask>>BSTEP)
					if d > ans {
						ans = d
						if ans == (1<<bi)-1 {
							break
						}
					}
				}
			}

		}
	}
	return ans
}
