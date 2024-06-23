// 274. H-Index
// https://leetcode.com/problems/h-index/

func hIndex(citations []int) int {
    
    sort.Ints(citations)
    size := len(citations)

    if citations[size-1] == 0 {
        return 0
    } else if size == 1 {
        return 1
    }

    h_index := 0
    for h := 1; h <= size; h++ {
        if citations[size -h] >= h {
            h_index = h
        } else {
            break
        }
    }
    return h_index
}
