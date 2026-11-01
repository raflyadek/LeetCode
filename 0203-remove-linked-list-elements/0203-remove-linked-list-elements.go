/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val {
        head = head.Next
    }
    if head == nil {
        return nil
    }
    temp := head
    for temp.Next != nil {
        if temp.Next.Val == val {
            temp.Next = temp.Next.Next
        } else {
            temp = temp.Next
        }
    }
    //return head because we just manipulate head wth another copy
    return head
}