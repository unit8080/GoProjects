// 347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/description/

func topKFrequent(nums []int, k int) []int {
    
    // build frequency map
    count := make(map[int]int)
    for _, num := range nums {
        count[num] += 1
    }

    // build unique number list
    unique := []int{}
    for key, _ := range count {
        unique = append(unique, key)
    }

    //partition in ascending order
    partition := func(l, r int) int {

        for l < r {
            if count[unique[l+1]] <= count[unique[l]] {
                unique[l+1], unique[l] = unique[l], unique[l+1]
                l++
            } else if count[unique[r]] > count[unique[l]] {
                r--
            } else {
                unique[r], unique[l+1] = unique[l+1], unique[r]
            }
        }
        return l

    }

    var qselect func(l, r int) int
    qselect = func(l, r int) int {
        
        pIdx := partition(l, r)
        if pIdx  < len(unique) - k {
            return qselect(pIdx+1, r)
        } else if pIdx  > len(unique) - k {
            return qselect(0, pIdx -1)
        } else {
            return pIdx
        }
    }

    kIdx := qselect(0, len(unique)-1)

    ans := []int{}
    for kIdx < len(unique) {
        ans = append(ans, unique[kIdx])
        kIdx++
    }
    return ans
}
