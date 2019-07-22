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
