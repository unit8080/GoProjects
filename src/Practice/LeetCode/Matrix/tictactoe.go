// 348. Design Tic-Tac-Toe
// https://leetcode.com/problems/design-tic-tac-toe/

type TicTacToe struct {
    rows []int
    cols []int
    diag1 int
    diag2 int
    size int
}


func Constructor(n int) TicTacToe {
    ttt := new(TicTacToe)
    ttt.size = n
    ttt.cols = make([]int, n)
    ttt.rows = make([]int, n)
    return *ttt
}


func (this *TicTacToe) Move(row int, col int, player int) int {
    currPlayer := 1
    if player % 2 == 0 {
        currPlayer = -1
    }

    this.rows[row] += currPlayer
    this.cols[col] += currPlayer
    if row == col {
        this.diag1 += currPlayer
    }
    if col == this.size - row - 1 {
        this.diag2 += currPlayer
    }

    n := this.size
    if player % 2 == 0 {
        n = -1 * this.size
    }
    if this.rows[row] == n ||
        this.cols[col] == n ||
        this.diag1 == n ||
        this.diag2 == n {
            return player
    }
    return 0
}


/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
