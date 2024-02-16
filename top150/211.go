package top150

type Trie struct {
	chars [26]*Trie
	word  bool
}

func (s *Trie) Insert(word string) {
	n := s
	for _, ch := range word {
		v := ch - 'a'
		if n.chars[v] == nil {
			n.chars[v] = &Trie{}
		}
		n = n.chars[v]
	}
	n.word = true
}

type WordDictionary struct {
	root *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{root: &Trie{}}
}

func (this *WordDictionary) AddWord(word string) {
	this.root.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	n := this.root
	var dfs func(int, *Trie) bool
	dfs = func(ind int, n *Trie) bool {
		if ind == len(word) {
			return n.word
		}
		ch := word[ind]
		if ch != '.' {
			v := ch - 'a'
			n = n.chars[v]
			if n != nil && dfs(ind+1, n) {
				return true
			}
		} else {
			for _, nn := range n.chars {
				if nn != nil && dfs(ind+1, nn) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, n)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
