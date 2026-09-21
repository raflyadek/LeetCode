func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    mapLetter := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        mapLetter[s[i]]++
    }
    for i := 0; i < len(t); i++ {
        mapLetter[t[i]]--
    }

    for _, v := range mapLetter {
        if v > 0 {
            return false
        }
    }

    return true
}