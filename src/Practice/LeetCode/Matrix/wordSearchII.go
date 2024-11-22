// 212. Word Search II
// https://leetcode.com/problems/word-search-ii/

type TrieNode struct {
    isLast bool
    word string 
    children map[byte] *TrieNode
}

var root *TrieNode

func newTrieNode() *TrieNode {
    node := new(TrieNode)
    node.isLast = false
    node.word = ""
    node.children = make(map[byte] *TrieNode, 0)
    return node
}
func findWords(board [][]byte, words []string) []string {
    
    m := len(board)    // rows
    n := len(board[0]) // columns

    root = newTrieNode()

    for _, word := range words {
        insertWord(root, word)
    }

    result := []string{}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
                findWord(board, i, j, m, n, root, &result)
        }
    }
    return result
}

var directions = [][]int{{1,0}, {-1,0},{0, -1}, {0, 1}}
func findWord(board [][]byte, i, j int, m, n int, root *TrieNode, result *[]string) {
    
    if i < 0 || i >= m || j < 0 || j >= n {
        return
    }
    if _, exists := root.children[board[i][j]]; !exists || board[i][j] == byte('$') {
        return
    }
    root = root.children[board[i][j]]
    temp := board[i][j]
    if root.isLast {
        root.isLast = false
        (*result) = append((*result), root.word)
        
    }
    board[i][j] = byte('$')
    for _, dir := range directions {
        new_i := i + dir[0]
        new_j := j + dir[1]
        findWord(board, new_i, new_j, m, n, root, result)
    }
    board[i][j] = temp
}
func insertWord(root *TrieNode, word string) {
    for i := 0; i < len(word); i++ {
        if _, exists := root.children[word[i]]; !exists {
            root.children[word[i]] = newTrieNode()
            root = root.children[word[i]]
        } else {
            root = root.children[word[i]]
        }
    }
    root.isLast = true
    root.word = word
}
