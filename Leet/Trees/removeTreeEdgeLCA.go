/* removeTreeEdgeLCA.go file implements concepts */

package main

import (
    "fmt"
)

type TreeNode struct {
    data int;
    left *TreeNode;
    right *TreeNode;
};

func getAncestor(root, p, q *TreeNode, output **TreeNode) bool{
    if (root != nil) {
        pqItself := false;
        if (root == p || root == q) {
            pqItself = true;
        }
        left := getAncestor(root.left, p, q, output);
        right := getAncestor(root.right, p, q, output);
        if ((left && right) || (pqItself && right) || (pqItself && left)) {
            *output = root;
        }
        return pqItself || left || right;
    }
    return false;
}

func _LCATree(root, p, q *TreeNode) *TreeNode {
    var output *TreeNode;
    getAncestor(root, p, q, &output);
    return output;
}

/* non-binary search tree */
func _removeEdge(root *TreeNode) {

}


func main() {
    fmt.Println();
}
