// file nthNode.go

package main

import (
  "fmt"
)
// 0 ms (100%), 2.2 MB (61%)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func lstLength(head *ListNode) int {
    currNode := head;
    length := 0;
    for (currNode != nil) {
        length++;
        currNode = currNode.Next;
    }
    return length;
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if (head == nil) {
        return head;
    } else {
        target := lstLength(head)+1-n;
        nNode := 1;
        currNode := &ListNode{0,nil};
        currNode.Next = head;
        var tempNode *ListNode = currNode;
        for (currNode != nil) {
            if (nNode == target) {
                temp := currNode.Next.Next;
                currNode.Next.Next = nil;
                currNode.Next = temp;
                currNode = currNode.Next;
            } else {
                currNode = currNode.Next;
            }
            nNode++;
        }
        return tempNode.Next;
    }
}
