// file mergeLinkedLists.go

package main

import (
  "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}


func clump(front *ListNode, max int) *ListNode {
  // concept: remove duplicates of current node examining, then insert duplicates seen at position.
    if (max <= 0) {
        panic(1);
    }
    if (front == nil || front.Next == nil) {
        return;
    } else {
        var connect *ListNode = front;
        // Does run in O(N^2);
        for (connect != nil) {
            var ref *ListNode = connect.Next;
            var curr *ListNode = connect;
            var currCount int = 1;
            for (curr.Next != nil) {
                if (curr.Next.Val == connect.Val) {
                    currCount++;
                    curr.Next = curr.Next.Next;
                } else {
                    curr = curr.Next;
                }
            }
            if (currCount > max) {
                currCount = max-1;
            } else {
                currCount--;
            }
            var temp *ListNode = connect.Next;
            for i := 0; i < currCount; i++ {
                connect.Next = &ListNode{Val:connect.Val};
                connect = connect.Next;
            }
            connect.Next = temp;
            connect = ref;
        }

    }
}


/* Further manipulation with combining duplicates in linked lists */

func printLst(front *ListNode) {
    var curr *ListNode = front;
    fmt.Print("{");
    for (curr != nil) {
        fmt.Print(curr.Val);
        if (curr.Next != nil) {
            fmt.Print(" ");
        }
        curr = curr.Next;
    }
    fmt.Print("}");
}


func combineDuplicates(front *ListNode) *ListNode {
    if (front != nil) {
        if (front != nil && front.Next != nil && front.Val == front.Next.Val) {
            var lastSeen int = front.Next.Val;
            for (front != nil && front.Next != nil && front.Next.Val == lastSeen) {
                var val int = front.Val;
                front = front.Next;
                front.Val += val;
            }
        }
        if (front != nil && front.Next != nil) {
            var curr *ListNode = front;
            var lastSeen int = curr.Next.Val-1;
            for (curr != nil && curr.Next != nil) {
                if (curr.Next.Val == lastSeen) { // during removal, curr->next = curr->next->next
                    var val int = curr.Next.Val;
                    curr.Next = curr.Next.Next;
                    curr.Val += val;
                } else {
                    lastSeen = curr.Next.Val;
                    curr = curr.Next;
                }
            }
        }
    }
    return front;
}

/* interleaving between regular and reversed linked list */
func reverse(front *ListNode) *ListNode {
    var prev *ListNode = nil;
    var newFront *ListNode = &ListNode{Val:0};
    var base *ListNode = newFront;
    var curr *ListNode = front;
    for (curr != nil) {
        newFront.Next = &ListNode{Val:curr.Val};
        newFront = newFront.Next;
        curr = curr.Next;
    }
    curr = base.Next;
    for (curr != nil) {
        var temp *ListNode = curr.Next;
        curr.Next = prev;
        prev = curr;
        curr = temp;
    }
    base = prev;
    return base;
}

func braid(front *ListNode) *ListNode {
    if (front != nil) {
        var start *ListNode = reverse(front);
        var tracer *ListNode = front;
        for (tracer != nil) {
            var temp *ListNode = &ListNode{Val: start.Val};
            var follow *ListNode = tracer.Next;
            tracer.Next = temp;
            temp.Next = follow;
            tracer = follow;
            start = start.Next;
        }
        return front;
    } else {
        return nil;
    }
}



func listLength(front *ListNode) int {
    var curr *ListNode = front;
    var length int = 0;
    for (curr != nil) {
        curr = curr.Next;
        length++;
    }
    return length;
}

func isSortedBy(front *ListNode, k int) bool {
    if (k <= 0) {
        panic(1);
    } else if (front == nil || front.Next == nil || k > listLength(front)) {
        return true;
    } else {
        var iters int = 1;
        for (iters <= k) {
            var start *ListNode = front;
            var follow *ListNode = front;
            for iter := 1; iter < iters; iter++ {
                start = start.Next;
                follow = follow.Next;
            }
            var exitWhile bool = false;
            for (start != nil) {
                for i := 0; i < k; i++ {
                    if (follow.Next != nil) {
                        follow = follow.Next;
                    } else {
                        exitWhile = true;
                    }
                }
                if (exitWhile) {
                    break;
                }
                if (follow.Val < start.Val) {
                    return false;
                }
                for j := 0; j < k; j++ {
                    start = start.Next;
                }
            }
            iters++;
        }
        return true;
    }
}


/* merges values until all values are at least value of N */
func mergeToN(front *ListNode, n int) {
    if (front != nil) {
        return;
    } else {
        for (front.Val < n && front.Next != nil) {
            var val int = front.Val;
            front.Next.Val += val;
            front = front.Next;
        }
        if (front.Val < n) {
            front = front.Next;
            return;
        }
        var curr *ListNode = &ListNode{Val:0};
        curr.Next = front;
        for (curr.Next != nil) { // use temporary new listNode
            if (curr.Next.Val < n) {
                if (curr.Next.Next != nil) {
                    var val int = curr.Next.Val;
                    curr.Next.Next.Val += val;
                    curr.Next = curr.Next.Next;
                } else {
                    curr.Next = curr.Next.Next;
                    return;
                }
            } else {
                curr = curr.Next;
            }
        }
    }
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

func main() {
    n4 := &ListNode{Val: 8, Next: nil};
    n3 := &ListNode{Val: 6, Next: n4};
    n2 := &ListNode{Val: 3, Next: n3};
    n1 := &ListNode{Val: 2, Next: n2};
    front := &ListNode{Val: 1, Next: n1};
    newFront := braid(front);
    printLst(newFront);
}
