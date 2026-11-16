func earliestTime(tasks [][]int) int {
    res := tasks[0][0] + tasks[0][1]

    for i := 0; i < len(tasks); i++ {
        if tasks[i][0] + tasks[i][1] < res {
            res = tasks[i][0] + tasks[i][1]
        }
    }
    return res
}