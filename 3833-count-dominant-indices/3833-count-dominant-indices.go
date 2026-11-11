func dominantIndices(nums []int) int {
    res := 0
    for i := 0; i < len(nums); i++ {
        average := 0
        leng := 0
        for j := i+1; j < len(nums); j++ {
            average += nums[j]
            leng++
            if j == len(nums)-1 && nums[i] > (average/leng) {
            res++
        }
        }
    }
    return res
}