package question

func numJewelsInStones(jewels string, stones string) int {
	cnt := 0
	jewel := make(map[int32]struct{})
	for _, c := range jewels {
		jewel[c] = struct{}{}
	}
	for _, c := range stones {
		if _, ok := jewel[c]; ok {
			cnt++
		}
	}
	return cnt
}
