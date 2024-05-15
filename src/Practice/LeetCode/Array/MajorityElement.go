// 169. Majority Element
// https://leetcode.com/problems/majority-element/
// Approach 7: Boyer-Moore Voting Algorithm
// TC : O(N)
// SC : O(1)

func majorityElement(nums []int) int {
    
    count := 0
    var candidate int
    for _, num := range nums {
        if count == 0 {
            candidate = num
        }
        if num == candidate {
            count++
        } else {
            count--
        }
    }
    return candidate
}

