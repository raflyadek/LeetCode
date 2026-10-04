func isPalindrome(s string) bool {
    if len(s) == 0 || len(s) == 1 {
        return true
    }
    strTemp := strings.ToLower(s)
    str := "abcdefghijklmnopqrstuvfwxyz0123456789"
    var str2 string
    for i := 0; i < len(strTemp); i++ {
        if strings.Contains(str, string(strTemp[i])){
            str2 += string(strTemp[i])
        }
    }
    if len(str2) == 0 || len(str2) == 1 {
        return true
    }
    j := len(str2)-1
    for i := 0; i < j; i++ {
        if str2[i] != str2[j-i] {
            return false
        }
    }

    return true
}