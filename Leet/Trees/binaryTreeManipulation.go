// file binaryTreeManipulation.go


package main

import (
  "fmt"
  "math"
  "strconv"
)

type BinaryTreeNode struct {
    data int;
    left *BinaryTreeNode;
    right *BinaryTreeNode;
};


func preOrder(root *BinaryTreeNode) {
    if (root != nil) {
        fmt.Print(root.data, " ");
        preOrder(root.left);
        preOrder(root.right);
    }
}


func removeMatchedLeaves(root1 **BinaryTreeNode, root2 *BinaryTreeNode) {
    if ((*root1) != nil && root2 != nil) {
        // preorder traversal ordering
        if ((*root1).data == root2.data && (*root1).left == nil && (*root1).right == nil) {
            (*root1) = (*root1).right;
            return;
        }
        removeMatchedLeaves(&(*root1).left, root2.left);
        removeMatchedLeaves(&(*root1).right, root2.right);
    }
}

func removeMatchingLeaves(root1 *BinaryTreeNode, root2 *BinaryTreeNode) *BinaryTreeNode {
    removeMatchedLeaves(&root1, root2);
    return root1;
}

func stretchWrapper(root *BinaryTreeNode, k int, left int) {
    if (root != nil) {
        var val int = root.data/k;
        if (left == -1) {
            var temp *BinaryTreeNode = root.left;
            var temp2 *BinaryTreeNode = root.right;
            var curr *BinaryTreeNode = &BinaryTreeNode{data:val};
            var base *BinaryTreeNode = curr;
            //BinaryTreeNode* curr = root;
            for i := 1; i < k; i++ {
                curr.left = &BinaryTreeNode{data:val};
                curr = curr.left;
            }
            curr.left = temp;
            curr.right = temp2;
            root = base;
            stretchWrapper(curr.left, k, -1);
            stretchWrapper(curr.right, k, 1);
        } else {
            var temp *BinaryTreeNode = root.left;
            var temp2 *BinaryTreeNode = root.right;
            var curr *BinaryTreeNode = &BinaryTreeNode{data:val};
            var base *BinaryTreeNode = curr;
            for i := 1; i < k; i++ {
                curr.right = &BinaryTreeNode{data:val};
                curr = curr.right;
            }
            curr.left = temp;
            curr.right = temp2;
            root = base;
            stretchWrapper(curr.right, k, 1);
            stretchWrapper(curr.left, k, -1);
        }
    }
}

func stretch(root *BinaryTreeNode, k int) {
    if (k <= 0) {
        panic(1);
    } else {
        if (root != nil && k > 0) {
            var val int = root.data/k;
            var temp *BinaryTreeNode = root.left;
            var temp2 *BinaryTreeNode = root.right;
            var curr *BinaryTreeNode = &BinaryTreeNode{data:val};
            var base *BinaryTreeNode = curr;
            for i := 1; i < k; i++ {
                base.left = &BinaryTreeNode{data:val};
                base = base.left;
            }
            base.left = temp;
            base.right = temp2;
            root = curr;
            stretchWrapper(base.left, k, -1);
            stretchWrapper(base.right, k, 1);
        }
    }
}


type TreeNode struct {
    Val int;
    Left *TreeNode;
    Right *TreeNode;
};


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMin(root *TreeNode) int {
    if (root != nil) {
        var currNode *TreeNode = root;
        for (currNode.Left != nil) {
            currNode = currNode.Left;
        }
        return currNode.Val;
    } else {
        panic(1);
    }
}

func nodeDeletion(root **TreeNode, key int) {
    if (*root != nil) { // check nullptr condition
        if ((*root).Val < key) {
            nodeDeletion(&(*root).Right, key);
        } else if ((*root).Val > key) {
            nodeDeletion(&(*root).Left, key);
        } else {
            if ((*root).Left == nil) {
                (*root) = (*root).Right;
            } else if ((*root).Right == nil) {
                (*root) = (*root).Left;
            } else {
                // neither left nor right is nil
                var min int = getMin((*root).Right);
                nodeDeletion(&(*root).Right, min);
                (*root).Val = min;
            }
        }
    }
}

func deleteNode(root *TreeNode, key int) *TreeNode {
    nodeDeletion(&root, key);
    return root;
}


/* Merging binary tree nodes */

func zipHelper(root *TreeNode, k int, step int) {
    if (root != nil) {
        if (step >= k) {
            step = 0;
            zipHelper(root.Left, k, step+1);
            zipHelper(root.Right, k, step+1);
        } else {
            if (root.Left != nil && root.Right == nil) {
                var val int = root.Val;
                root = root.Left;
                root.Val += val;
                zipHelper(root, k, step+1);
            } else if (root.Right != nil && root.Left == nil) {
                var val int = root.Val;
                root = root.Right;
                root.Val += val;
                zipHelper(root, k, step+1);
            } else {
                step := 0;
                zipHelper(root.Left, k, step+1);
                zipHelper(root.Right, k, step+1);
            }
        }
    }
}


func zip(root *TreeNode, k int) {
    if (k <= 0) {
        panic(1);
    } else if (k > 1) {
        zipHelper(root, k, 1);
    }
}


/* converting a tree into a string */
func treeToString(t *TreeNode) string {
    if (t != nil) {
        if (t.Left == nil && t.Right == nil) {
            return strconv.Itoa(t.Val);
        } else {
            if ((t.Left != nil && t.Right != nil) || (t.Right != nil)) {
                return strconv.Itoa(t.Val)+"("+treeToString(t.Left)+ ")" + "(" + treeToString(t.Right) + ")";
            } else {
                return strconv.Itoa(t.Val)+"("+treeToString(t.Left)+ ")";
            }
        }
    }
    return "";
}

