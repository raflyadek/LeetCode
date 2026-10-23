func maxProduct(nums []int) int {
    left, right := 0, 0
    for _, v := range nums {
        //if right change, then before that left = right\
        //so left will always behind right (smaller)
        if v > right {
            left = right 
            right = v
            //lets assume 3, 4, 5, 8, 6
            //when right hit 6 its not changing but left is change to 6
        } else if v > left {
            left = v
        }
    }
    return (left-1) * (right-1)
}