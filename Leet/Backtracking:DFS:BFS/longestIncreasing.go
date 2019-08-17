// file longestIncreasing.go

package findLength

import (
  "fmt"
)

func findLengthOfLCIS(nums []int) int {
    if (len(nums) <= 1) {
        return len(nums);
    }
    longestStrictlyIncreasing := 0;
    currStreak := 1;
    for pos := 0; pos < len(nums)-1; pos++ {
        if (nums[pos] < nums[pos+1]) {
            currStreak++;
        } else {
            if (currStreak > longestStrictlyIncreasing) {
                longestStrictlyIncreasing = currStreak;
            }
            currStreak = 1;
        }
    }
    if (currStreak > longestStrictlyIncreasing) {
        return currStreak;
    } else {
        return longestStrictlyIncreasing;
    }
}

// finding length of longest subsequence - not optimal
// (dynamic programming better)
func lengthLIS(nums []int, curr []int, longest *int, pos int) {
    if (pos == len(nums)) {
        if (len(curr) > *longest) {
            *longest = len(curr);
        }
    } else {
        if (len(curr) == 0) {
            lengthLIS(nums, curr, longest, pos+1);
            newLst := curr;
            newLst = append(newLst, nums[pos]);
            lengthLIS(nums, newLst, longest, pos+1);
        } else if (curr[len(curr)-1] < nums[pos]) {
            lengthLIS(nums, curr, longest, pos+1);
            otherLst := curr;
            otherLst = append(otherLst, nums[pos]);
            lengthLIS(nums, otherLst, longest, pos+1);
        } else {
            lengthLIS(nums, curr, longest, pos+1);
        }
    }
}
func lengthOfLIS(nums []int) int {
    var currSubsequence []int;
    pos := 0;
    longest := 0;
    lengthLIS(nums, currSubsequence, &longest, pos);
    return longest;
}
