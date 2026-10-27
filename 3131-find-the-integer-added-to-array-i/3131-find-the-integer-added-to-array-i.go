func addedInteger(nums1 []int, nums2 []int) int {
    min1, min2 := nums1[0], nums2[0]
    for i := 1; i < len(nums1); i++ {
        if nums1[i] < min1 {
            min1=nums1[i]
        }
        if nums2[i] < min2 {
            min2=nums2[i]
        }
    }
    res := min2-min1
    return res
}