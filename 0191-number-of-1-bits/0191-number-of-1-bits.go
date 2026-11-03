func hammingWeight(n int) int {
    res := 0
    for n > 0 {
        //&1 as a way to find if the first binary is 1 or no and then shift 1 right to erase the first binary, because 1 is 0001 <- then if the last binary is 1 too then the result is 1
        res += n&1
        n >>= 1
    }
    return res
}