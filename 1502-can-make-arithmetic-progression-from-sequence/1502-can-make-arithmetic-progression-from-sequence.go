func canMakeArithmeticProgression(arr []int) bool {
    slices.Sort(arr)
    diff := arr[1]-arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] - arr[i-1] != diff {
            return false
        }
    }
    return true
}