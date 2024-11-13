package trie

type Trie struct {
	alphabet [26]*Trie
	isWord   bool
}

func Constructor() Trie {
	return Trie{isWord: false}

}

func (this *Trie) Insert(word string) {
	trie := this
	for _, ch := range word {
		if trie.alphabet[ch-'a'] == nil {
			trie.alphabet[ch-'a'] = &Trie{}
		}
		trie = trie.alphabet[ch-'a']
	}
	trie.isWord = true
}

func (this *Trie) Search(word string) bool {
	trie := this
	for _, ch := range word {
		if trie.alphabet[ch-'a'] == nil {
			return false
		}
		trie = trie.alphabet[ch-'a']
	}
	return trie.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	trie := this
	for _, ch := range prefix {
		if trie.alphabet[ch-'a'] == nil {
			return false
		}
		trie = trie.alphabet[ch-'a']
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
