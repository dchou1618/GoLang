/* cycleLinkedList.go file uses basic "runners" to find cycles
   in linked lists.
*/

package main

import (
    "fmt"
    "testing"
)

type LinkedNode struct {
    data int;
    next *LinkedNode;
};

func _existsCycle(head *LinkedNode) bool {
    if (head != nil) {
        var curr *LinkedNode = head;
        var currNext *LinkedNode = head.next;
        for (curr != currNext) {
            if (curr == nil || currNext == nil) {
                // reached end of one linked list (no cycle);
                return false;
            } else {
                curr = curr.next;
                currNext = currNext.next.next;
            }
        }
        return true;
    } else {
        return false;
    }
}

func testExistsCycle(t *testing.T) {
    head := &LinkedNode{data:3, next: &LinkedNode{data:4,next: &LinkedNode{data:6, next: nil}}};
    var exists bool = _existsCycle(head);
    if (exists) {
        t.Errorf("Existence invalid: Got %v, Want %v", exists, false);
    }
    fmt.Println("Done");
}

func main() {}
