func addedInteger(nums1 []int, nums2 []int) int {
    slices.Sort(nums2)
    slices.Sort(nums1)
    res := 0
    for i := 0; i < len(nums1); i++ {
        if nums1[i] > nums2[i] {
            res = nums2[i]-nums1[i]
            return res
        } else {
            res = nums2[i]-nums1[i]
            return res
        }
    }
    return res
}