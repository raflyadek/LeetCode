func reversePrefix(word string, ch byte) string {
    words := ""
    gotcha := 0
    for i, v := range word {
        if byte(v) == ch {
            gotcha = i
            break
        }
    }
    if gotcha == 0 {
        return word
    }
    temp := gotcha+1
    for gotcha >= 0 {
        words += string(word[gotcha])
        gotcha--
    }
    words += word[temp:]
    return words
}