// file binaryTreeDiameter.go

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
func getHeight(root *TreeNode) int {
    if (root != nil) {
        left := 1 + getHeight(root.Left);
        right := 1 + getHeight(root.Right);
        if (left > right) {
            return left;
        } else {
            return right;
        }
    }
    return 0;
}

func maxDiameter(root *TreeNode, max *int) {
    if (root != nil) {
        left := getHeight(root.Left);
        right := getHeight(root.Right);
        if (left+right > *max) {
            *max = left+right;
        }
        maxDiameter(root.Left, max);
        maxDiameter(root.Right, max);
    }
}


func diameterOfBinaryTree(root *TreeNode) int {
    max := 0;
    maxDiameter(root, &max);
    return max;
}
