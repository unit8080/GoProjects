// 2963. Count the Number of Good Partitions
// https://leetcode.com/problems/count-the-number-of-good-partitions/
// Redo: 10/26/2024

func numberOfGoodPartitions(nums []int) int {
    
    n := len(nums)

    lastIndex := map[int]int{}
    for i := 0; i < n; i++ {
        lastIndex[nums[i]] = i
    }

    i := 0
    j := 0
    j = max(j, lastIndex[nums[i]])

    M := 1e9 + 7
    result := 1
    for i < n {
        if i > j {
            result = (result * 2) % int(M)
        }

        j = max(j, lastIndex[nums[i]])

        i++
    }
    return result
}
