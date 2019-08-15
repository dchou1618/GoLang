// file spiralOrder.go implementation of

package main

import (
    "fmt"
)


func spiralOrder(matrix [][]int) []int {
    if (len(matrix) == 0 || len(matrix[0]) == 0) {
        return []int{};
    }
    seen := make(map[string]bool);
    dirs := [][]int{[]int{0,1},[]int{1,0},[]int{0,-1},[]int{-1,0}};
    dirIndex := 0;
    var currRow, currCol int = 0, 0;
    order := []int{};
    for (!seen[string(currRow) + "," + string(currCol)]) {
        if (currRow < 0 || currRow >= len(matrix) || currCol < 0 || currCol >= len(matrix[0])) {
            break;
        }
        order = append(order, matrix[currRow][currCol]);
        seen[string(currRow) + "," + string(currCol)] = true;
        currRow += dirs[dirIndex][0];
        currCol += dirs[dirIndex][1];
        if (seen[string(currRow) + "," + string(currCol)]) {
            currRow -= dirs[dirIndex][0];
            currCol -= dirs[dirIndex][1];
            dirIndex = (dirIndex+1)%4;
            // divert direction
            currRow += dirs[dirIndex][0];
            currCol += dirs[dirIndex][1];
        }
        if (currRow < 0 || currRow >= len(matrix)) {
            currRow -= dirs[dirIndex][0];
            dirIndex = (dirIndex+1)%4;
            currRow += dirs[dirIndex][0];
            currCol += dirs[dirIndex][1];
        }
        if (currCol < 0 || currCol >= len(matrix[0])) {
            currCol -= dirs[dirIndex][1];
            dirIndex = (dirIndex+1)%4;
            currRow += dirs[dirIndex][0];
            currCol += dirs[dirIndex][1];
        }
    }
    return order;
}

func main() {
    
}
