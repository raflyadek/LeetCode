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
    copy := head
    //becaue we can only manipulate the next then the condition should be fufilled these 2
    for copy != nil && copy.Next != nil {
        if copy.Next.Val == val {
            copy.Next = copy.Next.Next
        } else {
            copy = copy.Next
        }
    }
    //return head because we just manipulate head wth another copy
    return head
}