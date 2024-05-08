

 type SparseVector struct {
    nonZeros map[int]int
}

func Constructor(nums []int) SparseVector {
    
    nonZeros := map[int]int{}

    for index, num := range nums {
        if num != 0 {
            nonZeros[index] = num
        }
    } 
    return SparseVector{ nonZeros }
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    var s1Map map[int]int
    var s2Map map[int]int

    // optimization
    if len(vec.nonZeros) < len((*this).nonZeros) {
        s1Map = vec.nonZeros
        s2Map = (*this).nonZeros
    } else {
        s1Map = (*this).nonZeros
        s2Map = vec.nonZeros
    }
    total := 0
    for key, val := range s1Map {
        if val2, exists := s2Map[key]; exists {
            total += val * val2
        }
    }
    return total
}

/*
type SparseVector struct {
    dst []int
}

func Constructor(nums []int) SparseVector {
    s := new(SparseVector)
    (*s).dst = make([]int, len(nums))
    copy((*s).dst, nums)
    return *s
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    s1 := (*this).dst
    s2 := vec.dst

    s1Map := make(map[int]int) // index to val
    s2Map := make(map[int]int)
    for i := 0; i < len(s1); i++ {
        if s1[i] != 0 {
            s1Map[i] = s1[i]
        }
    }
    for i := 0; i < len(s2); i++ {
        if s2[i] != 0 {
            s2Map[i] = s2[i]
        }
    }

    total := 0
    len1 := len(s1Map)
    len2 := len(s2Map)
    if len1 < len2 {
        for key, val := range s1Map {
            val2, exists := s2Map[key]
            if exists {
                total += val * val2
            }
        }
    } else {
        for key, val := range s2Map {
            val2, exists := s1Map[key]
            if exists {
                total += val * val2
            }
        }
    }
    return total
}
*/
/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
