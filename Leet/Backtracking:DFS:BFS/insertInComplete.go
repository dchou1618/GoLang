/* insertInComplete.go file */
package main

import (
    "fmt"
)

type TreeNode struct {
    data int;
    left *TreeNode;
    right *TreeNode;
};
func max(a int, b int) int {
    if (a > b) {
        return a;
    } else {
        return b;
    }
}

func height(t *TreeNode) int {
    if (t != nil) {
        return 1 + max(height(t.left), height(t.right));
    }
    return 0;
}

/* O(n) sol'n */
func _bfsQueue(t *TreeNode) *TreeNode {
    queue := []*TreeNode{t};
    for (len(queue) != 0) {
        var top *TreeNode = queue[0];
        queue = queue[1:];
        if (top.left == nil || top.right == nil) {
            return top;
        } else {
            queue = append(queue, top.left);
            queue = append(queue, top.right);
        }
    }
    return nil;
}

func _insertToComplete(t *TreeNode) *TreeNode {
    if (t != nil) {
        return _bfsQueue(t);
    } else {
        return nil;
    }
}

func getMin(tree *TreeNode) *TreeNode {
    var curr *TreeNode = tree;
    for (curr.left != nil) {
        curr = curr.left;
    }
    return curr;
}

func getMax(tree *TreeNode) *TreeNode {
    var currNode *TreeNode = tree;
    for (currNode.right != nil) {
        currNode = currNode.right;
    }
    return currNode;
}

/* O(log(n)) order log(n) sol'n */

func _completeInsertion(tree *TreeNode) *TreeNode {
    if (tree == nil || tree.left == nil || tree.right == nil) {
        return tree;
    } else if (height(tree.left) == height(tree.right)) {
        return getMin(tree.left);
    } else if (height(tree.left.left) != height(tree.left.right)) {
        return _completeInsertion(tree.left);
    } else {
        return _completeInsertion(tree.right);
    }
}


func main() {

    root := &TreeNode{data:6};
    root.left = &TreeNode{data:4};
    root.left.left = &TreeNode{data:3};
    root.left.right = &TreeNode{data:5};
    root.right = &TreeNode{data:8};
    fmt.Println(_completeInsertion(root));
}
