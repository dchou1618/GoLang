// file bstIterator.go


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

// runtime 48ms (37.64%), memory 11 mb(84%).

type BSTIterator struct {
    arr []int
    currIndex int
}

func inOrder(root *TreeNode, lst *[]int) {
    if (root != nil) {
        inOrder(root.Right, lst);
        *lst = append(*lst, root.Val);
        inOrder(root.Left, lst);
    }
}

func Constructor(root *TreeNode) BSTIterator {
    this := new(BSTIterator);
    lst := []int{};
    inOrder(root, &lst);
    this.arr = lst;
    this.currIndex = len(this.arr)-1;
    lst = nil;
    return *this;
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    curr := this.currIndex;
    this.currIndex--;
    return this.arr[curr];
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return this.currIndex >= 0;
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
