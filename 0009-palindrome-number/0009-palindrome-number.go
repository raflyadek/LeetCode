func isPalindrome(x int) bool {
    tempX := x
    res := 0
    for tempX > 0 {
        //
        res = res * 10 + tempX % 10
        //remove the last digit
        tempX /= 10
    }
    return res == x
}