func firstMatchingIndex(s string) int {
    res := -1
    n := len(s)
    for i := 0; i < len(s); i++ {
        if s[i] == s[n-i-1] {
            return i
        }
    }
    return res
}