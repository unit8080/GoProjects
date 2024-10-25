// 151. Reverse Words in a String
// https://leetcode.com/problems/reverse-words-in-a-string/

func reverseBytes(b []byte, l, r int) []byte {

	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
	return b
}
func reverseWords(s string) string {

	b := []byte(s)
    len1 := len(b)

    l := 0
    for l < len1 && b[l] == ' ' {
        l++
    }

    r := len1 - 1
    for r >= 0 && b[r] == ' ' {
        r--
    }
    
    len2 := r +1
	b = reverseBytes(b, l, r)

    i := 0
    l = 0
    r = 0

	for i < len2 {
		for i < len2 {
            if b[i] != ' ' {
                b[r] = b[i]
                r++
                i++
            } else {
                break
            }
		}

		if l < r {
			reverseBytes(b, l, r-1)
			if r >= len2 {
				break
			}
			b[r] = ' '
			r++
			l = r
		}
		i++
	}

    if l == r {
        b = b[:r-1]
    } else {
        b = b[:r]
    }
	
	return string(b)
}

