func majorityElement(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    mid := len(nums)/2
    counter := 0
    slices.Sort(nums)
    for i := 0; i < len(nums); i++ {
        counter++
        if nums[i+1] != nums[i] {
            counter = 0
        }
        if counter >= mid {
            return nums[i]
        }
    }
    return 0
}