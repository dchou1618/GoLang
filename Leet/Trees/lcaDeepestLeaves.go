/* file lcaDeepestLeaves.go */

package main

import (
    "fmt"
)

type TreeNode struct {
    data int;
    left *TreeNode;
    right *TreeNode;
};

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
    spin off of the LCA of binary tree (postorder traversal in nature)
    [1,2,null,3,4,null,null,5] - why failed?
             1
            /
           2
          / \
         3   4
        /
       5
    initially took two lowest (one depth may be second lowest)
    Will lead to false answer of 2, 3, 4, 5
 */

func max(a int, b int) int{
    if (a > b) {
        return a;
    } else {
        return b;
    }
}

func getDepth(root *TreeNode) int {
    if (root != nil) {
        return 1 + max(getDepth(root.left), getDepth(root.right));
    }
    return 0;
}


func _lcaLeaves(root *TreeNode, depth, deepest int, output **TreeNode) bool {
    if (root != nil) {
        var seenDeepest bool = false; // if the root is a deepest leaf (one of them);
        if (root.left == nil && root.right == nil && (depth == deepest)) {
            seenDeepest = true;
        }
        left := _lcaLeaves(root.left, depth+1, deepest, output);
        right := _lcaLeaves(root.right, depth+1, deepest, output);

        if ((left && right) || seenDeepest) {
            *output = root;
        }
        return left || right || seenDeepest;
    }
    return false;
}

func _lcaDeepestLeaves(root *TreeNode) *TreeNode {
    var deepest int = getDepth(root);
    var output *TreeNode;
    _lcaLeaves(root, 1, deepest, &output);
    return output;
}

func main() {
    root := &TreeNode{data: 3};
    root.left = &TreeNode{data: 5};
    root.left.left = &TreeNode{data: 6};
    root.left.right = &TreeNode{data: 2};
    root.left.right.left = &TreeNode{data: 7};
    root.left.right.right = &TreeNode{data: 4};
    fmt.Println(_lcaDeepestLeaves(root));
}
