// file bstLeaves.go

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

func lexSmallest(root *TreeNode, path string, smallest *string) {
    if (root != nil) {
        if (root.Left == nil && root.Right == nil) {
            if (string(97 + root.Val) + path < *smallest) {
                *smallest = string(97 + root.Val) + path;
            }
        }
        path = string(97+root.Val) + path;
        lexSmallest(root.Left, path, smallest);
        lexSmallest(root.Right, path, smallest);
    }
}

func smallestFromLeaf(root *TreeNode) string {
    smallest := "z";
    lexSmallest(root, "", &smallest);
    return smallest;
}

func bstPaths(root *TreeNode, path string, paths *[]string) {
    if (root != nil) {
        if (root.Left == nil && root.Right == nil) {
            *paths = append(*paths, path + strconv.Itoa(root.Val));
            return;
        }
        path += strconv.Itoa(root.Val) + "->"
        bstPaths(root.Left, path, paths);
        bstPaths(root.Right, path, paths);
    }
}

func binaryTreePaths(root *TreeNode) []string {
    if (root == nil) {
        return []string{};
    } else {
        paths := []string{};
        bstPaths(root, "", &paths);
        return paths;
    }
}
