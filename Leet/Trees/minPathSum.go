// file minPathSum.go

package main

import (
  "fmt"
)


func isValid(grid [][]int, r int, c int) bool {
    return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[r]);
}

func minPathSum(grid [][]int) int {
    // perceiving grid as minDP
    for r := range grid {
        for c := range grid[r] {
            var candidate1 int;
            var candidate2 int;
            if (isValid(grid, r-1, c)) {
                candidate1 = grid[r-1][c];
            } else {
                candidate1 = -1;
            }
            if (isValid(grid, r, c-1)) {
                candidate2 = grid[r][c-1];
            } else {
                candidate2 = -1;
            }
            if (candidate1 == -1 && candidate2 == -1) {
                continue;
            } else if (candidate1 == -1) {
                grid[r][c] += candidate2;
            } else if (candidate2 == -1) {
                grid[r][c] += candidate1;
            } else if (candidate1 > candidate2) {
                grid[r][c] += candidate2;
            } else {
                grid[r][c] += candidate1;
            }
        }
    }
    return grid[len(grid)-1][len(grid[0])-1];
}



func main() {
    fmt.Println();

}
