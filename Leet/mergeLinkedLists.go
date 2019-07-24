// file mergeLinkedLists.go

package main

import (
  "fmt"
)

type ListNode struct {
  Val int
  Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    currNode1 := l1;
    currNode2 := l2;
    if (currNode1 == nil) {
        return currNode2
    } else if (currNode2 == nil) {
        return currNode1;
    }
    var less *ListNode;
    var more *ListNode;
    if (currNode1.Val > currNode2.Val) {
        less = currNode2;
        more = currNode1;
    } else {
        more = currNode2;
        less = currNode1;
    }
    // if two values in smaller initial list less than curr more.Val,
    // move the pointer in the "less" list.
    // otherwise, for instance,
    // [1,2,4] (less) , [1,3,4] -> [1,1] // [2,4]
    for (less.Next != nil && more != nil) {
        if (less.Val <= more.Val && less.Next.Val <= more.Val) {
            less = less.Next;
        } else if (less.Val <= more.Val) {
            tempMore := more; // first iteration:
            more = less.Next; //more = 2
            less.Next = tempMore; // less.Next = 1, so 1->1 and 2->4
        }
    }
    if (more != nil) {
        less.Next = more; // then more is larger than less in length
    }
    if (l1.Val > l2.Val) {
        return l2;
    } else {
        return l1;
    }
}
