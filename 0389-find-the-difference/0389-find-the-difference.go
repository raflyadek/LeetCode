func findTheDifference(s string, t string) byte {
    if len(s) == 0 {
        return t[0]
    }
    mapByteS := make(map[byte]int)
    mapByteT := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        if i < len(t) - 1{
            mapByteT[t[i]]++
            mapByteS[s[i]]++
            continue
        }
        mapByteT[t[i]]++
    }

    for k, _ := range mapByteT {
        if _, ok := mapByteS[k]; !ok {
            return k
        }
        if mapByteS[k] != mapByteT[k] {
            return k
        }
    }
    return s[0]
}