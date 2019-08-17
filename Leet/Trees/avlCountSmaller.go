/* AVL Tree: left and right children differ by at most 1 for every node
   height of a node is the maximum distance from the node to a leaf.

   Proving AVL trees are balanced. Worst-Case: right subtree having height
   of one more than the left for every node.

   N_h = minimum number of nodes in AVL for height h
   N_h = 1 + N_{h-1} + N_{h-2}. 1 is for the root, N_{h-1} being the taller
   RIGHT subtree and N_{h-2} being the shorter left subtree.
         o
     /_|  /\
         /_\
   In turn, N_h > F_h = (phi^h)/sqrt(5).
   n > (phi^h)/sqrt(5) => log_{phi}(n) > 1/sqrt(5)(h)
   AVL insert:
   (1) Traverse down tree and insert where key isn't
   (2) Fix AVL
   - walk up tree until heavier side is of difference greater than 1
   - left rotate(x) being that root is moved over to the left.
*/

package main

import (
    "fmt"
)

type AVLNode struct {
    height int;
    data int;
    leftCount int;
    left *AVLNode;
    right *AVLNode;
};

func makeAVLNode(data int) (*AVLNode){
    this := new(AVLNode);
    this.height = 1;
    this.data = data;
    return this;
}

func max(n1 int, n2 int) int {
    if (n1 > n2) {
        return n1;
    } else {
        return n2;
    }
}

func minNode(root *AVLNode) int {
    if (root != nil) {
        for (root.left != nil) {
            root = root.left;
        }
        return root.data;
    } else {
        panic(1);
    }
}

func height(root *AVLNode) int { // height at root node
    if (root != nil) {
        return root.height;
    } else {
        return 0;
    }
}

func getBalanceFactor(root *AVLNode) int {
    if (root == nil) {
        return 0;
    } else {
        return height(root.left)-height(root.right);
    }
}

func leftRotate(root *AVLNode) (*AVLNode){
    var rootRight *AVLNode = root.right;
    var leftSubtree *AVLNode = rootRight.left;
    rootRight.left = root;
    root.right = leftSubtree;
    rootRight.leftCount += 1+root.leftCount;
    // leftTree of root becomes rootRight's leftTree.
    rootRight.height = max(height(rootRight.left), height(rootRight.right)) + 1;
    root.height = max(height(root.left), height(root.right)) + 1;
    return rootRight;
}
/*
          y
      /       \
     x       /_\
 /_\  /T2\
 Shifts y over to the right
           x
       /      \
     /_\      y
            /   \
          /T2\  /_\
and vice versa for left rotations
*/
func rightRotate(root *AVLNode) (*AVLNode){
    var x *AVLNode = root.left;
    var leftX *AVLNode = x.right;
    x.right = root;
    root.left = leftX;
    root.leftCount -= 1+x.leftCount;
    x.height = max(height(x.left),height(x.right))+1;
    root.height = max(height(root.left),height(root.right))+1;
    return x;
}

func insertAVL(root *AVLNode, val int, counter int, counts *[]int) *AVLNode{
    if (root == nil) {
        *counts = append([]int{counter}, *counts...);
        return makeAVLNode(val);
    } else if (root.data == val) {
        root.leftCount++;
        return root;
    } else if (root.data < val) {
        /* current value will have 1 (root) + everything to left of root to
           left of root */
        root.right = insertAVL(root.right, val, counter + root.leftCount + 1, counts);
    } else {
        root.leftCount++; // one more node on the left of root
        root.left = insertAVL(root.left, val, counter, counts);
    }
    root.height = max(height(root.left), height(root.right)) + 1;
    // updating heights
    var balanceFactor = getBalanceFactor(root);
    // balanceFactor defined to be left-right, so left-left (rotate right)
    if (balanceFactor > 1 && val < root.left.data) {
        return rightRotate(root);
    } else if (balanceFactor < -1 && val > root.right.data) {
      // right-right
        return leftRotate(root);
    } else if (balanceFactor > 1 && val > root.left.data) {
      // left-right
        root.left = leftRotate(root.left);
        return rightRotate(root);
    } else if (balanceFactor < -1 && val < root.right.data) {
        // root is right heavy and value belongs on left-side of right
        root.right = rightRotate(root.right);
        return leftRotate(root);
    }
    return root;
}

func countSmaller(nums []int) []int {
    if (len(nums) == 0) {
        return []int{};
    } else {
        var counts []int;
        var root *AVLNode;
        for i := len(nums)-1; i >= 0; i-- {
            root = insertAVL(root, nums[i], 0, &counts);
        }
        return counts;
    }
}


func main() {
    nums := []int{5,2,6,1,-1,-2,0};
    fmt.Println(countSmaller(nums));

}
