func countGoodSubstrings(s string) int {
    left, mid, right := 0, 1, 2
    res := 0
    for right < len(s) {
        if s[left] == s[mid] || s[left] == s[right] || s[mid] == s[right] {
            left++
            mid++
            right++
        } else {
            left++
            mid++
            right++
            res++
        }
    }
    return res
}