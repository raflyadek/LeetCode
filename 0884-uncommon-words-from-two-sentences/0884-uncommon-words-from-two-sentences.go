func uncommonFromSentences(s1 string, s2 string) []string {
    res := make([]string, 0, 0)
    mapStr := make(map[string]int)
    tempStr := ""
    for i := 0; i < len(s1); i++ {
        if string(s1[i]) != " " {
            tempStr += string(s1[i])
        } else {
            mapStr[tempStr]++
            tempStr = ""
        }
        if i == len(s1)-1 {
            mapStr[tempStr]++
            tempStr = ""
        }
    }

    for j := 0; j < len(s2); j++ {
        if string(s2[j]) != " " {
            tempStr += string(s2[j])
        } else {
            mapStr[tempStr]++
            tempStr = ""
        }
        if j == len(s2)-1 {
            mapStr[tempStr]++
            tempStr = ""
        }
    }

    for k, v := range mapStr {
        if v == 1 {
            res = append(res, k)
        }
    }
    return res
}