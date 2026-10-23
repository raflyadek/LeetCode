func countCompleteDayPairs(hours []int) int {
    res := 0
    left, right := 0, 1
    for left < len(hours)-1 {
        
        if (hours[left] + hours[right]) % 24 == 0 {
            res++
        }
        right++ 
        if right == len(hours) {
            left++ 
            right = left+1
        }
    }
    return res
}