func findWordsContaining(words []string, x byte) []int {
    res := make([]int, 0)
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words[i]); j++ {
            if words[i][j] == x {
                res = append(res, i)
                break
            }
        }
    }
    return res
}