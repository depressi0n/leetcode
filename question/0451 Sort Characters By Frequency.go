package question

import "bytes"

func Answer0451(s string) string {
	return frequencySort(s)
}

func frequencySort(s string) string {
	// 主要思想：先计数r，然后排序
	//// 计数
	//cnt := map[byteint{}
	//for i := range s {
	//	cnt[s[i]]++
	//}
	//// 用pair进行排序
	//type pair struct {
	//	ch  byte
	//	cnt int
	//}
	//pairs := make([]pair, 0, len(cnt))
	//for k, v := range cnt {
	//	pairs = append(pairs, pair{k, v})
	//}
	//sort.Slice(pairs, func(i, j int) bool { return pairs[i].cnt > pairs[j].cnt })
	//// 输出结果
	//ans := make([]byte, 0, len(s))
	//for _, p := range pairs {
	//	ans = append(ans, bytes.Repeat([]byte{p.ch}, p.cnt)...)
	//}
	//return string(ans)

	// 优化思路：次数会有上限，用桶排序优化排序，实际上就是两次map，第一次是统计次数，第二次对次数进行map
	// 计数，并统计最高次数
	cnt := map[byte]int{}
	maxFreq := 0
	for i := range s {
		cnt[s[i]]++
		maxFreq = max(maxFreq, cnt[s[i]])
	}
	buckets := make([][]byte, maxFreq+1)
	for ch, c := range cnt {
		buckets[c] = append(buckets[c], ch)
	}

	ans := make([]byte, 0, len(s))
	for i := maxFreq; i > 0; i-- {
		for _, ch := range buckets[i] {
			ans = append(ans, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(ans)
}
