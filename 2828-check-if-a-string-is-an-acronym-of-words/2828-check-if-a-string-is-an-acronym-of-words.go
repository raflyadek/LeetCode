func isAcronym(words []string, s string) bool {
    // ptr := 0
    if len(words) != len(s) {
        return false
    }
    for i := 0; i < len(words); i++ {
        if words[i][0] != s[i] {
            return false
        }
    }
    
    return true
}