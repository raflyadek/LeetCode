func findPermutationDifference(s string, t string) int {
    res := 0
    //map to store index and char
    mapS := make(map[rune]int)
    for i, v := range s {
        mapS[v] = i
    }
    for i, v := range t {
        res += abs(mapS[v] - i)
    }
    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}