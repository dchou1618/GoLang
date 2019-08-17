// file numIslands.go

package main

import (
  "fmt"
)



func isValid(grid [][]int, r int, c int) bool {
    return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[r]);
}

func getNeighbors(grid [][]int, r int, c int) [][]int {
    neighbors := [][]int{};
    if (grid[r][c] != 0) {
        if (isValid(grid, r+1, c) && grid[r+1][c] != 0) {
            neighbors = append(neighbors, []int{r+1,c});
        }
        if (isValid(grid, r-1, c) && grid[r-1][c] != 0) {
            neighbors = append(neighbors, []int{r-1,c});
        }
        if (isValid(grid, r, c+1) && grid[r][c+1] != 0) {
            neighbors = append(neighbors, []int{r,c+1});
        }
        if (isValid(grid, r, c-1) && grid[r][c-1] != 0) {
            neighbors = append(neighbors, []int{r,c-1});
        }
    }
    return neighbors
}

func textify(grid [][]byte) [][]int {
    newGrid := [][]int{};
    for r := 0; r < len(grid); r++ {
        row := make([]int, len(grid[0]));
        newGrid = append(newGrid,row);
    }
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[r]); c++ {
            if (grid[r][c] == '1') {
                newGrid[r][c] = math.MaxInt64;
            } else {
                newGrid[r][c] = 0;
            }
        }
    }
    return newGrid;
}

func fillCells(grid *[][]int, r int, c int, num int) {
    if (isValid(*grid, r, c) && (*grid)[r][c] != num && (*grid)[r][c] != 0) {
        (*grid)[r][c] = num;
        fillCells(grid, r+1, c, num);
        fillCells(grid, r, c+1, num);
        fillCells(grid, r-1, c, num);
        fillCells(grid, r, c-1, num);
    }
}


func numIslands(grid [][]byte) int {
    currNumIslands := 0;
    islandMap := textify(grid);
    for r := 0; r < len(islandMap); r++ {
        for c := 0; c < len(islandMap[0]); c++ {
            // always 4 neighbors
            if (islandMap[r][c] == math.MaxInt64) {
                currNumIslands++;
                fillCells(&islandMap, r, c, currNumIslands);
            }
        }
    }
    islandMap = nil;
    return currNumIslands;
}
