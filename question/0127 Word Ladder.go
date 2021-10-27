package question

import "math"

// 字典wordList 中从单词 beginWord和 endWord 的 转换序列 是一个按下述规格形成的序列：
//序列中第一个单词是 beginWord 。
//序列中最后一个单词是 endWord 。
//每次转换只能改变一个字母。
//转换过程中的中间单词必须是字典wordList 中的单词。
//给你两个单词 beginWord和 endWord 和一个字典 wordList ，找到从beginWord 到endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0。
func ladderLength(beginWord string, endWord string, wordList []string) int {
	return ladderLengthCore2(beginWord, endWord, wordList)
}

// 主要思想：第一步是处理wordlist，endWord必须在wordlist中
// 第二步，将wordlist中所有单词进行分类，相差一个字母的归为一类
// 前面两个步骤通过用map存储的方式进行处理
// 第三步，从startWord开始，经过DFS的方式，每次转变一个字符，直到能变成endWord
// 这种做法依赖wordList的大小，预处理过程的复杂度是O(n^2)，而后的DFS复杂度将十分高
func ladderLengthCore(beginWord string, endWord string, wordList []string) int {
	cache := make(map[string][]int)
	compare := func(s1 string, s2 string) bool {
		index, diff := 0, 0
		for index < len(s1) {
			if s1[index] != s2[index] {
				diff++
			}
			index++
		}
		if diff <= 1 {
			return true
		}
		return false
	}
	for i := 0; i < len(wordList); i++ {
		// 和beginWord也要对比
		if compare(wordList[i], beginWord) {
			if _, ok := cache[beginWord]; !ok {
				cache[beginWord] = []int{}
			}
			cache[beginWord] = append(cache[beginWord], i)
		}
		for j := i + 1; j < len(wordList); j++ {
			if compare(wordList[i], wordList[j]) {
				if _, ok := cache[wordList[i]]; !ok {
					cache[wordList[i]] = []int{}
				}
				cache[wordList[i]] = append(cache[wordList[i]], j)
				if _, ok := cache[wordList[j]]; !ok {
					cache[wordList[j]] = []int{}
				}
				cache[wordList[j]] = append(cache[wordList[j]], i)
			}
		}
	}
	//for s, ints := range cache {
	//	fmt.Printf("%s:", s)
	//	for _, index := range ints {
	//		fmt.Printf("%s,", wordList[index])
	//	}
	//	fmt.Println()
	//}
	if _, ok := cache[endWord]; !ok {
		return 0
	}
	res := len(wordList) + 1
	used := make(map[string]struct{})
	var dfs func(cur string, depth int)
	dfs = func(cur string, depth int) {
		if cur == endWord {
			res = min(res, depth)
			return
		}
		if depth >= res {
			return
		}
		for _, index := range cache[cur] {
			if _, ok := used[wordList[index]]; !ok {
				used[wordList[index]] = struct{}{}
				dfs(wordList[index], depth+1)
				delete(used, wordList[index])
			}
		}
	}
	dfs(beginWord, 1)
	if res == len(wordList)+1 {
		return 0
	}
	return res
}

// 为了降低上面方案的复杂度，修改预处理的方式:将单词中某个字母替换为*（创建虚拟节点），然后去匹配其他的能匹配的单词，加入同一个序列中
// 同时采用BFS的方式去轮询
func ladderLengthCore2(beginWord string, endWord string, wordList []string) int {
	// 预处理
	similar := make(map[string][]int)
	for i := 0; i < len(wordList); i++ {
		for j := 0; j < len(wordList[i]); j++ {
			t := wordList[i][:j] + "*" + wordList[i][j+1:]
			if _, ok := similar[t]; !ok {
				similar[t] = []int{}
			}
			similar[t] = append(similar[t], i)
		}
	}
	used, l := make([]bool, len(wordList)), 1
	wordList = append(wordList, beginWord)
	// 从beginWord开始
	q := []int{len(wordList) - 1}
	for len(q) > 0 {
		nextQ := []int{}
		l++
		// 遍历下一轮可能的值
		for i := 0; i < len(q); i++ {
			for j := 0; j < len(wordList[q[i]]); j++ {
				t := wordList[q[i]][:j] + "*" + wordList[q[i]][j+1:]
				for _, index := range similar[t] {
					if wordList[index] == endWord {
						return l
					}
					// 以前使用过的都不能再后面轮次中用到
					if !used[index] {
						used[index] = true
						nextQ = append(nextQ, index)
					}
				}
			}
		}
		q = nextQ
	}
	return 0
}

// 双向BFS
// 空间大小仍旧依赖于wordList，广度优先搜索的搜索空间大小依赖于每层节点的分支数量，随着层次的提升，时间复杂度将指数上升。
// 为了降低层数，考虑两个同时进行的广度优先搜索可以有效减少搜索空间
// 那么考虑一遍从beginWord开始，另一边从endWord开始
// 每次从两边扩展一层，当某个时刻发现两边都访问过同一顶点时停止搜索。
func ladderLengthCore3(beginWord string, endWord string, wordList []string) int {
	wordId := make(map[string]int)
	graph := make([][]int, 0, len(wordList))
	// 注册一个单词
	addWord := func(word string) int {
		id, ok := wordId[word]
		if !ok {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id := addWord(word)
		for i := 0; i < len(word); i++ {
			t := word[:i] + "*" + word[i+1:]
			id2 := addWord(t)
			graph[id] = append(graph[id], id2)
			graph[id2] = append(graph[id2], id)
		}
		return id
	}
	// 建立一个图，其中有一些虚拟节点
	for i := 0; i < len(wordList); i++ {
		addEdge(wordList[i])
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	distBegin := make([]int, len(wordId))
	for i := 0; i < len(distBegin); i++ {
		distBegin[i] = math.MaxInt64
	}
	distBegin[beginId] = 0
	qBegin := []int{beginId}

	distEnd := make([]int, len(wordId))
	for i := 0; i < len(distEnd); i++ {
		distBegin[i] = math.MaxInt64
	}
	distEnd[endId] = 0
	qEnd := []int{endId}

	for len(qBegin) > 0 && len(qEnd) > 0 {
		q := qBegin
		qBegin = nil
		for _, v := range q {
			if distEnd[v] < math.MaxInt64 {
				return (distBegin[v]+distEnd[v])>>1 + 1
			}
			for _, w := range graph[v] {
				if distBegin[v] == math.MaxInt64 {
					distBegin[v] = distBegin[v] + 1
					qBegin = append(qBegin, w)
				}
			}
		}
		q = qEnd
		qEnd = nil
		for _, v := range q {
			if distBegin[v] < math.MaxInt64 {
				return (distBegin[v]+distEnd[v])>>1 + 1
			}
			for _, w := range graph[v] {
				if distEnd[v] == math.MaxInt64 {
					distEnd[v] = distEnd[v] + 1
					qEnd = append(qEnd, w)
				}
			}
		}
	}
	return 0
}
