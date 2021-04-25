package question

type Trie struct {
	end      bool
	children [26]*Trie
}

/** Initialize your data structure here. */
func TrieConstructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	pre := this
	for i := 0; i < len(word); i++ {
		if pre.children[int(word[i]-'a')] == nil {
			pre.children[int(word[i]-'a')] = &Trie{}
		}
		pre = pre.children[word[i]-'a']
	}
	pre.end = true
	return
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		if cur.children[int(word[i]-'a')] == nil {
			return false
		}
		cur = cur.children[int(word[i]-'a')]
	}
	if cur.end {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		if cur.children[int(prefix[i]-'a')] == nil {
			return false
		}
		cur = cur.children[int(prefix[i]-'a')]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
