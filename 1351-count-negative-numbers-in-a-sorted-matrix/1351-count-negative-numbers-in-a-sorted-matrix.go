func countNegatives(grid [][]int) int {
    result := 0
    for i := 0; i < len(grid); i++ {
        for j := len(grid[i])-1 ; j >= 0; j-- {
            if grid[i][j] >= 0 {
                break
            } else {
                result++
            }
        }
    }
    return result
}