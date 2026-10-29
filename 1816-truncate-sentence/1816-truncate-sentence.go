func truncateSentence(s string, k int) string {
    res := ""
    count := 0
    for i := 0; i < len(s); i++ {
        if count == k {
            break
        }
        if string(s[i]) == " " {
            count++
            if count == k {
                break
            }
        }
        res += string(s[i])
    }
    return res
}