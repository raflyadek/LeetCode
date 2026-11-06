func duplicateNumbersXOR(nums []int) int {
    res := make(map[int]int)
    for _, v := range nums {
        res[v]++
    }
    res1 := 0
    for i, v := range res {
        if v == 2 {
            res1 ^= i
        }
    }
    
    return res1
}