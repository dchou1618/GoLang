// file isUniversalTree.go

package uvt_main

import (
  "fmt"
)

type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

func isUVT(root *TreeNode, val int) bool {
    if (root != nil) {
        if (root.Val != val) {
            return false;
        } else {
            return isUVT(root.Left, val) && isUVT(root.Right, val);
        }
    }
    return true;
}

func isUnivalTree(root *TreeNode) bool {
    if (root == nil) {
        return true;
    } else {
        val := root.Val;
        return isUVT(root, val);
    }
}

func uvt_main() {

}
