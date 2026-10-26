/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
    if head == nil {
        return 0
    }
    result := 0
    mul := 1
    //reverse first?
    curr := head
    var prev *ListNode
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    for prev != nil {
        t := prev.Val * mul
        result += t
        mul *= 2
        prev = prev.Next
    }
    return result
}