func isAcronym(words []string, s string) bool {
    // ptr := 0
    if len(words) != len(s) {
        return false
    }
    t := 0
    for _, v := range words {
        if v[0] != s[t] {
            return false
        }
        t++
    }
    
    return true
}