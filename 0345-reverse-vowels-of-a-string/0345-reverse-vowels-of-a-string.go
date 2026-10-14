func reverseVowels(s string) string {
    //convert to rune so we can modify it
    sRune := []rune(s)
    vowelMap := map[rune]bool{
        'a': true,
        'i': true,
        'u': true,
        'e': true,
        'o': true,
        'A': true,
        'I': true,
        'U': true,
        'E': true,
        'O': true,
    }
    left, right := 0, len(sRune)-1
    for left < right {
        if !vowelMap[sRune[left]] {
            left++
        }
        if !vowelMap[sRune[right]] {
            right--
        }
        if vowelMap[sRune[left]] && vowelMap[sRune[right]] {
            sRune[left], sRune[right] = sRune[right], sRune[left]
            right--
            left++
        }
    }
    return string(sRune)
}