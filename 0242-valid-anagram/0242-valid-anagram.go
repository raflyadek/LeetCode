func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    mapStringS := make(map[byte]int)
    mapStringT := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        mapStringS[s[i]]++
        mapStringT[t[i]]++
    }

    for k, v := range mapStringS {
        if v != mapStringT[k] {
            return false
        }
    }
    return true
}