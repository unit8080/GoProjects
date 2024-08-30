// 3. Longest Substring Without Repeating Characters
/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/


func lengthOfLongestSubstring(str string) int {

	left := 0
	right := 0
	strMap := map[byte]int{}

	ans := 0
	for right < len(str) {

		strMap[str[right]] += 1
		for strMap[str[right]] > 1 {
			strMap[str[left]] -= 1
			left++
		}

		ans = max(ans, right-left+1)
		right++
	}
	return ans
}
