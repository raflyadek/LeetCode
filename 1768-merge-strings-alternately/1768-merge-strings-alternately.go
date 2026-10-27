func mergeAlternately(word1 string, word2 string) string {
    count := 0
    str := ""
    if len(word1) > len(word2) {
        for i := 0; i < len(word2); i++ {
            str += string(word1[i])
            str += string(word2[i])
            count++
        }
        str += word1[count:]
    } else {
        for i := 0; i < len(word1); i++ {
            str += string(word1[i])
            str += string(word2[i])
            count++
        }
        str += word2[count:]
    }
    return str
}