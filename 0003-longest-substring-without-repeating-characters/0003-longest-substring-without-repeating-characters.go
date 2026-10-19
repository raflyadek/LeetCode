func lengthOfLongestSubstring(s string) int {
    result := 0
    temp := 0 
    mapS := make(map[byte]bool)
    left, right := 0, 0
    for right < len(s) {
        if mapS[s[right]] {
            left++
            if temp > result {
                result = temp
            }
            temp = 0
            right= left
            clear(mapS)
        }
        mapS[s[right]] = true
        temp++
        right++
    }
    if temp > result {
        result = temp 
    }
    return result
}