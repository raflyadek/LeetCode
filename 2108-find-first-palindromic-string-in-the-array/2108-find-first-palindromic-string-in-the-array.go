func firstPalindrome(words []string) string {
    flag := false
    for i := 0; i < len(words); i++ {
        lastIndex := len(words[i])-1
        if len(words[i]) == 1 {
            return string(words[i])
        }
        for j := 0; j < len(words[i])/2; j++ {
            if words[i][j] != words[i][lastIndex] {
                flag = false
                break
            } else {
                flag = true
                lastIndex--
            }
        }
        if flag == true {
            return string(words[i])
        }
    }
    return ""
}