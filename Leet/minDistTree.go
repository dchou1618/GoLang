// file minDistTree.go

// Leet 783 - Basic

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
func inOrder(root *TreeNode, vals *[]int){
    if (root != nil) {
        inOrder(root.Left,vals);
        *vals = append(*vals,root.Val);
        inOrder(root.Right,vals);
    }
}
func minDiffInBST(root *TreeNode) int {
    orderedBST := []int{};
    inOrder(root,&orderedBST);
    minDiff := 1<<32-1; // maximum uint 32 bytes
    for i := 0; i < len(orderedBST)-1; i++ {
        if (orderedBST[i+1]-orderedBST[i] < minDiff) {
            minDiff = orderedBST[i+1]-orderedBST[i];
        }
    }
    return minDiff;
}

func main() {

}
