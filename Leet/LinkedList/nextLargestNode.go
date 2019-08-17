// file

package main
import (
  "fmt"
)

type ListNode struct {
    Val int;
    Next *ListNode;
};

func nextLargerNodes(head *ListNode) []int {
    nextLargest := []int{};
    startNode := head;
    currNode := head;
    for (startNode != nil) {
        largest := startNode.Val;
        for (currNode != nil) {
            if (currNode.Val > largest) {
                largest = currNode.Val;
                break;
            }
            currNode = currNode.Next;
        }
        if (largest == startNode.Val) {
            nextLargest = append(nextLargest,0);
        } else {
            nextLargest = append(nextLargest,largest);
        }
        startNode = startNode.Next;
        currNode = startNode;
    }
    return nextLargest;
}

func main() {

}
