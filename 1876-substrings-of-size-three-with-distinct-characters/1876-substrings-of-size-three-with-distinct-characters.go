func countGoodSubstrings(s string) int {
    //just fucking around with the value here lol
    left, mid, right := 0, 1, 2
    res := 0
    for right < len(s) {
        if s[left] != s[mid] && s[left] != s[right] && s[mid] != s[right] {
            left++
            mid++
            right++
            res++
        } else {

        left++
            mid++
       right++
        }

    }
    return res
}