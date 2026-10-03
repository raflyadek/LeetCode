func romanToInt(s string) int {
    result := 0
    mapToInt := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    if len(s) == 1 {
        result += mapToInt[s[0]]
    }
    for i := 1; i < len(s); i++ {
        if mapToInt[s[i]] > mapToInt[s[i-1]] {
            result += mapToInt[s[i]] - mapToInt[s[i-1]]
            i++
            if i == len(s)-1 {
                result += mapToInt[s[i]]
            }
        } else {
            result += mapToInt[s[i-1]]
            if i == len(s)-1 {
                result += mapToInt[s[i]]
            }
        }
    }
    return result
}