func rowAndMaximumOnes(mat [][]int) []int {
    res := make([]int, 0)
    temp := 0
    temp2 := 0
    for i := 0; i < len(mat); i++ {
        temp1 := 0
        for j := 0; j < len(mat[i]); j++ {
            temp1 += mat[i][j]
        }
        if temp1 > temp {
            temp = temp1
            temp2 = i
        }
    }
    res = append(res, temp2, temp)
    return res
}