/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    res := 0
    mul := 1
    //reverse first
    var prev *ListNode 
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    for prev != nil {
        res += prev.Val * mul
        mul *= 2
        prev = prev.Next
    }

    return res
}