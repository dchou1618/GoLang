// file longestConsecutive.go


package main

import (
  "fmt"
)
func getSet(nums []int) map[int]bool {
    set := make(map[int]bool);
    for _, val := range nums {
        if (!set[val]) {
            set[val] = true;
        }
    }
    return set;
}

func longestConsecutive(nums []int) int {
    if (len(nums) == 0) {
        return 0;
    }
    set := getSet(nums);
    seen := make(map[int]bool)
    longest := 1;
    for _, val := range nums {
        if (!seen[val] && !set[val-1]) {
            currLength := 1;
            for (set[val+1]) {
                if (set[val+1]) {
                    currLength++;
                }
                val++;
            }
            if (longest < currLength) {
                longest = currLength;
            }
        }
        seen[val] = true;
    }
    set = nil;
    seen = nil;
    return longest;
}
