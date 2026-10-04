func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    temp1 := ""
    temp2 := ""
    for i := 0; i < len(word1); i++ {
        temp1 += string(word1[i])
    }

    for i := 0; i < len(word2); i++ {
        temp2 += string(word2[i])
    }
    return temp1 == temp2
}