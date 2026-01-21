func numberOfSteps(num int) int {
    result := 0 
    resultNum := num

    for resultNum > 0 {
        if resultNum % 2 == 0 {
            resultNum /= 2
            result++
        } else if resultNum % 2 != 0 {
            resultNum -= 1
            result++
        }   
    }

    return result
}