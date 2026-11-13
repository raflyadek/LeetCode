func isSameAfterReversals(num int) bool {
    for num > 0 {
        s := num % 10 
        if s == 0 {
            return false
        } else {
            return true
        }
    }

    return true
}