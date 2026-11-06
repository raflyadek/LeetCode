func stableMountains(height []int, threshold int) []int {
    res := make([]int, 0)
    for i := 1; i < len(height); i++ {
        if height[i-1] > threshold {
            res = append(res, i)
        } 
    }
    return res
}