// 121. Best Time to Buy and Sell Stock
/*
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
*/
// func maxProfit(prices []int) int {
//     minPrice := math.MaxInt
//     maxProfit := 0

//     for i := 0; i < len(prices); i++ {
//         if prices[i] < minPrice {
//             minPrice = prices[i]
//         } else {
//             maxProfit = max(maxProfit, prices[i] - minPrice)
//         }
//     }
//     return maxProfit
// }

func maxProfit(prices []int) int {

    n := len(prices)
    profit := 0

    left := 0
    right := 1
    lastPrice := prices[left]
    for right < n {
        price := prices[right]
        if price > lastPrice {
            profit = max(profit, price - lastPrice)
        }
        if price < lastPrice {
            lastPrice = price
        }
        right++
    }
    return profit
}
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
