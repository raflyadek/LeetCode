func isPrefixOfWord(sentence string, searchWord string) int {
    tempArr := make([]string, 0)
    tempStr := ""
    for i := 0; i < len(sentence); i++ {
        if string(sentence[i]) != " " {
            tempStr += string(sentence[i])
        } else {
            tempArr = append(tempArr, tempStr)
            tempStr = ""
        }
        if i == len(sentence)-1 {
            tempArr = append(tempArr, tempStr)
            tempStr = ""
        }
    }

    //just contains no?? no, because its prefix
    i, j, k := 0, 0, 0
    
    for i < len(tempArr) {
        if len(tempArr[i]) < len(searchWord) {
            i++
            continue
        }
        if tempArr[i][j] == searchWord[k] {
            if k == len(searchWord) - 1 {
                return i + 1
            }
            j++
            k++ 
        } else {
            i++
            j = 0
            k = 0
            continue
        }
    }

    return -1
}