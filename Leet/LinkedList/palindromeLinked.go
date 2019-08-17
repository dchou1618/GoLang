// file palindromeLinked.go


package main

import (
  "fmt"
)

func listLength(head *ListNode) (int) {
    currNode := head;
    l := 0;
    for (currNode != nil) {
        l++;
        currNode = currNode.Next;
    }
    return l;
}
func reverse(head *ListNode) (*ListNode) {
    var prevNode *ListNode;
    currNode := head;
    follow := currNode.Next;
    l := 0;
    halfLen := listLength(head)/2;
    for (l < halfLen) {
        follow = currNode.Next;
        currNode.Next = prevNode;
        prevNode = currNode;
        currNode = follow;
        l++;
    }
    head = prevNode;
    return head;
}

func isPalindrome(head *ListNode) bool {
    if (head == nil || head.Next == nil) {
        return true;
    }
    currNode := head;
    fullLen := listLength(head);
    l := 0;
    for (l < fullLen/2) {
        currNode = currNode.Next;
        l++;
    }
    if (fullLen%2 != 0) {currNode = currNode.Next;}
    tempNode := reverse(head);
    for (tempNode != nil && currNode != nil) {
        if (tempNode.Val != currNode.Val) {
            return false;
        }
        currNode = currNode.Next;
        tempNode = tempNode.Next;
    }
    return true;
}

func main() {
  
}
