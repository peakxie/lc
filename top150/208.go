package top150

type Trie struct {
	child [26]*Trie
	word  bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		ind := v - 'a'
		if node.child[ind] == nil {
			node.child[ind] = &Trie{}
		}
		node = node.child[ind]
	}
	node.word = true

}

func (this *Trie) searchPrefix(prefix string) *Trie {
	node := this
	for _, v := range prefix {
		ind := v - 'a'
		if node.child[ind] == nil {
			return nil
		}
		node = node.child[ind]
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	if node != nil {
		return node.word
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
