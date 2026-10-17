func firstPalindrome(words []string) string {
    flag := false

    for i := 0; i < len(words); i++ {
        l := len(words[i])-1
        if len(words[i]) == 1 {
            return words[i]
        }
        for j := 0; j < len(words[i])/2; j++ {
            if words[i][j] != words[i][l] {
                break
            } 
            if j == len(words[i])/2 - 1 {
                flag = true
            }
            l--
        }
        if flag == true {
            return words[i]
        }
    }
    return ""
}