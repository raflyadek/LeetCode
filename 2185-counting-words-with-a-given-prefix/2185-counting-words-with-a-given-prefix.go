func prefixCount(words []string, pref string) int {
    n := len(pref)
    res := 0 
    for _, v := range words {
        if len(v) < n {
            continue
        }
        if v[0:n] == pref {
            res++
        }
    }
    return res
}