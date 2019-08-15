// file rollDice.go

// Ex: input = [1,2,3]

package main

import (
  "fmt"
  "math"
)

func minPipMoves(topFaces []int) []int {
    opposingSides := map[int]int{1:6, 6:1, 2:5, 5:2, 3:4, 4:3};
    minMoves := make([]int, len(topFaces)+1);
    minSum := math.MaxInt64;
    pipTarget := 0;
    // all equal to some number from 1 to 6
    for num := 1; num <= 6; num++ {
        moves := []int{};
        sum := 0;
        for _, face := range topFaces {
            if (opposingSides[face] == num) {
                sum += 2;
                moves = append(moves, 2);
            } else if (face == num) {
                moves = append(moves, 0);
            } else {
                sum++;
                moves = append(moves, 1);
            }
        }
        if (sum < minSum) {
            minSum = sum;
            minMoves = moves;
            moves = nil;
            pipTarget = num;
        }
    }
    fmt.Println(pipTarget);
    return minMoves;
}


func main() {
    dice := []int{1,1,6};
    fmt.Println(minPipMoves(dice));
}
