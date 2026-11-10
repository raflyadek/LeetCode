func vowelStrings(words []string, left int, right int) int {
    res := 0 
    mapVowel := map[byte]bool{
        'a': true,
        'i': true,
        'u': true,
        'e': true,
        'o': true,
    }
    for i := left; i <= right; i++ {
        lastIndex := len(words[i])-1
        if mapVowel[words[i][0]] && mapVowel[words[i][lastIndex]] {
            res++
        }
    }
    return res 
}