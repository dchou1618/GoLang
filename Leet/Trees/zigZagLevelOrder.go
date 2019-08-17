// file zigZagLevelOrder.go

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

 // Runtime - 9%, Memory - 14.81%
func findHeight(root *TreeNode) int {
    if (root != nil) {
        left := 1+findHeight(root.Left);
        right := 1+findHeight(root.Right);
        if (left>right) {
            return left;
        } else {
            return right;
        }
    }
    return 0;
}

func inOrder(root *TreeNode, zigLevel *[][]int, depth int) {
    if (root != nil) {
        inOrder(root.Left, zigLevel, depth+1);
        if (depth%2 == 0) {
            (*zigLevel)[depth] = append((*zigLevel)[depth], root.Val);
        } else {
            (*zigLevel)[depth] = append([]int{root.Val},(*zigLevel)[depth]...);
        }
        inOrder(root.Right, zigLevel, depth+1);
    }
    return;
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    zigLevel := make([][]int,findHeight(root));
    inOrder(root, &zigLevel, 0);
    return zigLevel;
}
