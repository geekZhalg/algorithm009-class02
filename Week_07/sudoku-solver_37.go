var  nums = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
func dfs(board [][]byte, row int, col int, col_s [9][10]bool, row_s [9][10]bool, box_s [9][10]bool) bool{
    if col == 9 {
        col = 0
        row++

        if row == 9 {
            return true
        }
    }

    if board[row][col] != '.' {
        return dfs(board, row, col + 1, col_s, row_s, box_s)
    } else {
        for i := 1; i <= 9; i++ {
            idx := (row/3)*3 + col/3
            if  col_s[col][i] == true || row_s[row][i] == true || box_s[idx][i] == true {
                continue
            }

            col_s[col][i] = true
            row_s[row][i] = true
            box_s[idx][i] = true
            board[row][col] = nums[i-1]
            if dfs(board, row, col + 1, col_s, row_s, box_s) == true {
                return true
            }

            col_s[col][i] = false
            row_s[row][i] = false
            box_s[idx][i] = false
            board[row][col] = '.'
        }

        return false
        
    }
}

func solveSudoku(board [][]byte)  {
    //
    var col_s [9][10]bool
    var row_s [9][10]bool
    var box_s [9][10]bool

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] != '.' {
                num := board[i][j] - '0'
                col_s[j][num] = true
                row_s[i][num] = true
                idx := (i/3)*3 + j/3
                box_s[idx][num] = true
            }
        }
    }

    dfs(board, 0, 0, col_s, row_s, box_s)
}
