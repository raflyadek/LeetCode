func hammingWeight(n int) int {
    res := 0
    for n > 0 {
        //&1 because 1 is 0001 <- then it acts as a checking
        //how many 1, and if both 1 then add to res += 1, and shift right 1/delete the first binary 
        res += n&1
        n = n >> 1
    }
    return res
}