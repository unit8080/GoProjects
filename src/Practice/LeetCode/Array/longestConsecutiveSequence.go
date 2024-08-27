
// 128. Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/

func longestConsecutive(nums []int) int {
    
    numMap := make(map[int]bool)
    for _, num := range nums {
        numMap[num] = true
    }

    longestStreak := 0

    for _, num := range nums {
        
        
        if _, exists := numMap[num-1]; !exists {
            currentStreak := 1
            currNum := num
            for exists := numMap[currNum+1]; exists; exists = numMap[currNum+1] {
                currentStreak++
                currNum++
            }
            if currentStreak > longestStreak {
                longestStreak = currentStreak
            }
        }
        
    }
    return longestStreak
}
