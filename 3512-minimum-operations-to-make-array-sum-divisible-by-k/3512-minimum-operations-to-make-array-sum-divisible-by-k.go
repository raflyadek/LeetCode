func minOperations(nums []int, k int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    temp := sum % k
    if temp == 0 {
        return 0
    }
    return temp
}