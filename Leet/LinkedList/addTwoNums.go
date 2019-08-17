// addTwoNums.go file

package main

import (
  "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    if (head == nil) {
        return head;
    }
    var prevNode *ListNode;
    currNode := head;
    temp := currNode.Next;
    for (currNode != nil) {
        temp = currNode.Next;
        currNode.Next = prevNode;

        prevNode = currNode;
        currNode = temp;
    }
    head = prevNode;
    return head;
}

func lstLength(l1 *ListNode) (int) {
    currNode := l1;
    length := 0;
    for (currNode != nil) {
        length++;
        currNode = currNode.Next;
    }
    return length;
}
// Problem 2: Efficiency 75%, Memory 100%
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (*ListNode){
    carryOver := 0;
    longer := l1;
    shorter := l2;
    if (lstLength(l1) > lstLength(l2)) {
        longer = l1;
        shorter = l2;
    } else {
        longer = l2;
        shorter = l1;
    }
    currNodeLonger := longer;
    currNodeShorter := shorter;
    for (currNodeLonger != nil || currNodeShorter != nil) {
        if (currNodeLonger != nil && currNodeShorter == nil) {
            if (carryOver == 1) {
                currNodeLonger.Val += carryOver;
                if (currNodeLonger.Val >= 10) {
                    currNodeLonger.Val %= 10;
                    carryOver = 1;
                } else {
                    carryOver = 0;
                }
            }
            if (carryOver == 1 && currNodeLonger.Next == nil) {
                currNodeLonger.Next = &ListNode{1, nil};
                break;
            }
            currNodeLonger = currNodeLonger.Next;
        } else {
            currNodeLonger.Val += currNodeShorter.Val + carryOver;
            if (currNodeLonger.Val >= 10) {
                currNodeLonger.Val %= 10;
                carryOver = 1;
            } else {
                carryOver = 0;
            }
            if (carryOver == 1 && currNodeLonger.Next == nil) {
                currNodeLonger.Next = &ListNode{1, nil};
                break;
            }
            currNodeLonger = currNodeLonger.Next;
            currNodeShorter = currNodeShorter.Next;
        }
    }
    return longer;
}

func main() {

}
