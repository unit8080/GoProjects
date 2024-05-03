// 1216. Valid Palindrome III
// https://leetcode.com/problems/valid-palindrome-iii/


func isValidPalindrome(s string, k int) bool {
    queue := [][]int{{0, len(s) - 1, 0}}
	visited := make([][]bool, len(s))
	for i := range visited {
		visited[i] = make([]bool, len(s))
	}

	for len(queue) > 0 {
		l, r, currK := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]

		if currK > k {
			return false
		}

		for s[l] == s[r] {
			l++
			r--
			if l >= r {
				return true
			}
		}

		if !visited[l+1][r] {
			queue = append(queue, []int{l + 1, r, currK + 1})
			visited[l+1][r] = true
		}

		if !visited[l][r-1] {
			queue = append(queue, []int{l, r - 1, currK + 1})
			visited[l][r-1] = true
		}
	}
	return false
}

