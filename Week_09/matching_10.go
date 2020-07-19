
var memo map[string]bool

func dp(s string, i int, p string, j int) bool{
    if j == len(p) {
        return i == len(s)
    } 

    if i == len(s) {
        if (len(p) - j) % 2 == 1 {
            return false
        }

        for ; j + 1 < len(p); j+=2 {
            if p[j+1] != '*' {
                return false
            }
        }

        return true
    }

    key := strconv.Itoa(i) + strconv.Itoa(j)
    if _,ok := memo[key]; ok {
        return memo[key]
    }

    res := false
    if s[i] == p[j] || p[j] == '.' {
        if j + 1 < len(p) && p[j+1] == '*' {
            res = dp(s, i, p, j + 2) || dp(s, i + 1, p,j)
        } else {
            res = dp(s, i+1, p, j+1)
        }
    } else {
        if j + 1 < len(p) && p[j+1] == '*' {
            res = dp(s, i, p, j + 2)
        } else {
            res = false
        }
    }

    memo[key] = res
    return res
}

func isMatch(s string, p string) bool {
    memo = make(map[string]bool)
    return dp(s, 0, p, 0)
}
