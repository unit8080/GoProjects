// 88. Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int)  {

    i := m -1
    j := n -1

    for k := m + n - 1; k >= 0; k-- { // in place len(nums1) = m + n

        // exhausted nums2 then all nums1 are in place
        if j < 0 {
            return
        }

        // nums1 is exhausted, but nums2 needs to be copied
        if i < 0 { 
            nums1[k] = nums2[j]
            j--
            continue
        }

        if nums1[i] >= nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
    }
}

// using Auxilliary space
/*
func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    newArr := make([]int, m+n)
    index := 0
    var i, j int

    for i < m && j < n {
        
        if nums1[i] <= nums2[j] {
            newArr[index] = nums1[i]
            i = i+1
        } else {
            newArr[index] = nums2[j]
            j = j + 1
        }
        index = index + 1
    }
    for i < m {
        newArr[index] = nums1[i]
        index = index +1
        i = i +1
    }
    for j < n {
        newArr[index] = nums2[j]
        index = index +1
        j = j +1
    }
    copy(nums1[:], newArr[:])
}
*/
