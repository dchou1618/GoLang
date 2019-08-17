// file pathFinding.go

package main

import (
    "fmt"
)

/* finding nodes in complete tree using bfs level-order */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// breadth-first-search
func countNodes(root *TreeNode) int {
    if (root == nil) {
        return 0;
    } else {
        // level-order traversal
        numNodes := 1;
        queue := []*TreeNode{root};
        for (len(queue) != 0) {
            var top *TreeNode = queue[0];
            queue = queue[1:];
            if (top.Left != nil) {
                queue = append(queue, top.Left);
                numNodes++;
            }
            if (top.Right != nil) {
                queue = append(queue, top.Right);
                numNodes++;
            } else {
                return numNodes;
            }
        }
        queue = nil;
        return numNodes;
    }
}

// implementing BFS with largest value per level in binary tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    if (root == nil) {
        return []int{};
    }
    queue := []*TreeNode{root}; // queue with independent counter
    rows := []int{0};
    var largest []int;
    for (len(queue) != 0) {
        top := queue[0];
        row := rows[0];
        queue = queue[1:];
        rows = rows[1:];
        if (len(largest) == 0 || len(largest) <= row) {
            largest = append(largest, top.Val);
        } else {
            if (largest[row] < top.Val) {
                largest[row] = top.Val;
            }
        }
        if (top.Left != nil && top.Right != nil) {
            rows = append(rows, row+1);
            rows = append(rows, row+1);
            queue = append(queue, top.Left);
            queue = append(queue, top.Right);
        } else if (top.Left != nil) {
            rows = append(rows, row+1);
            queue = append(queue, top.Left);
        } else if (top.Right != nil) {
            rows = append(rows, row+1);
            queue = append(queue, top.Right);
        }
    }
    return largest;
}


// finding unique paths in grid by marking visited locations with -1 (obstacles)
func isValid(grid [][]int, r int, c int) bool {
    return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[r]) && grid[r][c] != -1;
}

func uniquePathsDFS(grid [][]int, r int, c int, walked int, paths *int) {
    if (isValid(grid, r, c) && grid[r][c] == 2 && walked == -1) {
        *paths++;
        return;
    } else if (grid[r][c] == 2) {
        return;
    } else if (isValid(grid, r, c)) {
        grid[r][c] = -1; // marks visited location
        if (isValid(grid, r+1, c)) {
            walked--;
            uniquePathsDFS(grid, r+1, c, walked, paths);
            walked++;
        }
        if (isValid(grid, r, c-1)) {
            walked--;
            uniquePathsDFS(grid, r, c-1, walked, paths);
            walked++;
        }
        if (isValid(grid, r-1, c)) {
            walked--;
            uniquePathsDFS(grid, r-1, c, walked, paths);
            walked++;
        }
        if (isValid(grid, r, c+1)) {
            walked--;
            uniquePathsDFS(grid, r, c+1, walked, paths);
            walked++;
        }
        grid[r][c] = 0;
    }
}

func uniquePathsIII(grid [][]int) int {
    numPaths := 0;
    nonObstacles := 0;
    var startRow, startCol int;
    for r, row := range grid {
        for c, elem := range row {
            if (elem == 1) {
                startRow = r;
                startCol = c;
            }
            if (elem == 0) {
                nonObstacles++;
            }
        }
    }
    uniquePathsDFS(grid, startRow, startCol, nonObstacles, &numPaths);
    return numPaths;
}

// implementing dijkstra's algorithm

// dijkstra is best suited for sparse graphs - priority queue
// remove min-cost vertex from pq

func leastCostPath(paths [][]int) []int {
    if (len(paths) > 0) {
        
        for (len(paths) != 0) {

        }
    }
}






// First consider nodes with infinite cost
// initially store only one
