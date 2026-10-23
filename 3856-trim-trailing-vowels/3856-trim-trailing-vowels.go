func trimTrailingVowels(s string) string {
    n := len(s)
    res := n
    mapV := map[byte]bool{
        'a': true,
        'i': true,
        'u': true,
        'e': true,
        'o': true,
    }
    for i := n-1; i >= 0; i-- {
        if mapV[s[i]] {
            res--
        } else {
            break
        }
    }
    return s[0:res]
}