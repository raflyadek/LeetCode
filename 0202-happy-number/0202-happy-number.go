func isHappy(n int) bool {
    mapNum := make(map[int]int)
    for n != 1 {
    sum := 0
        for n > 0 {
            digit := n%10
            sum += digit * digit
            n /= 10  
        }
        n = sum
        mapNum[n]++
        if mapNum[n] > 1 {
            return false
        }
    }
    return true
}