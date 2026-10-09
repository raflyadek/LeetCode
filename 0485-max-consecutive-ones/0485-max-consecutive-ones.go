func findMaxConsecutiveOnes(nums []int) int {
    counter := 0 
    temp := 0
    for _, v := range nums {
        if v == 1 {
            temp++
        } else {
            if temp > counter {
                counter = temp
            }
            temp = 0
        }
    }
    if temp > counter { 
        counter = temp 
    }
    return counter
}