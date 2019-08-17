/* limitLevelSum.go file */

package main

import (
    "fmt"
)

type TreeNode struct {
    data int;
    left *TreeNode;
    right *TreeNode;
};

func displayTree(root *TreeNode, indent string) {
    if (root != nil) {
        displayTree(root.left, indent + "  ");
        fmt.Println(indent, root.data);
        displayTree(root.right, indent + "  ");
    }
}

func _removeAll(root **TreeNode) {
    if (*root != nil) {
        _removeAll(&(*root).left);
        _removeAll(&(*root).right);
        *root = nil;
    }
}

func _limitSum(root **TreeNode, sum int, depth int, level int,
               currSum int, removed *[]int) {
    if (*root != nil) {
        if (depth == level) {
            currSum += (*root).data;
            if (currSum > sum) {
                *removed = append(*removed, (*root).data);
                _removeAll(root);
                return;
            }
        }
        _limitSum(&(*root).left, sum, depth+1, level, currSum, removed);
        _limitSum(&(*root).right, sum, depth+1, level, currSum, removed);
    }
}

func _limitLevelSum(root *TreeNode, level int, sum int) []int {
    var removed []int;
    _limitSum(&root, sum, 1, level, 0, &removed);
    return removed;
}


func main() {
    root := &TreeNode{data: 19};
    root.left = &TreeNode{data: 10};
    root.right = &TreeNode{data: 23};
    root.right.left = &TreeNode{data: 21};
    root.left.right = &TreeNode{data: 15};
    fmt.Println(_limitLevelSum(root, 3, 6));
    displayTree(root, "");
}
