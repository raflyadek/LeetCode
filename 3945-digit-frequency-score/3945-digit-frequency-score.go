func digitFrequencyScore(n int) int {
    mapN := make(map[int]int)
    for n > 0 {
        d := n % 10 
        mapN[d]++
        n /= 10
    }
    res := 0
    for i, v := range mapN {
        res += i * v
    }
    return res
}