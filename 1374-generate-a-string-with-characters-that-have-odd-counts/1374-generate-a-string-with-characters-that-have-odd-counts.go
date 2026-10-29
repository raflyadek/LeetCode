func generateTheString(n int) string {
    res := ""
    if n % 2 == 0 {
        //if even then just divide by two and use that number to fill up the res
        //and find if its still even then if its -1
        //nah just -1, its enough no???
        temp := n-1
        res += "a"
        for i := 0; i < temp; i++ {
            res += "b"
        }
    } else {
        for i := 0; i< n; i++{
            res += "a"
        }
    }

    return res
}