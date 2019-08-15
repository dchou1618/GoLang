// targetSum.go file

//
package main

import (
    "fmt"
)

func targetSums(nums []int, S int, ways *int) {
    if (len(nums) == 0) {
        if (S == 0) {
            (*ways)++;
        }
    } else {
        var val int = nums[0];
        targetSums(nums[1:], S-val, ways);
        targetSums(nums[1:], S+val, ways);
    }
}
func findTargetSumWays(nums []int, S int) int {
    var ways int = 0;
    targetSums(nums, S, &ways);
    return ways;
}

func main() {
    
}
