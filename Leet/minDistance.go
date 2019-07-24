// file editing distance minDistance.go

package main

import (
  "fmt"
)

// applies memoization to saving results from recursive statements
func minDistWrapper(word1 string, word2 string, pos1 int, pos2 int, memo [][]int) int {
    if (memo[pos1][pos2] != -1) {
        return memo[pos1][pos2];
    }
    if (word1 == word2) {
        return 0;
    }
    if (pos1 == len(word1)) {
        memo[pos1][pos2] = len(word2)-pos2;
        return len(word2)-pos2;
    }
    if (pos2 == len(word2)) {
        memo[pos1][pos2] = len(word1)-pos1;
        return len(word1)-pos1;
    }
    if (word1[pos1] == word2[pos2]) {
        memo[pos1][pos2] = minDistWrapper(word1,word2, pos1+1, pos2+1, memo);
        return memo[pos1][pos2];
    }
    remove := minDistWrapper(word1,word2, pos1+1, pos2, memo);
    insert := minDistWrapper(word1,word2, pos1, pos2+1, memo);
    del := minDistWrapper(word1,word2, pos1+1,pos2+1, memo);
    // obtain minimum
    min := remove;
    if (insert < min) {
        min = insert;
    }
    if (del < min) {
        min = del;
    }
    memo[pos1][pos2] = 1+min;
    return memo[pos1][pos2];
}

func minDistance(word1 string, word2 string) int {
    if (len(word1) == 0) {
        return len(word2);
    } else if (len(word2) == 0){
        return len(word1);
    }
    pos1 := 0;
    pos2 := 0;
    memo := [][]int{};
    for i := 0; i <= len(word1); i++ {
        row := []int{};
        for j := 0; j <= len(word2); j++ {
            row = append(row,-1);
        }
        memo = append(memo,row);
    }
    return minDistWrapper(word1,word2, pos1, pos2, memo);
}
