func numOfStrings(patterns []string, word string) int {
    res := 0
    for _, v := range patterns {
        if strings.Contains(word, v) {
            res++
        }
    }
    return res
}