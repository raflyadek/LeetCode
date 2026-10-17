func maxProfit(prices []int) int {
    left, right := prices[0], prices[0]
    sum := 0
    for _, v := range prices {
        //we have to include right because, right must be greater than left in terms of index
        if v < left {
            left = v
            right = v
        } else if v > right {
            right = v
        }

        //sum it up
        temp := right - left 
        if temp > sum {
            sum = temp
        }
    }
    return sum
}