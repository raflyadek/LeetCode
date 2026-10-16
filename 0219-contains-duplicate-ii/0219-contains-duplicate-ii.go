func containsNearbyDuplicate(nums []int, k int) bool {
    mapNum := make(map[int]int)
    for key, v := range nums {
        if keyz, found := mapNum[v]; found {
            if abs(keyz-key) <= k {
                return true
            } 
        }
        mapNum[v] = key
        
    }

    return false
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}