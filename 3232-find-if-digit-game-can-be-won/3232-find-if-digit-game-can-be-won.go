func canAliceWin(nums []int) bool {
    n1 := 0
    n2 := 0
    for _, v := range nums {
        if v >= 10 {
            n2 += v
        } else if v < 10 {
            n1 += v
        }
    }

    if n1 == n2 {
        return false
    }
    return true
}