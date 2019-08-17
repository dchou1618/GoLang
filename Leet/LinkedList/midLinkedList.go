// file midLinkedList.go

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// middle node 0 ms (100%), 2mb (59.38%).
func middleNode(head *ListNode) *ListNode {
    var currNode *ListNode = head;
    var jumpNode *ListNode = head;
    for (jumpNode != nil && jumpNode.Next != nil) {
        currNode = currNode.Next;
        jumpNode = jumpNode.Next.Next;
    }
    return currNode;
}
