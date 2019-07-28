// file removeDuplicates.go

package main

import (
  "fmt"
  "math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if (head == nil || head.Next == nil) {
        return head;
    } else {
        currNode := head;
        var duplicate int = math.MaxInt64;
        var tempNode *ListNode = &ListNode{0,nil};
        first := tempNode;
        for (currNode != nil) {
            if (currNode.Next != nil && currNode.Val == currNode.Next.Val) {
                duplicate = currNode.Val;
            }
            // accounting for nonduplicate nodes
            if (currNode.Val != duplicate) {
                val := &ListNode{currNode.Val,nil};
                tempNode.Next = val;
                tempNode = tempNode.Next;
                val = nil;
            }
            currNode = currNode.Next;
        }

        return first.Next;
    }
}

func main() {
  
}
