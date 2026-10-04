func canConstruct(ransomNote string, magazine string) bool {
    mapMag := make(map[byte]int)
    for i := 0; i < len(magazine); i++ {
        mapMag[magazine[i]]++
    }
    for i := 0; i < len(ransomNote); i++ {
        val, ok := mapMag[ransomNote[i]]
        if !ok {
            return false
        }
        if val == 0 {
            return false
        }
        mapMag[ransomNote[i]]--
    }
    return true
}