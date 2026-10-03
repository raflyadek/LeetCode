func majorityElement(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    mapN := make(map[int]int)
    mid := len(nums)/2
    for i := 0; i < len(nums); i++ {
        mapN[nums[i]]++
        if mapN[nums[i]] >= mid+1 {
            return nums[i]
        }
    }
    return 0
}