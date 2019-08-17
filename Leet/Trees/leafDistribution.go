/* leafDistribution.go file */


package main

import (
    "fmt"
)

type TreeNode struct {
    data string;
    salary int;
    left *TreeNode;
    right *TreeNode;
};

func showTree(root *TreeNode, indent string) {
    if (root != nil) {
        showTree(root.right, indent + "  ");
        fmt.Println(indent, root.data,":",root.salary);
        showTree(root.left, indent + "  ");
    }
}

func storeIntoTree(root **TreeNode, newData string, salary int) {
    if (*root != nil) {
        if ((*root).salary > salary) {
            if ((*root).left == nil) {
                (*root).left = &TreeNode{data:newData, salary: salary,
                                      left: nil, right: nil};
                return;
            }
            storeIntoTree(&(*root).left, newData, salary);

        } else if ((*root).salary < salary) {
            if ((*root).right == nil) {
                (*root).right = &TreeNode{data:newData, salary: salary,
                                       left: nil, right: nil};
                return;
            }
            storeIntoTree(&(*root).right, newData, salary);
        }
    } else {
        *root = &TreeNode{data:newData, salary: salary, left: nil, right: nil};
    }
}

func getLeafDistribution(root *TreeNode) map[string]int {
    
}

func main() {
    var root *TreeNode;
    storeIntoTree(&root, "Dylan", 100000);
    storeIntoTree(&root, "Daniel", 90000);
    storeIntoTree(&root, "Rufus", 180000);
    storeIntoTree(&root, "Monroe", 95000);
    showTree(root, "");
}
