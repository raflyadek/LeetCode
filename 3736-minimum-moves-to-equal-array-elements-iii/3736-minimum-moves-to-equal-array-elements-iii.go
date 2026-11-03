func minMoves(nums []int) int {
    highest := nums[0]
    for _, v := range nums {
        if v > highest {
            highest = v
        }
    }
    res := 0
    for _, v := range nums {
        if v < highest {
            res += highest - v
        }
    }
    return res
}