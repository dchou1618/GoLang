/* */


package main

import (
    "fmt"
)

type ListNode struct {
    Val int;
    Next *ListNode;
};


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func printLst(head *ListNode) {
    var curr *ListNode = head;
    fmt.Print("{");
    for (curr != nil) {
        fmt.Print(curr.Val, ",");
        curr = curr.Next;
    }
    fmt.Print("}");
}

func partition(head *ListNode, x int) *ListNode {
    if (head == nil) {
        return head;
    } else {
        var curr *ListNode = head;
        var lowerHalf *ListNode = &ListNode{0, nil};
        lowerHalf.Next = curr;
        var seenLess bool = false;
        for (lowerHalf.Next != nil && lowerHalf.Next.Val < x) {
            lowerHalf = lowerHalf.Next;
            seenLess = true;
        }
        var insertNode *ListNode = lowerHalf;
        for (lowerHalf != nil && lowerHalf.Next != nil) {
            if (lowerHalf.Next.Val < x && lowerHalf.Next != head) {
                var temp *ListNode = &ListNode{lowerHalf.Next.Val, nil};
                lowerHalf.Next = lowerHalf.Next.Next;
                if (!seenLess) {
                    temp.Next = head;
                    head = temp;
                    seenLess = true;
                    insertNode = head;
                } else {
                    followTemp := insertNode.Next;
                    insertNode.Next = temp;
                    //fmt.Println("H", insertNode.Val, insertNode.Next.Val);
                    insertNode.Next.Next = followTemp;
                    // printLst(head); fmt.Println(insertNode.Next.Val);
                    // printLst(insertNode);
                    insertNode = insertNode.Next;
                    //fmt.Println(insertNode.Val, insertNode.Next.Val,insertNode.Next.Next.Val);
                }
            } else {
                lowerHalf = lowerHalf.Next;
            }
        }
        return head;
    }
}

func main() {
    fmt.Println();
}
