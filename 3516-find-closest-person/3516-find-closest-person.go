func findClosest(x int, y int, z int) int {
    res := x
    stepX := abs(x-z)
    stepY := abs(y-z)
    if stepX > stepY {
        res = 2
    } else if stepX == stepY {
        res = 0
    } else {
        res = 1
    }
    return res    
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}