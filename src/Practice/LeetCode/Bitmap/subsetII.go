// 90. Subsets II
// https://leetcode.com/problems/subsets-ii/
// import "math"
// func subsetsWithDup(nums []int) [][]int {
    
//     ans := [][]int{}

//     n := len(nums)
//     for i := math.Pow(2, float64(n)); i < math.Pow(2, float64(n+1)); i++ {
//         bitmask := fmt.Sprintf("%b", int(i))[1:]

//         output := []int{}
//         for j := 0; j < n; j++ {
//             if bitmask[j] == '1' {
//                 output = append(output, nums[j])
//             }
//         }
//         ans = append(ans, output)
//     }
//     return ans
// }

func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    
    ans := [][]int{}
    subset(nums, 0, []int{}, &ans)
    return ans
}

func subset(nums []int, startIdx int, curr []int, ans *[][]int) {

    newCurr := make([]int, len(curr))

    copy(newCurr, curr)
    *ans = append((*ans), newCurr)

    if startIdx > len(nums) {
        return
    }
    for i := startIdx; i < len(nums); i++ {
        if i != startIdx && nums[i] == nums[i-1] {
            continue
        }
        subset(nums, i+1, append(curr, nums[i]), ans)
    }
}
