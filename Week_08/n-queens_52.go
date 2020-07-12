var cnt int

func dfs(row int, n int, col int, pie int, na int) {
    if row >= n {
        cnt++
        return
    }

    p := ((1<<n) - 1) & (^(col | pie | na))
    for p > 0 {
        l := p & -p
        p = p & (p-1)
        dfs(row + 1, n, col|l, (pie|l) << 1, (na|l) >> 1)
    }
}

func totalNQueens(n int) int {
    cnt = 0
    //p := (1 << n) - 1
    dfs(0, n, 0, 0, 0)
    return cnt
}
