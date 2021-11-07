package question

import (
	"container/heap"
	"sort"
)

func TopKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	// 计数
	for _, w := range words {
		cnt[w]++
	}
	// 出现过的单词
	uniqueWords := make([]string, 0, len(cnt))
	for w := range cnt {
		uniqueWords = append(uniqueWords, w)
	}
	// 对出现过的单词排序
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return cnt[s] > cnt[t] || cnt[s] == cnt[t] && s < t
	})
	return uniqueWords[:k]
}

//
type pair struct {
	w string
	c int
}

// 优先级队列，优先按照c排序，其次按照w排序
type hp0692 []pair

func (h hp0692) Len() int            { return len(h) }
func (h hp0692) Less(i, j int) bool  { a, b := h[i], h[j]; return a.c < b.c || a.c == b.c && a.w > b.w }
func (h hp0692) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp0692) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp0692) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	h := &hp0692{}
	for w, c := range cnt {
		heap.Push(h, pair{w, c}) //先push，push完再弹出最小的一个
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(pair).w
	}
	return ans
}
