// file splitLinkedList.go

package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func lstLength(root *ListNode) int {
    var currNode *ListNode = root;
    var length int = 0;
    for (currNode != nil) {
        length++;
        currNode = currNode.Next;
    }
    return length;
}
func splitListToParts(root *ListNode, k int) []*ListNode {
    parts := make([]*ListNode, k);
    for i := range parts {
        parts[i] = &ListNode{Val: 0, Next: nil};
    }
    length := lstLength(root);
    size := length/k;
    larger := length%k;
    var currLength int = 0;
    var index int = 0;
    var currNode *ListNode = root;
    var temp *ListNode = parts[index];
    for (currNode != nil) {
        if (index < larger) {
            if (currLength < size+1) {
                temp.Next = &ListNode{Val: currNode.Val};
                temp = temp.Next;
                currLength++;
            }
            if (currLength == size+1) {
                parts[index] = parts[index].Next;
                currLength = 0;
                index++;
                if (index < k) {
                    temp = parts[index];
                }
                currNode = currNode.Next;
                continue;
            }
        }
        if (index < k && index >= larger) {
            if (currLength < size) {
                temp.Next = &ListNode{Val: currNode.Val};
                temp = temp.Next;
                currLength++;
            }
            if (currLength == size) {
                parts[index] = parts[index].Next;
                currLength = 0;
                index++;
                if (index < k) {
                    temp = parts[index];
                }
            }
        }
        currNode = currNode.Next;
    }
    if (index < k) {
        for (index < k) {
            parts[index] = parts[index].Next;
            index++;
        }
    }
    return parts;
}

func main() {
    node8 := &ListNode{};
    node7 := &ListNode{Val: 4, Next: node8}
    node6 := &ListNode{Val: 12, Next: node7}
    node5 := &ListNode{Val: 10, Next: node6}
    node4 := &ListNode{Val: 8, Next: node5}
    node3 := &ListNode{Val: 2, Next: node4}
    node2 := &ListNode{Val: 6, Next: node3}
    node1 := &ListNode{Val: 2, Next: node2}
    front := &ListNode{Val: 3, Next: node1};
    fmt.Println(splitListToParts(front, 3));
}
