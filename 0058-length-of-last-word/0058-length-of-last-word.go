func lengthOfLastWord(s string) int {
    result := 0 
    isWord := false
    for i := len(s)-1; i >= 0; i-- {
        if string(s[i]) != " " {
            result++
            isWord = true
        } else {
            if isWord == true {
                break
            }
            isWord = false
        }
    }
    return result
}