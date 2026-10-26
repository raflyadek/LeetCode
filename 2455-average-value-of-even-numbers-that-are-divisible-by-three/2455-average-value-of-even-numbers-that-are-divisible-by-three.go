func averageValue(nums []int) int {
    sum := 0
    count := 0
    for _, v := range nums {
        if v % 2 ==0 && v % 3 == 0{
            sum +=v
            count++
        }
    }
    if count == 0 {
        return 0
    }
    return sum/count
}