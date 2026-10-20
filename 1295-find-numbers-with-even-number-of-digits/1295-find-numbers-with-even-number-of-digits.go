func findNumbers(nums []int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        digit := 0
        num := nums[i]
        for num > 0 {
            num /= 10
            digit++
        }
        if digit % 2 == 0 {
            res++
        }
    }
    return res
}