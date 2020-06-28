func minPathSum(grid [][]int) int {
    row := len(grid)
    if row == 0 {
        return 0
    }
    
    col := len(grid[0])
    if col == 0 {
        return 0
    }
    
    var dp [][]int
    for i := 0; i < row; i++ {
        tmp := make([]int, col)
        dp = append(dp, tmp)
    }
    
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if i == 0 && j == 0 {
                dp[i][j] = grid[i][j]
            } else if i == 0 {
                dp[i][j] = dp[i][j-1] + grid[i][j]
            } else if j == 0 {
                dp[i][j] = dp[i-1][j] + grid[i][j]
            } else {
                dp[i][j] = minV(dp[i-1][j], dp[i][j-1]) + grid[i][j]
            }
        }
    }
    
    return dp[row-1][col-1]
}

func minV(i, j int) int {
    if i > j {
        return j
    }
    return i
}

