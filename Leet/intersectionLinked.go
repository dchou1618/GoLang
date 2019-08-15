// file intersectionLinked.go


// Iterative solution

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if (headA == nil || headB == nil) {
        return nil;
    }
    // simply iterate over the linked list el'ts until two items match up
    currNodeA := headA;
    currNodeB := headB;

    for (currNodeA != currNodeB) {
        if (currNodeA != nil) {
            currNodeA = currNodeA.Next;
        } else {
            currNodeA = headA;
        }
        if (currNodeB != nil) {
            currNodeB = currNodeB.Next;
        } else {
            currNodeB = headB;
        }
    }
    return currNodeB;
}

// Recursive solution
func getIntersect(headA, headB *ListNode) (*ListNode) {
    if (headA != nil && headB != nil) {
        if (headA.Next == headB.Next && headA != headB) {
            return headA.Next;
        } else {
            l1 := getIntersect(headA.Next,headB);
            l2 := getIntersect(headA,headB.Next);
            l3 := getIntersect(headA.Next,headB.Next);
            if (l1 != nil) {
                return l1;
            } else if (l2 != nil) {
                return l2;
            } else {
                return l3;
            }
        }
    }
    return nil;
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if (headA == nil || headB == nil) {
        return nil;
    } else if (headA == headB) {
        return headA;
    } else if (headA == headB.Next) {
        return headA;
    } else if (headB == headA.Next) {
        return headB;
    } else {
        return getIntersect(headA,headB);
    }
}

// push node to delete to the end of the linked list
func deleteNode(node *ListNode) {
    currNode := node;
    for (currNode != nil && currNode.Next != nil) {
        var val int = currNode.Next.Val;
        currNode.Next.Val = currNode.Val;
        currNode.Val = val;
        if (currNode.Next.Next == nil) {
            currNode.Next = currNode.Next.Next;
            break;
        }
        currNode = currNode.Next;
    }
}

func swapPairs(head *ListNode) *ListNode {
    if (head == nil || head.Next == nil) {
        return head;
    }
    var prev *ListNode = head;
    var curr *ListNode = head.Next;
    var temp *ListNode;
    for (curr != nil) {
        if (curr != nil && curr.Next != nil && curr.Next.Next != nil) {
            temp = curr.Next.Next;
        } else {
            temp = curr.Next;
        }
        if (curr == head.Next) {
            head = curr;
        }
        var nextPrev *ListNode = curr.Next;
        curr.Next = prev;
        prev.Next = temp;
        prev = nextPrev;
        curr = temp;
    }
    return head;
}


// rotate linked list

func lstLength(head *ListNode) int {
    var length int = 0;
    var currNode *ListNode = head;

    for (currNode != nil) {
        currNode = currNode.Next;
        length++;
    }
    return length;
}
func rotateRight(head *ListNode, k int) *ListNode {
    if (k == 0 || head == nil) {
        return head;
    }
    var length int = lstLength(head);
    if (k%length == 0) {
        return head;
    }
    k %= length;
    var currNode *ListNode = head;
    var rotateStart *ListNode;
    var currNum int = 0;
    for (currNode != nil) {
        if (currNum+1 == length-k) {
            rotateStart = currNode.Next;
            var start *ListNode = rotateStart;
            currNode.Next = nil;
            for (rotateStart != nil && rotateStart.Next != nil) {
                rotateStart = rotateStart.Next;
            }
            rotateStart.Next = head;
            head = start;
            break;
        }
        currNum++;
        currNode = currNode.Next;
    }
    return head;
}
