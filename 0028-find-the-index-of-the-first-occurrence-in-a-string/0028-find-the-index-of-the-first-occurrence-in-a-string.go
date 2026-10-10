func strStr(haystack string, needle string) int {
    if len(haystack) == 1 {
        return 0
    }
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