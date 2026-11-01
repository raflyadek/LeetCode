func fib(n int) int {
    a, b := 0, 1
    //b will be a number before a
    for n > 0 {
        a, b = a + b, a
        n--
    }

    return a
}