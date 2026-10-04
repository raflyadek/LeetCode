func isHappy(n int) bool {
    mapN := make(map[int]bool)
    for n != 1 {
        temp := 0
        for n > 0 {
            //get last digit
            tempMod := n % 10
            temp += tempMod * tempMod
            //remove last digit
            n /= 10
        }
        n = temp
        if mapN[temp] {
            return false
        }
        mapN[temp] = true
    }
    return true
}