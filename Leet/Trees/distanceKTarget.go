/* */

package main

import (
    "fmt"
)

type TreeNode struct {
    data int
    left *TreeNode
    right *TreeNode
}
/**
 * Definition for a binary tree node.
   reinforce graph + trees
 */
func adjacencyLst(root *TreeNode, neighbors *map[*TreeNode][]*TreeNode) {
    if (root != nil) {
        if (root.left != nil) {
            (*neighbors)[root] = append((*neighbors)[root], root.left);
            (*neighbors)[root.left] = append((*neighbors)[root.left], root);
        }
        if (root.right != nil) {
            (*neighbors)[root] = append((*neighbors)[root], root.right);
            (*neighbors)[root.right] = append((*neighbors)[root.right], root);
        }
        adjacencyLst(root.left, neighbors);
        adjacencyLst(root.right, neighbors);
    }
}

/* Trees being acyclic graphs - BFS comes to mind with distance K from a target
   With BFS, remember to keep track of the visited nodes in the tree.
*/
func distanceK(root *TreeNode, target *TreeNode, K int) []int {
    if (K < 0) {
        return []int{};
    } else if (K == 0) {
        return []int{target.data};
    }
    var distanceK []int;
    neighbors := make(map[*TreeNode][]*TreeNode);
    adjacencyLst(root, &neighbors);
    queue := []*TreeNode{target};
    dists := []int{1};
    visited := make(map[int]bool);
    visited[target.data] = true;
    for (len(queue) != 0) {
        top := queue[0];
        dist := dists[0];
        queue = queue[1:];
        dists = dists[1:];
        for _, neighbor := range neighbors[top] {
            if (!visited[neighbor.data]) {
                if (dist == K) {
                    distanceK = append(distanceK, neighbor.data);
                }
                queue = append(queue, neighbor);
                dists = append(dists, dist+1);
                visited[neighbor.data] = true;
            }
        }
    }
    return distanceK;
}

func main() {
    root := &TreeNode{data:12};
    root.left = &TreeNode{data:8};
    root.left.left = &TreeNode{data:6};
    root.left.right = &TreeNode{data:7};
    root.left.right.left = &TreeNode{data:5};
    root.left.right.right = &TreeNode{data:4};
    root.right = &TreeNode{data:15};
    root.right.right = &TreeNode{data:14};
    fmt.Println(distanceK(root, root.left.right.left, 2));
}
