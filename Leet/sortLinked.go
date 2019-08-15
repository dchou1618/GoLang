// file sortLinked.go

package main

import (
  "fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// using fast/slow nodes to get mid and split

func mergeLsts(node1 *ListNode, node2 *ListNode, temp *ListNode) *ListNode {
    currNode := temp; // always should initialize current node to be referenced to the head.
    for (node1 != nil && node2 != nil) {
        if (node1.Val < node2.Val) {
            currNode.Next = node1;
            node1 = node1.Next;
        } else {
            currNode.Next = node2;
            node2 = node2.Next;
        }
        currNode = currNode.Next;
    }
    if (node1 == nil) {
        currNode.Next = node2;
    }
    if (node2 == nil) {
        currNode.Next = node1;
    }
    return temp;
}

// func getMid(node *ListNode) *ListNode {
//     currNode := node;
//     jumpNode := node;
//     for (jumpNode != nil && jumpNode.Next != nil) {
//         currNode = currNode.Next;
//         jumpNode = jumpNode.Next.Next;
//     }
//     return currNode;
// }

func mergeSplit(head *ListNode, temp *ListNode) *ListNode {
    if (head.Next == nil) {
        return head;
    }
    currNode := head;
    jumpNode := head;
    var mid *ListNode;

    for (jumpNode != nil && jumpNode.Next != nil) {
        mid = currNode;
        currNode = currNode.Next;
        jumpNode = jumpNode.Next.Next;
    }
    mid.Next = nil;
    left := mergeSplit(head, temp);
    // fmt.Println(mid.Val, currNode.Next);
    right := mergeSplit(currNode, temp);
    temp = mergeLsts(left, right, temp);
    return temp.Next;
}

func sortList(head *ListNode) *ListNode {
    if (head == nil) {
        return head;
    } else {
        temp := &ListNode{-1,nil};
        return mergeSplit(head, temp);
    }
}
