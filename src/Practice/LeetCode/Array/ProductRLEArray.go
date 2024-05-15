// 1868. Product of Two Run-Length Encoded Arrays
// https://leetcode.com/problems/product-of-two-run-length-encoded-arrays/

func findRLEArray(e1 [][]int, e2 [][]int) [][]int {

    res := make([][]int, 0) // len(e1)

    i := 0
    j := 0
    for i < len(e1) && j < len(e2) {
        if e1[i][1] == e2[j][1] {
            res = append(res, []int{e1[i][0]*e2[j][0], e1[i][1]})
            i++ 
            j++
        } else if e1[i][1] < e2[j][1] {
            res = append(res, []int{e1[i][0]*e2[j][0], e1[i][1]})
            e2[j][1] = e2[j][1] - e1[i][1]
            i++
        } else {
            res = append(res, []int{e1[i][0]*e2[j][0], e2[j][1]})
            e1[i][1] = e1[i][1]-e2[j][1]
            j++
        }
    }
    
    
    rlen := len(res)
    p := 0
    
    for i := 1; i < len(res); i++ {
        if res[p][0] == res[i][0] {
            res[p][1] += res[i][1]
            res[i] = nil
        
            rlen--
        } else {
            p = i
        }
    }

    final := make([][]int, 0, rlen)
    for i := range res { 
        if res[i] != nil {
            final = append(final, res[i])
        }
    }
    return final
}
