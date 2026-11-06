func mirrorDistance(n int) int {
    reverse := 0
    temp := n
    for temp > 0 {
        reverse = (reverse * 10) + (temp%10)
        temp /= 10 
    }
    return abs(n-reverse)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}