// 157. Read N Characters Given Read4
// https://leetcode.com/problems/read-n-characters-given-read4/

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf4 []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf4 := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf4) // read4 returns 4. Now buf = ['a','b','c','d'], fp points to 'e'
 * read4(buf4) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf4) // read4 returns 3. Now buf = ['i','j','k',...], fp points to end of file
 */

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		var data []byte

		count := 0
		b := make([]byte, 4)
		for len(data) < n {
			c := read4(b)
			if c == 0 {
				break
			}

			data = append(data, b[:c]...)
			count += c
		}
		count = min(count, n)
        copy(buf, data[:count])
		return count
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// var solution = func(read4 func([]byte) int) func([]byte, int) int {
//     // implement read below.
//     return func(buf []byte, n int) int {
//         count := 0
//         buf4 := make([]byte, 4)
//         for n > 0 {

//             c := read4(buf4)

//             if n < c  {
//                 c = n
//             }

//             if c > 0 {
//                 for i := 0; i < c; i++ {
//                     buf[count] = buf4[i]
//                     count += 1
//                 }
//             } else {
//                 return count
//             }
//             n = n - c
//         }
//         return count
//     }
// }
