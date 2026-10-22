/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dummy := head
    for dummy.Next != nil {
        if dummy.Val == dummy.Next.Val {
            dummy.Next = dummy.Next.Next
        } else {
            dummy = dummy.Next
        }
    }
    return head
}