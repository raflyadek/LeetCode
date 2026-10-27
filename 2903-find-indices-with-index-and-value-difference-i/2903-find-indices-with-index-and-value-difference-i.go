func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    res := make([]int, 2)
    // left, right := 0, 0 
    // for left < len(nums) {
    //     if abs(left-right) >= indexDifference && abs(nums[left]-nums[right]) >= valueDifference {
    //         res[0] = left
    //         res[1] = right
    //         return res
    //     }
    //     right++
    //     if right == len(nums)-1 {
    //         left++
    //         right = left 
    //     }
    // }
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if abs(i-j) >= indexDifference && abs(nums[i]-nums[j]) >= valueDifference {
                res[0]=i
                res[1]=j
                return res
            }
        }
    }
    res[0]=-1
    res[1]=-1
    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}