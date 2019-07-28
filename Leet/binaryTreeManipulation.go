// file binaryTreeManipulation.go


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
// idea of swapping left with right branches
func invertTree(root *TreeNode) *TreeNode {
    if (root != nil) {
        left := invertTree(root.Left);
        right := invertTree(root.Right);
        root.Right = left;
        root.Left = right;
        return root;
    }
    return nil;
}

func sumLeft(root *TreeNode, side int) int {
    if (root != nil) {
        if (root.Left == nil && root.Right == nil && side == 0) {
            return root.Val;
        }
        return sumLeft(root.Left, 0) + sumLeft(root.Right, 1);
    }
    return 0;
}

func sumOfLeftLeaves(root *TreeNode) int {
    return sumLeft(root, -1);
}

func sumNodes(root *TreeNode) int {
    if (root != nil) {
        return root.Val + sumNodes(root.Left) + sumNodes(root.Right);
    }
    return 0;
}

func nodeSumFrequency(root *TreeNode, freqs *[]int) {
    if (root != nil) {
        *freqs = append(*freqs, sumNodes(root));
        nodeSumFrequency(root.Left, freqs);
        nodeSumFrequency(root.Right, freqs);
    }
}
func findFrequentTreeSum(root *TreeNode) []int {
    freqSums := []int{};
    nodeSumFrequency(root, &freqSums);
    counts := make(map[int]int);
    max := 0;
    for _, val := range freqSums {
        counts[val]++;
        if (counts[val] > max) {
            max = counts[val];
        }
    } // clear
    freqSums = nil;
    sums := []int{};
    for sum, freq := range counts {
        if (freq == max) {
            sums = append(sums, sum);
        }
    }
    return sums;
}
