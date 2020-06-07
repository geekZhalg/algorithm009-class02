func mycombine(n int, k int, from int,res[]int, result *[][]int) {
    if len(res) == k {
        tmp:=make([]int, k)
        copy(tmp, res)
        (*result) = append(*result, tmp)
        return
    }

    for i:=from; i<= n - (k - len(res)) + 1; i++ {
        res = append(res, i)
        mycombine(n, k, i+1,res, result)
        res = res[:len(res) - 1]
    }

    return
} 

func combine(n int, k int) [][]int {
   result := [][]int{}

   mycombine(n, k, 1, []int{}, &result)

   return result
}
