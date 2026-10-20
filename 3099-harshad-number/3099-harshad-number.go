func sumOfTheDigitsOfHarshadNumber(x int) int {
    //if below 10 then it must be harshad 
    //because its divise by the number itself

    res := 0
    temp := x
    for x > 0 {
        s := x % 10
        res += s
        x /= 10
    }
    if temp % res == 0 {
        return res
    }

    return -1
}