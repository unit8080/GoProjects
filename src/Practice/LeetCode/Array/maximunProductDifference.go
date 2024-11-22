// 1913. Maximum Product Difference Between Two Pairs
// https://leetcode.com/problems/maximum-product-difference-between-two-pairs/

func maxProductDifference(nums []int) int {
    
    largest := math.MinInt64
    secLargest := math.MinInt64

    smallest := math.MaxInt64
    secSmallest := math.MaxInt64

    for i := 0; i < len(nums); i++ {
        num := nums[i]
        if num > largest {
            secLargest = largest
            largest = num
        } else if num > secLargest {
            secLargest = num
        }
        if num < smallest {
            secSmallest = smallest
            smallest = num
        } else if num < secSmallest {
            secSmallest = num
        }
    }
    return largest * secLargest - smallest * secSmallest
}
