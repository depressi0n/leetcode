package question

// WordDictionary - word dictionary
type WordDictionary struct {
	trie *PrefixTrie
}

// Constructor - Initialize your data structure here.
func WordDictionaryConstructor() WordDictionary {
	return WordDictionary{
		trie: &PrefixTrie{},
	}
}

// AddWord - Adds a word into the data structure.
func (wd *WordDictionary) AddWord(word string) {
	wd.trie.Insert(word)
}

// Search - Returns if the word is in the data structure.
// A word could contain the dot character '.' to represent any one letter.
func (wd *WordDictionary) Search(word string) bool {
	return wd.trie.Search(word, 0)
}

// PrefixTrie - Prifix Tree
type PrefixTrie struct {
	links [26]*PrefixTrie
	isEnd bool
}

// Insert - insert a word
func (tr *PrefixTrie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		ch := word[i] - 'a'
		if tr.links[ch] == nil {
			tr.links[ch] = &PrefixTrie{}
		}
		tr = tr.links[ch]
	}
	tr.isEnd = true
}

// Search - search a word
func (tr *PrefixTrie) Search(word string, k int) bool {
	for i := k; i < len(word); i++ {
		if word[i] == '.' { //遇到通配符的处理
			for j := 0; j < len(tr.links); j++ {
				if tr.links[j] != nil && tr.links[j].Search(word, i+1) {
					return true
				}
			}
			return false
		}

		ch := word[i] - 'a'
		if tr.links[ch] == nil {
			return false
		}
		tr = tr.links[ch]
	}
	return tr.isEnd
}
