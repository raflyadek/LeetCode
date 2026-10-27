func checkIfPangram(sentence string) bool {
    //if map length is not 26 then its false, use rune/string as a key so its not double in map
    mapS := make(map[rune]bool)
    if len(sentence) < 26 {
        return false
    }
    for _, v := range sentence {
        mapS[v] = true
    }
    if len(mapS) != 26 {
        return false
    }
    return true
}