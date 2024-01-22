package top150

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	dirs := [][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	trie := &Trie{}

	for _, v := range words {
		trie.Insert(v)
	}

	m, n := len(board), len(board[0])
	resMap := map[string]bool{}

	var dfs func(*Trie, int, int)
	dfs = func(node *Trie, x, y int) {
		ch := board[x][y]
		node = node.child[ch-'a']
		if node == nil {
			return
		}

		if len(node.word) != 0 {
			resMap[node.word] = true
		}

		board[x][y] = '*'
		for _, v := range dirs {
			nx, ny := v[0]+x, v[1]+y
			if nx >= 0 && nx < m && ny >= 0 && ny < n && board[nx][ny] != '*' {
				dfs(node, nx, ny)
			}
		}
		board[x][y] = ch
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(trie, i, j)
		}
	}

	for k, _ := range resMap {
		res = append(res, k)
	}

	return res
}

type Trie struct {
	child [26]*Trie
	word  string
}

func (s *Trie) Insert(word string) {
	node := s
	for _, ch := range word {
		v := ch - 'a'
		if node.child[v] == nil {
			node.child[v] = &Trie{}
		}
		node = node.child[v]
	}
	node.word = word
}
