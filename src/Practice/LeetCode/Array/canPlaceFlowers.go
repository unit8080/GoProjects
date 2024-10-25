// 605. Can Place Flowers
// https://leetcode.com/problems/can-place-flowers/

// You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.

 

// Example 1:

// Input: flowerbed = [1,0,0,0,1], n = 1
// Output: true
// Example 2:

// Input: flowerbed = [1,0,0,0,1], n = 2
// Output: false
 

// Constraints:

// 1 <= flowerbed.length <= 2 * 104
// flowerbed[i] is 0 or 1.

func canPlaceFlowers(flowerbed []int, n int) bool {

	if len(flowerbed) < n {
		return false
	}

	i := 0
	for i < len(flowerbed) {
		if i == 0 && flowerbed[i] == 0 {
			if len(flowerbed) == 1 || flowerbed[i+1] == 0 {
				n--
				flowerbed[i] = 1
				i += 2
				continue
			} else {
                i++
            }
		}
		if flowerbed[i] == 1 || flowerbed[i-1] == 1 ||
			i < len(flowerbed)-1 && flowerbed[i+1] == 1 {
			i++
			continue
		}
		flowerbed[i] = 1
		n--
		i++
	}
	return n <= 0
}
