func isAnagram(s string, t string) bool {
    var s_a [26] int

    if len(s) != len(t) {
        return false
    }

    for _,v := range s {
        s_a[v - 'a'] ++
    }

    for _,v := range t {
        s_a[v - 'a'] --
        if s_a[v - 'a'] < 0 {
            return false
        }
    }

    return true
}
