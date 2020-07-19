/* if s[i] == '0' && s[i-1] >= '1' && s[i-1] <= '2'
        dp[i] =  dp[i-2]
     else 
        return 0
*/
func numDecodings(s string) int {
    if s[0] == '0' {
        return 0
    }
    
    dp_i_2 := 1
    dp_i_1 := 1
    dp_i:=1

    for i := 1; i < len(s); i++ {
        if s[i] == '0' {
            if s[i-1] >= '1' && s[i-1] <= '2' {
                dp_i = dp_i_2
            } else {
                return 0
            }        
        } else if s[i-1] == '1' {
            dp_i = dp_i_1 + dp_i_2
        } else if s[i-1] == '2' && s[i] >= '1' && s[i] <= '6' {
            dp_i = dp_i_1 + dp_i_2
        }

        dp_i_2 = dp_i_1
        dp_i_1 = dp_i
    }

    return dp_i
}
