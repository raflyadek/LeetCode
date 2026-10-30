func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := make([]int, 0)
    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            merged = append(merged, nums1[i])
            i++
        } else {
            merged = append(merged, nums2[j])
            j++
        }
    }
    for i < len(nums1) {
        merged = append(merged, nums1[i])
        i++
    }
    for j < len(nums2) {
        merged = append(merged, nums2[j])
        j++
    }

    var res float64
    mid := len(merged)/2 
    if len(merged) % 2 == 0 {
        res += (float64(merged[mid]) + float64(merged[mid-1]) )/ 2
    } else {
        res += float64(merged[mid])
    }
    return res
}