func tree2str(t *TreeNode) string {
    return treeToString(t);
}

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


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstHelper(root *TreeNode, minVal int, maxVal int) bool {
    if (root != nil) {
        if ((maxVal != math.MinInt64 && root.Val >= maxVal) || (minVal != math.MaxInt64 && root.Val <= minVal)) {
            return false;
        }
        return bstHelper(root.Left, minVal, root.Val) && bstHelper(root.Right, root.Val, maxVal);
    }
    return true;
}
func isValidBST(root *TreeNode) bool {
    var minVal int = math.MaxInt64;
    var maxVal int = math.MinInt64;
    return bstHelper(root, minVal, maxVal);
}

// Find modal value in bst

type BinaryTree struct {
    root *BinaryTreeNode;
}

func modality(root *TreeNode, counts *map[int]int, max *int) {
    if (root != nil) {
        (*counts)[root.Val]++;
        if ((*counts)[root.Val] > *max) {
            *max = (*counts)[root.Val];
        }
        modality(root.Left, counts, max);
        modality(root.Right, counts, max);
    }
}
func findMode(root *TreeNode) []int {
    counts := make(map[int]int);
    var max int = math.MinInt64;
    modality(root, &counts, &max);
    var modes []int;
    for candidate, count := range counts {
        if (count == max) {
            modes = append(modes, candidate);
        }
    }
    return modes;
}


func existsPath(root *BinaryTreeNode, start int, end int,
                startSeen bool, endSeen bool, exists *bool) {
    if (root != nil) {
        if (root.data == start && !endSeen) { // seen start before end
            startSeen = true;
        }
        if (root.data == end) {
            endSeen = true;
        }
        if (startSeen && endSeen) {
            *exists = true;
            return;
        }
        existsPath(root.left, start, end, startSeen, endSeen, exists);
        existsPath(root.right, start, end, startSeen, endSeen, exists);
    }
}

func (this *BinaryTree) hasPath(start int, end int) bool {
    if (this.root != nil) {
        return false;
    } else {
        var exists bool = false;
        existsPath(this.root, start, end, false, false, &exists);
        return exists;
    }
}


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// sum of paths from root to leaves
func rootToLeaf(root *TreeNode, currSum int, totalSum *int) {
    if (root != nil) {
        if (root.Left == nil && root.Right == nil) {
            *totalSum += currSum*10 + root.Val;
        }
        rootToLeaf(root.Left, currSum*10 + root.Val, totalSum);
        rootToLeaf(root.Right, currSum*10 + root.Val, totalSum);
    }
}

func sumNumbers(root *TreeNode) int {
    var totalSum int;
    rootToLeaf(root, 0, &totalSum);
    return totalSum;
}


func sequentialEdit(root **BinaryTreeNode, lower int, upper int, left bool) {
    if (lower < upper) {
        var mid int = (lower+upper)/2;
        if (left) {
            (*root).left = &BinaryTreeNode{data:mid};
            sequentialEdit(&(*root).left, lower, mid-1, true);
            sequentialEdit(&(*root).left, mid+1, upper, false);
        } else {
            (*root).right = &BinaryTreeNode{data:mid};
            sequentialEdit(&(*root).right, lower, mid-1, true);
            sequentialEdit(&(*root).right, mid+1, upper, false);
        }
        //cout << "H";
    } else if (lower == upper) {
        var mid int = (lower+upper)/2;
        if (left) {
            (*root).left = &BinaryTreeNode{data:mid};
        } else {
            (*root).right = &BinaryTreeNode{data:mid};
        }
        return;
    } else {
        return;
    }
}

func sequential(N int) *BinaryTreeNode {
    if (N == 0) {
        return nil;
    } else if (N < 0) {
        panic(1);
    } else if (N == 1) {
        return &BinaryTreeNode{data:0};
    } else {
        var root *BinaryTreeNode;
        if (N%2 != 0) {
            var mid int = N/2;
            root = &BinaryTreeNode{data:mid};
            sequentialEdit(&root, 0, mid-1, true);
            sequentialEdit(&root, mid+1, N-1, false);
        } else {
            var mid int = (N-1)/2;
            root = &BinaryTreeNode{data:mid};
            sequentialEdit(&root, 0, mid-1, true);
            sequentialEdit(&root, mid+1, N-1, false);
        }
        return root;
    }
}


func main() {
    root1 := &BinaryTreeNode{data:1};
    root1.left = &BinaryTreeNode{data:2};
    root1.left.left = &BinaryTreeNode{data:3};
    root1.left.left.right = &BinaryTreeNode{data:2};
    root1.right = &BinaryTreeNode{data:6};
    root1.right.left = &BinaryTreeNode{data:2};
    root2 := &BinaryTreeNode{data:1};
    root2.left = &BinaryTreeNode{data:2};
    root2.left.left = &BinaryTreeNode{data:3};
    root2.left.left.right = &BinaryTreeNode{data:2};
    root2.right = &BinaryTreeNode{data:16};
    root2.right.left = &BinaryTreeNode{data:2};
    removeMatchedLeaves(&root1, root2);
    preOrder(root1);
    fmt.Println();
    root3 := sequential(10);
    preOrder(root3);
}
