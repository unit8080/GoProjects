// 238. Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/?envType=company&envId=microsoft&favoriteSlug=microsoft-thirty-days

func productExceptSelf(nums []int) []int {
    
    L := 1
    R := 1
    ans := make([]int, len(nums))
    
    // build Left prefix product array
    for i := 0; i < len(nums); i++ {
        ans[i] = L
        L = L * nums[i]
    }
    for j := len(nums) -1; j >= 0; j-- {
        ans[j] = R * ans[j]
        R = R * nums[j]
    }
    return ans
}
