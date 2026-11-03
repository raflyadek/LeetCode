func hammingWeight(n int) int {
    res := 0
    for n > 0 {
        res += n&1
        n = n >> 1
    }
    return res
}