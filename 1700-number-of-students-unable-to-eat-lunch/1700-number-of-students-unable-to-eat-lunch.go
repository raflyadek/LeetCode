func countStudents(students []int, sandwiches []int) int {
    //0 circular 1 square
    temp := 0
    // stuck := 0
    for temp < len(students) {
        if students[0] == sandwiches[0] {
            //delete the queue
            // stuck = 0
            students = students[1:]
            sandwiches = sandwiches[1:]
            temp = 0
        } else {
            // if stuck > temp{
            //     break
            // }
            temp++
            front := students[0]
            students = students[1:]
            students = append(students, front)
            // stuck++
        }
    }
    return len(students)
}