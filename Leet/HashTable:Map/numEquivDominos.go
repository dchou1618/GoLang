// file numEquivDominos.go

package main

import (
  "fmt"
)

func numEquivDominoPairs(dominoes [][]int) int {
    seen := make(map[string]int);
    pairs := 0;
    for i := 0; i < len(dominoes); i++ {
        _, included := seen[string(dominoes[i][0])+string(dominoes[i][1])];
        if (included) {
            seen[string(dominoes[i][0])+string(dominoes[i][1])] += 1;
        } else {
            _, include := seen[string(dominoes[i][1])+string(dominoes[i][0])];
            if (include) {
                seen[string(dominoes[i][1])+string(dominoes[i][0])] += 1;
            } else {
                seen[string(dominoes[i][0])+string(dominoes[i][1])] += 1;
            }
        }
    }
    for _, value := range seen {
        if (value > 1) {
            pairs += (value*(value-1))/2;
        } // nC2 quick hand
    }
    seen = nil;
    return pairs;
}
