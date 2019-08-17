// file trimTree.go


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, L int, R int) *TreeNode {
    if (root != nil) {
        if (root.Val > R) {
            return trimBST(root.Left, L, R); // exclude values above and below L, R
        } else if (root.Val < L) {
            return trimBST(root.Right, L, R);
        } else {
            root.Left = trimBST(root.Left, L, R);
            root.Right = trimBST(root.Right, L, R);
            return root;
        }
    } else {
        return nil;
    }
}
