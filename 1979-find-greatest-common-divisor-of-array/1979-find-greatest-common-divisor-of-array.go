func findGCD(nums []int) int {
    small := nums[0]
    large := nums[0]
    for _, v := range nums {
        if v < small {
            small = v
        }
        if v > large {
            large = v
        }
    }
    if large % small == 0 {
        return small
    }
    i := 2
    temp := 0
    for i < small {
        if small % i == 0 && large % i == 0 {
            temp = i
        }
        i++
    }
    if temp > 0 {
        return temp
    }
    return 1
}