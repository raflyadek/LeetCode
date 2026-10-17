func maxProfit(prices []int) int {
    left, right := prices[0], prices[0]
    sum := 0
    for _, v := range prices {
        //check the val 
        if v < left {
            left = v
            right = v
        } else if v > right {
            right = v
        }
        temp := right - left
        if temp > sum {
            sum = temp
        }
    }
    return sum
}