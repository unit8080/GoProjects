// 1424. Diagonal Traverse II
// https://leetcode.com/problems/diagonal-traverse-ii/

func findDiagonalOrder(nums [][]int) []int {
    
    dmap := make(map[int][]int) // key : i + j 
    maxKey := -1

    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums[i]); j++ {
            key := i + j
            if key > maxKey {
                maxKey = key
            }
            dmap[key] = append([]int{ nums[i][j] }, dmap[key]...)
        }
    }

    ans := []int{}
    for i := 0; i <= maxKey; i++ {
        ans = append(ans, dmap[i]...)
    }
    return ans
}
