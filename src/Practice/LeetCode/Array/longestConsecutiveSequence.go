
// 128. Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/

// Algorithm:
// - find the number which does not have precedent number (a less than this)
// - to verify above insert all nums into map
// - keep looking next sequence number until find it in mmap
// - compute logest sequence so far

func longestConsecutive(nums []int) int {

    longestSequence := 0

    numMap := map[int]bool {}
    for _, num := range nums {
        numMap[num] = true
    }
    for _, num := range nums {
        if _, exists := numMap[num - 1]; !exists {

            currNum := num
            currSequence := 1

            for {
                currNum++
                if _, exists := numMap[currNum]; exists {
                   currSequence++
                } else {
                    break
                }
            }
            if currSequence > longestSequence {
                longestSequence = currSequence
            }
        }
    }
    return longestSequence
}


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
