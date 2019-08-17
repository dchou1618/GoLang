// file mergeBinaryTrees.go

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

func mergeOverlap(node1 *TreeNode, node2 *TreeNode) {
    if (node1 != nil && node2 != nil) {
        node1.Val += node2.Val;
        if (node1.Left == nil && node2.Left != nil) {
            node1.Left = &TreeNode{0, nil, nil};
        }
        if (node1.Right == nil && node2.Right != nil) {
            node1.Right = &TreeNode{0, nil, nil};
        }
        mergeOverlap(node1.Left,node2.Left);
        mergeOverlap(node1.Right,node2.Right);
    }
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if (t1 != nil) {
        mergeOverlap(t1, t2);
        return t1;
    } else {
        mergeOverlap(t2, t1);
        return t2;
    }
}
