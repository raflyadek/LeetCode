func smallerNumbersThanCurrent(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    for i := 0; i < n; i++ {
    temp := 0
        for j := 1; j < n; j++ {
            if nums[i] > nums[(i+j)%n] {
                temp++
            }
        }
    res[i] = temp
    }
    return res
}