/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        //traverse 2 times faster than slow and if fast reached a point where it is nil OR fast.next is nil then it stopped and slow should be in the middle because it traverse once per iteration
        fast = fast.Next.Next
    }

    return slow
}