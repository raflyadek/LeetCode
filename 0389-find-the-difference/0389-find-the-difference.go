func findTheDifference(s string, t string) byte {
    res := 0
    for i := 0; i < len(s); i++ {
        res = res ^ int(s[i])
    }

    for i := 0; i < len(t); i++ {
        res = res ^ int(t[i])
    }
    return byte(res)
}