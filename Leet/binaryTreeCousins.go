// file binaryTreeCousins.go

package main

import (
  "fmt"
)

// problem 993: Cousins in binary tree: (0 ms 100% & 2.5 MB memory 75%)
func depthAndParent(root *TreeNode, x int, depth int) ([]int) {
    if (root != nil) {
        if (root.Val == x) {
            return []int{root.Val,depth};
        }
        if (root.Left != nil) {
            if (root.Left.Val == x) {
                return []int{root.Val,depth+1};
            }
        }
        if (root.Right != nil) {
            if (root.Right.Val == x) {
                return []int{root.Val,depth+1};
            }
        }
        left := depthAndParent(root.Left,x,depth+1);
        right := depthAndParent(root.Right,x,depth+1);
        if (left != nil) {
            return left;
        } else {
            return right;
        }
    }
    return nil;
}
func isCousins(root *TreeNode, x int, y int) bool {
    parentDepthX := depthAndParent(root, x, 0);
    parentDepthY := depthAndParent(root, y, 0);
    return parentDepthX[0] != parentDepthY[0] && parentDepthX[1] == parentDepthY[1];
}
