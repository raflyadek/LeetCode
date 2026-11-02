func countStudents(students []int, sandwiches []int) int {
    //0 circular 1 square
    n := len(students)
    temp := n
    stuck := 0
    for n > 0 {
        if students[0] == sandwiches[0] {
            //delete the queue
            stuck = 0
            students = students[1:]
            sandwiches = sandwiches[1:]
            n--
        } else {
            if stuck > temp*2 {
                break
            }
            front := students[0]
            students = students[1:]
            students = append(students, front)
            stuck++
        }
    }
    return len(students)
}