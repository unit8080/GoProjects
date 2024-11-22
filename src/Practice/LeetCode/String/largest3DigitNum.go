// 2264. Largest 3-Same-Digit Number in String
// https://leetcode.com/problems/largest-3-same-digit-number-in-string/

func largestGoodInteger(num string) string {

	if len(num) < 3 {
		return ""
	}
	maxDigit := -1
	for i := 2; i < len(num); i++ {
		if num[i] == num[i-1] && num[i] == num[i-2] {
			// found three same digits
			if int(num[i]) > maxDigit {
				maxDigit = int(num[i])
			}
		}
	}
	if maxDigit != -1 {
		return string(maxDigit) + string(maxDigit) + string(maxDigit)
	} else {
		return ""
	}
}

