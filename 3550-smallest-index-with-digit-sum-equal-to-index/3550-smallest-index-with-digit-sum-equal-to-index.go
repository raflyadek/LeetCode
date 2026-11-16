func smallestIndex(nums []int) int {
    res := -1
    for i, v := range nums {
        temp := 0
        for v > 0 {
            d := v % 10
            temp += d
            v /= 10
        }
        if temp == i {
            return i
        }
    }
    return res
}