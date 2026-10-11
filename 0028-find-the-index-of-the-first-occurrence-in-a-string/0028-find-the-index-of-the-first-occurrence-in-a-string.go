func strStr(haystack string, needle string) int {
    //edge case
    if len(haystack) == 0 && haystack == needle {
        return 0
    }

    //variable to point out two index
    left := 0
    for right := len(needle); right <= len(haystack); right++ {
        if haystack[left:right] == needle {
            return left
        } else {
            left++
        }
    }
    return -1
}