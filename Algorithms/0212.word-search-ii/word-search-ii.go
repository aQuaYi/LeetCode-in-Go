package problem0212

func findWords(board [][]byte, words []string) []string {
	var results []string

	m := len(board)
	if m == 0 {
		return results
	}

	n := len(board[0])
	if n == 0 {
		return results
	}

	trie := buildTrie(words)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, i, j, trie, &results)
		}
	}

	return results
}

func dfs(board [][]byte, i int, j int, trie *TrieNode, results *[]string) {
	c := board[i][j]
	if c == '#' || trie.next[int(c-'a')] == nil {
		return
	}

	trie = trie.next[int(c-'a')]
	if len(trie.word) > 0 {
		// Found one
		*results = append(*results, trie.word)
		trie.word = ""
	}

	board[i][j] = '#'
	if i > 0 {
		dfs(board, i-1, j, trie, results)
	}

	if i < len(board)-1 {
		dfs(board, i+1, j, trie, results)
	}

	if j > 0 {
		dfs(board, i, j-1, trie, results)
	}

	if j < len(board[0])-1 {
		dfs(board, i, j+1, trie, results)
	}

	board[i][j] = c
}

func buildTrie(words []string) *TrieNode {
	root := new(TrieNode)
	for _, word := range words {
		cur := root
		for _, c := range word {
			cidx := int(c - 'a')
			if cur.next[cidx] == nil {
				cur.next[cidx] = new(TrieNode)
			}
			cur = cur.next[cidx]
		}
		cur.word = word
	}

	return root
}

// TrieNode 是 trie 的节点
type TrieNode struct {
	next [26]*TrieNode
	word string
}
