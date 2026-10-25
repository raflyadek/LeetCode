/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    //middle
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    curr := slow
    var prev *ListNode
    //reverse it
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    for prev != nil && head != nil{
        if prev.Val != head.Val {
            return false
        }
        prev = prev.Next
        head = head.Next
    }
    return true
}