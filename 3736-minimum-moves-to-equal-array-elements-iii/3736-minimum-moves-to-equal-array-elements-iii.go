func minMoves(nums []int) int {
    highest := slices.Max(nums)
    res := 0
    for _, v := range nums {
        if v < highest {
            res += highest - v
        }
    }
    return res
}