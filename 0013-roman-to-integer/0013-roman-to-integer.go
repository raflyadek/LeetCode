func romanToInt(s string) int {
    result := 0
    mapToInt := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000,
    }
    lastIndex := len(s)-1
    for i := 0; i < len(s); i++ {
        if i == lastIndex {
            result += mapToInt[string(s[i])]
            continue
        }
        if mapToInt[string(s[i+1])] > mapToInt[string(s[i])] {
            result += mapToInt[string(s[i+1])] - mapToInt[string(s[i])]
            i += 1
        } else {
            result += mapToInt[string(s[i])]
        }
    }
    return result
}