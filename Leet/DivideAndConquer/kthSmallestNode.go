// file kthSmallestNode.go

package main

import (
  "fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inOrder(root *TreeNode, lst *[]int) {
    if (root != nil) {
        inOrder(root.Left,lst)
        *lst = append(*lst, root.Val);
        inOrder(root.Right,lst);
    }
    return;
}

func kthSmallest(root *TreeNode, k int) int {
    // sol'n in O(n), but memory usage can be improved
    orderedLst := []int{};
    inOrder(root, &orderedLst);
    return orderedLst[k-1];
}
