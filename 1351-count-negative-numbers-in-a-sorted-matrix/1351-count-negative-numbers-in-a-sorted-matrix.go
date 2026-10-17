func countNegatives(grid [][]int) int {
    result := 0
    for lastIndexRow := len(grid)-1; lastIndexRow >= 0; lastIndexRow-- {
        for lastIndexCol := len(grid[lastIndexRow])-1; lastIndexCol >= 0; lastIndexCol-- {
            if grid[lastIndexRow][lastIndexCol] >= 0 {
                break
            } else {
                result++
            }
        }
    }
    return result
}