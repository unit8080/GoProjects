// 42. Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/

// Two-pointer approach:
// So, we can say that if there is a larger bar at one end (say right), 
// we are assured that the water trapped would be dependant on height of bar 
// in current direction (from left to right). As soon as we find the bar at 
// other end (right) is smaller, we start iterating in opposite direction (from 
// right to left).
// We must maintain left_max and right_max during the iteration, but now we 
// can do it in one iteration using 2 pointers, switching between the two.

// Algorithm
// Initialize left pointer to 0 and right pointer to size-1
// While left<right, do:
//      If height[left] is smaller than height[right]
//          If height[left]≥left_max, update left_max
//          Else add left_max−height[left] to ans
//          Add 1 to left.
//      Else
//          If height[right]≥right_max, update right_max
//          Else add right_max−height[right] to ans
//          Subtract 1 from right.

func trap(height []int) int {
    left := 0
    right := len(height) -1

    left_max := 0
    right_max := 0
    ans := 0

    for left < right {
        if height[left] < height[right] {
            if height[left] > left_max {
                left_max = height[left]
            } else {
                ans += left_max - height[left]
            }
            left++
        } else {
            if height[right] > right_max {
                right_max = height[right]
            } else {
                ans += right_max - height[right]
            }
            right--
        }
    }
    return ans
}
