func repeatedCharacter(s string) byte {
    mapS := make(map[rune]int)
    for _, v := range s {
        mapS[v]++
        if mapS[v] >=2 {
            return byte(v)
        }
    }
    return '0'
}