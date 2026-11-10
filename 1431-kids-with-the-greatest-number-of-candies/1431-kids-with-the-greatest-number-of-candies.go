func kidsWithCandies(candies []int, extraCandies int) []bool {
    res := make([]bool, len(candies))
    highest := 0
    for _, v := range candies {
        if v > highest {
            highest = v
        }
    }
    for i := 0; i < len(candies); i++ {
        if candies[i] + extraCandies >= highest {
            res[i] = true
        } else {
            res[i] = false
        }
    }
    return res
}