/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil {
        return false
    }
    //middle
    temp := head
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    //REVERSE THE SLOW
    var prev *ListNode
    for slow != nil {
        next := slow.Next
        slow.Next = prev
        prev = slow
        slow = next
    }
    for temp != nil && prev != nil {
        if temp.Val != prev.Val {
            return false
        }
        temp, prev = temp.Next, prev.Next
    }
    return true
}