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
    var data []byte

    return func(buf []byte, n int) int {
        b := make([]byte, 4)
        
        for len(data) < n {
            c := read4(b)
            if c == 0 {
                break
            }
            
            data = append(data, b[:c]...)
        }
        
        c := min(len(data), n)
        copy(buf, data[:c])
        data = data[c:]

        return c
    }
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}
// var buffer [500]byte

// var wrIdx int = 0
// var rdIdx int = 0
// var total int = 0

// var solution = func(read4 func([]byte) int) func([]byte, int) int {
//     // implement read below.
//     return func(buf []byte, n int) int {
        
//         count := 0
//         for n > 0 {
//             buf4 := make([]byte, 4)
//             c := read4(buf4)
//             if c == 0 {
//                 return count
//             }
//             for i := 0; i < c; i++ {
//                     buffer[wrIdx] = buf4[i]
//                     wrIdx += 1
//                     total += 1
//             }
//             fmt.Println("buffer:", buffer[:3])
//             if n < c  {
//                 c = n 
//             }

//             if c > 0 {
//                 for i := 0; i < c; i++ {
//                     buf[count] = buffer[rdIdx]
//                     rdIdx += 1
//                     count += 1
//                 }
//                 fmt.Println("buf:", buf)
//             } 

//             if c < 4 {
//                 return count
//             }
//             n = n - c
//         }
//         return count
//     }
// }
