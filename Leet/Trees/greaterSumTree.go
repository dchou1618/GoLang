// file greaterSumTree.go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// perhaps implement prefixed-summation

// 0 ms (100%), 2.5 MB (100%).
func inOrderStorage(root *TreeNode, currSum *int) {
    // simply reverse inorder traversal and add to currSum (reverse-prefixed sum).
    if (root != nil) {
        inOrderStorage(root.Right, currSum);
        *currSum += root.Val;
        root.Val = *currSum;
        inOrderStorage(root.Left, currSum);
    }
}

func bstToGst(root *TreeNode) *TreeNode {
    curr := 0;
    inOrderStorage(root, &curr);
    return root;
}
