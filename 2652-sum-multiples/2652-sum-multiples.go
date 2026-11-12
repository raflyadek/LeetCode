func sumOfMultiples(n int) int {
    res :=0 
    for n > 0 {
        if n % 3 == 0 || n % 5 == 0 || n % 7 == 0 {
            res += n
        }
        n--
    }
    return res
}