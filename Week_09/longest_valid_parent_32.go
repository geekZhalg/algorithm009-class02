/*
func isValidParentheses(s []byte) bool {
    stack := []byte{}
    
    
    for _,c := range s {
        if c == ')' && len(stack) > 0 && stack[len(stack) - 1] == '(' {
            stack = stack[:len(stack) - 1]
        } else {
            stack = append(stack, c)
        }
    }
    //fmt.Println(s,stack, len(stack))
    return len(stack) == 0
}

func longestValidParentheses(s string) int {
    if len(s) < 2 {
        return 0
    }

    //fmt.Println(s)
    maxlen := 0
    ss := []byte(s)
    for i := 2; i <= len(ss); i+=2 {
        for j := 0; j <= len(ss) - i; j++ {
            if isValidParentheses(ss[j:j+i]) == true {
                if i > maxlen {
                    maxlen = i
                }
                break
            }
        }
    } 

    return maxlen
} */

// if s[i] == ')' && s[i-dp[i-1] - 1] == '(' dp[i] = 2 + dp[i-1] + dp[i-dp[i-1] - 2]
/*
func longestValidParentheses(s string) int {
    dp := make([]int, len(s))

    for i := 0; i < len(dp); i++ {
        dp[i] = 0
    }

    for i:=1; i < len(s); i++ {
        if s[i] == ')' && i-dp[i-1] - 1 >= 0 && s[i-dp[i-1] - 1] == '(' {
            dp[i] = 2 + dp[i-1]
            if i - dp[i-1] - 2 >= 0 {
                dp[i] += dp[i-dp[i-1]-2]
            }
        }
    }

    maxlen:=0
    for i := 1; i < len(dp); i++ {
        if dp[i] > maxlen {
            maxlen = dp[i]
        }
    }

    return maxlen
}*/



