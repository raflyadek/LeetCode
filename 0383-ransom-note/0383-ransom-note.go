func canConstruct(ransomNote string, magazine string) bool {
    mapRansom := make(map[rune]int)
    for _, v := range ransomNote {
        mapRansom[v]++
    }

    for _, v := range magazine {
        mapRansom[v]--
    }
    for _, v:= range mapRansom {
        if v > 0 {
            return false
        }
    }

    return true
}