// 605. Can Place Flowers
// https://leetcode.com/problems/can-place-flowers/
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
