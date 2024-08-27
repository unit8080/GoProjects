// 11. Container With Most Water
/*
https://leetcode.com/problems/container-with-most-water/
*/

// Algorith:
// - take two pointer approach
// - calculate area by width * min of l & rh height
// - move the left or right pointer whichever height is sorter
//   in search of  bigger height though decreasing the width
// - take max of area at each step

func maxArea(height []int) int {

	left := 0
	right := len(height) - 1
	ans := 0

	for left < right {

		lh := height[left]
		rh := height[right]
		water := min(lh, rh) * (right - left)
		ans = max(ans, water)

		if lh < rh {
			left++
		} else if lh > rh {
			right--
		} else {
			left++
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
