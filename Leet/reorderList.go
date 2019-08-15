/* */


package main

import (
    "fmt"
)

type ListNode struct {
    Val int;
    Next *ListNode;
};

func reverse(head *ListNode) (*ListNode, int) {
    var newHead *ListNode = &ListNode{Val:head.Val};
    var base *ListNode = newHead;
    var currNode *ListNode = head.Next;
    numNodes := 1;
    for (currNode != nil) {
        numNodes++;
        newHead.Next = &ListNode{Val:currNode.Val};
        currNode = currNode.Next;
        newHead = newHead.Next;
    }
    var prevNode *ListNode;
    currNode = base;
    for (currNode != nil) {
        var temp *ListNode = currNode.Next;
        currNode.Next = prevNode;
        prevNode = currNode;
        currNode = temp;
    }
    newHead = prevNode;
    return newHead, numNodes;
}

func reorderList(head *ListNode)  {
    if (head != nil && head.Next != nil) {
        reverseNode, numNodes := reverse(head);
        var midNode int;
        if (numNodes%2 == 0) {
            midNode = (numNodes-1)/2;
        } else {
            midNode = numNodes/2;
        }
        var curr *ListNode = head;
        var currReverse *ListNode = reverseNode;
        for (midNode > 0) {
            var temp *ListNode = curr.Next;
            curr.Next = &ListNode{Val:currReverse.Val};
            curr.Next.Next = temp;
            currReverse = currReverse.Next;
            curr = curr.Next.Next;
            midNode--;
        }
        if (numNodes%2 != 0) {
            curr.Next = nil;
        } else {
            curr.Next.Next = nil;
        }
    }
}

func main() {

}
