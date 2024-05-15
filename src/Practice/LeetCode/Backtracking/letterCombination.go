// 17. Letter Combinations of a Phone Number
// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// backtracking algorithm

func letterCombinations(digits string) []string {

    if len(digits) == 0 {
        return []string{}
    }
    ans := []string{}
    letters := map[rune]string{ 
        '2':"abc", '3': "def", '4':"ghi", '5':"jkl",
        '6':"mno", '7':"pqrs", '8':"tuv", '9':"wxyz",
    }

    var backtracking func(path []rune, index int) 
    backtracking = func(path []rune, index int) {
        if len(path) == len(digits) { // got one of results
            ans = append(ans, string(path))
            return
        } // get letters for a digit
        allLetters := letters[rune(digits[index])]
        for _, letter := range allLetters {
            path = append(path, rune(letter))
            backtracking(path, index+1) // call recursively here
            path = path[:len(path)-1] // backtrack - remove previous letter
        }
    }

    path := []rune{}
    backtracking(path, 0)
    return ans
    
}
